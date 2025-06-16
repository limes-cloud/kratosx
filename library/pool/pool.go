package pool

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/panjf2000/ants/v2"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/logger"
)

type Runner interface {
	Run()
}

type Pool interface {
	Go(runner Runner) error
}

type pool struct {
	pf *ants.PoolWithFunc
}

var (
	ins *pool
)

func Instance() Pool {
	return ins
}

func Init(conf *config.Pool, watcher config.Watcher) {
	if conf == nil {
		return
	}

	p, err := ants.NewPoolWithFunc(conf.Size, func(i any) {
		if run, ok := i.(Runner); ok {
			run.Run()
		}
	},
		ants.WithExpiryDuration(conf.ExpiryDuration),
		ants.WithMaxBlockingTasks(conf.MaxBlockingTasks),
		ants.WithNonblocking(conf.Nonblocking),
		ants.WithPreAlloc(conf.PreAlloc),
	)
	if err != nil {
		panic("协程池初始化失败：" + err.Error())
	}

	ins = &pool{pf: p}

	watcher("pool.size", func(value config.Value) {
		size, err := value.Int()
		if err != nil {
			log.Errorf("Pool配置变更失败：%s", err.Error())
			return
		}
		if size != 0 {
			ins.pf.Tune(int(size))
		}
	})
}

// Go 执行协程任务
func (c *pool) Go(runner Runner) error {
	return c.pf.Invoke(runner)
}

type runner struct {
	ctx context.Context
	fn  func()
}

// Run 执行协程任务
func (r runner) Run() {
	select {
	case <-r.ctx.Done():
		return
	default:
		r.fn()
	}
}

// AddRunner 添加协程任务
func AddRunner(ctx context.Context, fn func()) Runner {
	return &runner{fn: fn, ctx: ctx}
}

type WaitRunner interface {
	AddTasks(f ...func() error)
	AddTask(f func() error)
	Wait() error
	ErrorList() []error
}

type waitRunner struct {
	options *WaitRunnerOption
	size    atomic.Int32
	wg      sync.WaitGroup
	ctx     context.Context
	err     error
	errList []error
	errMux  sync.Mutex
}

type waitTask struct {
	runner *waitRunner
	fn     func() error
}

type WaitRunnerOption struct {
	max           int
	errorBreak    bool
	retryCount    int
	retryWaitTime time.Duration
}

type WaitRunnerOptionFunc func(o *WaitRunnerOption)

// WithMaxWaitRunnerOption 限制最大的等待队列数量
func WithMaxWaitRunnerOption(max int) WaitRunnerOptionFunc {
	return func(o *WaitRunnerOption) {
		o.max = max
	}
}

// WithErrorBreakOption 在执行过程中，如果遇到错误则停止后续的执行
func WithErrorBreakOption(is bool) WaitRunnerOptionFunc {
	return func(o *WaitRunnerOption) {
		o.errorBreak = is
	}
}

// WithRetryCountOption 在执行过程中，如果遇到错误是否重试执行
func WithRetryCountOption(retryCount int) WaitRunnerOptionFunc {
	if retryCount <= 0 {
		retryCount = 1
	}
	return func(o *WaitRunnerOption) {
		o.retryCount = retryCount
	}
}

// WithRetryWaitTimeOption 在执行过程中，如果遇到错误则等待一段时间后重试执行
func WithRetryWaitTimeOption(duration time.Duration) WaitRunnerOptionFunc {
	return func(o *WaitRunnerOption) {
		o.retryWaitTime = duration
	}
}

// Run 执行协程任务
func (w *waitTask) Run() {
	defer func() {
		if r := recover(); r != nil {
			logger.Helper().WithContext(w.runner.ctx).Errorw("exec task error", r)
		}
		w.runner.size.Add(-1)
		w.runner.wg.Done()
	}()

	// 如果已经出错，则不再执行后续任务
	if w.runner.options.errorBreak && w.runner.err != nil {
		return
	}

	var (
		count = 0
		err   error
	)
	for count < w.runner.options.retryCount {
		count++
		select {
		case <-w.runner.ctx.Done():
			break
		default:
			if err = w.fn(); err == nil {
				return
			}
			if w.runner.options.retryWaitTime > 0 {
				time.Sleep(w.runner.options.retryWaitTime)
			}
		}
	}
	if err != nil {
		w.runner.errMux.Lock()
		w.runner.err = err
		w.runner.errList = append(w.runner.errList, err)
		w.runner.errMux.Unlock()
	}
}

// NewWaitRunner 创建等待协程任务
func NewWaitRunner(ctx context.Context, opts ...WaitRunnerOptionFunc) WaitRunner {
	opt := &WaitRunnerOption{
		max:        -1,
		errorBreak: false,
		retryCount: 1,
	}
	for _, o := range opts {
		o(opt)
	}

	return &waitRunner{
		ctx:     ctx,
		options: opt,
		wg:      sync.WaitGroup{},
		size:    atomic.Int32{},
		err:     nil,
		errList: []error{},
		errMux:  sync.Mutex{},
	}
}

// AddTasks 批量添加协程任务
func (w *waitRunner) AddTasks(fs ...func() error) {
	if len(fs) == 0 {
		return
	}
	for _, f := range fs {
		w.AddTask(f)
	}
}

// AddTask 添加协程任务
func (w *waitRunner) AddTask(f func() error) {
	// 如果开启了出错，则不再执行后续任务
	if w.options.errorBreak && w.err != nil {
		return
	}

	w.wg.Add(1)

	// 当前运行池已满，则等待一段时间再执行
	for w.options.max > 0 && w.size.Load() >= int32(w.options.max) {
		time.Sleep(100 * time.Microsecond)
	}

	w.size.Add(1)
	if err := ins.Go(w.taskRunner(f)); err != nil {
		w.size.Add(-1)
		w.wg.Done()
	}
}

// taskRunner 协程任务
func (w *waitRunner) taskRunner(f func() error) Runner {
	return &waitTask{
		runner: w,
		fn:     f,
	}
}

// Wait 等待协程任务执行完成
func (w *waitRunner) Wait() error {
	w.wg.Wait()
	return w.err
}

// ErrorList 获取错误列表
// 请在Wait之后调用
func (w *waitRunner) ErrorList() []error {
	return w.errList
}
