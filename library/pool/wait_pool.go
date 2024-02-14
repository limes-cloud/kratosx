package pool

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
)

type WaitRunner interface {
	AddTasks(ctx context.Context, f ...func() error)
	AddTask(ctx context.Context, f func() error)
	Wait()
}

type waitRunner struct {
	wg *sync.WaitGroup
}

type waitTask struct {
	wg *sync.WaitGroup
	fn func() error
}

// Run 执行协程任务
func (w *waitTask) Run() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
		w.wg.Done()
	}()
	return w.fn()
}

// NewWaitRunner 创建等待协程任务
func NewWaitRunner() WaitRunner {
	return &waitRunner{
		wg: &sync.WaitGroup{},
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
		if err := ins.Go(w.taskRunner(f)); err != nil {
			log.Error(ctx, fmt.Sprintf("添加携程任务失败:%s", ctx.Err()))
		}
	}
}

// taskRunner 协程任务
func (w *waitRunner) taskRunner(f func() error) Runner {
	return &waitTask{
		wg: w.wg,
		fn: f,
	}
}

// Wait 等待协程任务执行完成
func (w *waitRunner) Wait() {
	w.wg.Wait()
}
