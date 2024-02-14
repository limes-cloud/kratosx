package pool

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/panjf2000/ants/v2"

	"github.com/limes-cloud/kratosx/config"
)

type Runner interface {
	Run() error
}

type Pool interface {
	Go(runner Runner) error
}

type pool struct {
	pf *ants.PoolWithFunc
}

var ins *pool

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

func (c *pool) Go(runner Runner) error {
	return c.pf.Invoke(runner)
}

type runner struct {
	fn func() error
}

func (r runner) Run() error {
	return r.fn()
}

func AddRunner(fn func() error) Runner {
	return &runner{fn: fn}
}
