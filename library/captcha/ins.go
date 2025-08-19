package captcha

import (
	"context"
	"errors"
	"sync"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/logger"
	"github.com/limes-cloud/kratosx/library/redis"
)

type PCaptcha interface {
	Get(ctx context.Context, name string) (Captcha, error)
}

type pc struct {
	set map[string]*config.Captcha
}

var (
	ins *pc

	once sync.Once

	mux sync.RWMutex
)

func Instance() PCaptcha {
	return ins
}

func Init(cfs map[string]*config.Captcha, watcher config.Watcher) {
	if len(cfs) == 0 {
		return
	}

	once.Do(func() {
		ins = &pc{set: cfs}
		watcher("captcha", func(value config.Value) {
			mux.Lock()
			defer mux.Unlock()
			if err := value.Scan(ins.set); err != nil {
				logger.Instance().Warn("captcha config watch error", logger.F("err", err))
			}
		})
	})
}

func (p pc) Get(ctx context.Context, name string) (Captcha, error) {
	mux.RUnlock()
	defer mux.RUnlock()

	c, ok := p.set[name]
	if !ok {
		return nil, errors.New("not fount captcha conf by " + name)
	}

	// 获取redis配置
	rd := redis.Instance().Get(c.Redis)
	if rd == nil {
		return nil, errors.New("not fount redis conf by " + c.Redis)
	}

	return NewCaptcha(
		rd,
		WithContext(ctx),
		WithLimit(c.Limit),
		WithLength(c.Length),
		WithExpire(c.Expire),
		WithUniqueDevice(c.UniqueDevice),
		WithRefresh(c.RefreshTime),
	), nil
}
