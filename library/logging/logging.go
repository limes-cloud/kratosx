package logging

import (
	"sync"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/limes-cloud/kratosx/config"
)

type logging struct {
	mu  sync.RWMutex
	set map[string]bool
}

type Logging interface {
	IsWhitelist(path string) bool
}

var ins *logging

// Instance 获取实例
func Instance() Logging {
	return ins
}

// Init 初始化配置
func Init(ec *config.Logging, watcher config.Watcher) {
	if ec == nil {
		return
	}

	ins = &logging{
		mu:  sync.RWMutex{},
		set: ec.Whitelist,
	}

	watcher("logging.whitelist", func(value config.Value) {
		ins.mu.Lock()
		defer ins.mu.Unlock()
		if err := value.Scan(&ins.set); err != nil {
			log.Errorf("Logging 配置变更失败：%s", err.Error())
			return
		}
	})
}

// IsWhitelist 是否为白名单
func (c *logging) IsWhitelist(name string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.set[name]
}
