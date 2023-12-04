package pool

import (
	kratosConfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/limes-cloud/kratosx/config"
	"github.com/panjf2000/ants/v2"
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

var instance *pool

func Instance() Pool {
	return instance
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

	instance = &pool{pf: p}

	watcher("pool.size", func(value kratosConfig.Value) {
		size, err := value.Int()
		if err != nil {
			log.Errorf("Pool配置变更失败：%s", err.Error())
			return
		}
		if size != 0 {
			instance.pf.Tune(int(size))
		}
	})
}

func (c *pool) Go(runner Runner) error {
	return c.pf.Invoke(runner)
}
