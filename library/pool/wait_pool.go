package pool

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type WaitRunner interface {
	AddTasks(ctx context.Context, f ...func() error)
	AddTask(ctx context.Context, f func() error)
	Wait()
}

type waitRunner struct {
	max  int32
	size *atomic.Int32
	wg   *sync.WaitGroup
}

type waitTask struct {
	size *atomic.Int32
	wg   *sync.WaitGroup
	fn   func() error
	ctx  context.Context
}

type WaitRunnerOption struct {
	max int32
}

type WaitRunnerOptionFunc func(o *WaitRunnerOption)

func WithWaitRunnerOption(max int32) WaitRunnerOptionFunc {
	return func(o *WaitRunnerOption) {
		o.max = max
	}
}

func (w *waitTask) Ctx() context.Context {
	return w.ctx
}

// Run 执行协程任务
func (w *waitTask) Run() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("执行协程任务失败：%v", r)
		}
		w.wg.Done()
		w.size.Add(-1)
	}()
	select {
	case <-w.ctx.Done():
		return w.ctx.Err()
	default:
		return w.fn()
	}
}

// NewWaitRunner 创建等待协程任务
func NewWaitRunner(opts ...WaitRunnerOptionFunc) WaitRunner {
	opt := &WaitRunnerOption{
		max: -1,
	}
	for _, o := range opts {
		o(opt)
	}

	return &waitRunner{
		wg:   &sync.WaitGroup{},
		max:  opt.max,
		size: &atomic.Int32{},
	}
}

// AddTasks 批量添加协程任务
func (w *waitRunner) AddTasks(ctx context.Context, fs ...func() error) {
	if len(fs) == 0 {
		return
	}
	for _, f := range fs {
		w.AddTask(ctx, f)
	}
}

// AddTask 添加协程任务
func (w *waitRunner) AddTask(ctx context.Context, f func() error) {
	select {
	case <-ctx.Done():
		log.Error(ctx, fmt.Sprintf("添加携程任务失败:%s", ctx.Err()))
	default:
		w.wg.Add(1)
		for w.max > 0 && w.size.Load() >= w.max {
			time.Sleep(100 * time.Microsecond)
		}

		w.size.Add(1)
		if err := ins.Go(w.taskRunner(ctx, f)); err != nil {
			w.size.Add(-1)
			w.wg.Done()
			log.Error(ctx, fmt.Sprintf("添加携程任务失败:%s", ctx.Err()))
		}
	}
}

// taskRunner 协程任务
func (w *waitRunner) taskRunner(ctx context.Context, f func() error) Runner {
	return &waitTask{
		wg:   w.wg,
		size: w.size,
		fn:   f,
		ctx:  ctx,
	}
}

// Wait 等待协程任务执行完成
func (w *waitRunner) Wait() {
	w.wg.Wait()
}
