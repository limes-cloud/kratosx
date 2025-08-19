package pool

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/logger"
)

type Runner interface {
	Run()
}

type Pool interface {
	// WithContext 返回一个上下文
	WithContext(ctx context.Context) Pool

	// Go 执行任务
	Go(runner Runner) error

	// GoFunc 执行函数任务
	GoFunc(fn func()) error

	// NewWaitRunner 创建等待协程任务
	NewWaitRunner(opts ...WaitRunnerOptionFunc) WaitRunner
}

type pool struct {
	pf  *ants.PoolWithFunc
	ctx context.Context
}

var (
	ins  *pool
	once sync.Once
)

const (
	poolSize             = 100000
	poolMaxBlockingTasks = 10000
)

// Instance 获取全局日志器
func Instance() Pool {
	return ins
}

// Init 初始化协程
func Init(conf *config.Pool) {
	if conf == nil {
		return
	}

	once.Do(func() {
		if conf.Size == 0 {
			conf.Size = poolSize
		}
		if conf.MaxBlockingTasks == 0 {
			conf.MaxBlockingTasks = poolMaxBlockingTasks
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
		ins = &pool{pf: p, ctx: context.Background()}
	})
}

// GoFunc 执行函数任务
func (p *pool) GoFunc(fn func()) error {
	return p.pf.Invoke(AddRunner(p.ctx, fn))
}

// Go 执行协程任务
func (p *pool) Go(runner Runner) error {
	return p.pf.Invoke(runner)
}

// WithContext 载入业务ctx
func (p *pool) WithContext(ctx context.Context) Pool {
	return &pool{pf: p.pf, ctx: ctx}
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
	// AddTasks 批量添加协程任务
	AddTasks(f ...func() error)

	// AddTask 添加协程任务
	AddTask(f func() error)

	// Wait 等待协程任务执行完成
	Wait() error

	// ErrorList 获取协程任务执行错误列表
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
			logger.Instance().WithContext(w.runner.ctx).Error("exec task error", logger.F("err", r))
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
func (p *pool) NewWaitRunner(opts ...WaitRunnerOptionFunc) WaitRunner {
	opt := &WaitRunnerOption{
		max:        -1,
		errorBreak: false,
		retryCount: 1,
	}
	for _, o := range opts {
		o(opt)
	}

	return &waitRunner{
		ctx:     p.ctx,
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
