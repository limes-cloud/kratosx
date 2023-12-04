package logging

import (
	kratosConfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/limes-cloud/kratosx/config"
	"sync"
)

type logging struct {
	mu  sync.RWMutex
	set map[string]bool
}

type Logging interface {
	IsWhitelist(path string) bool
}

var instance *logging

func Instance() Logging {
	return instance
}

func Init(ec *config.Logging, watcher config.Watcher) {
	if ec == nil {
		return
	}

	instance = &logging{
		mu:  sync.RWMutex{},
		set: ec.Whitelist,
	}

	newConf := map[string]bool{}
	watcher("logging.whitelist", func(value kratosConfig.Value) {
		if err := value.Scan(&newConf); err != nil {
			log.Errorf("Logging 配置变更失败：%s", err.Error())
			return
		}

		instance.mu.Lock()
		instance.set = newConf
		instance.mu.RLock()
	})
}

func (c *logging) IsWhitelist(name string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.set[name]
}
