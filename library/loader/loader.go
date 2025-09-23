package loader

import (
	"io"
	"os"
	"sync"

	"github.com/limes-cloud/kratosx/library/logger"

	"github.com/limes-cloud/kratosx/config"
)

type loader struct {
	mux  sync.RWMutex
	conf map[string]string
	set  map[string][]byte
}

type Loader interface {
	Get(name string) []byte
}

var (
	ins *loader

	once sync.Once
)

func Instance() Loader {
	return ins
}

func Init(conf map[string]string, watcher config.Watcher) {
	// 不存在跳过初始化
	if len(conf) == 0 {
		return
	}

	once.Do(func() {
		ins = &loader{
			mux:  sync.RWMutex{},
			set:  make(map[string][]byte),
			conf: conf,
		}

		// 连接数据库
		initFunc := func() {
			for key, path := range ins.conf {
				if err := ins.initLoader(key, path); err != nil {
					panic("loader init error:" + err.Error())
				}
			}
		}

		watcher("loader", func(value config.Value) {
			ins.mux.Lock()
			defer ins.mux.Unlock()

			if err := value.Scan(&ins.conf); err != nil {
				logger.Instance().Error("loader config watch error", logger.F("err", err))
				return
			}

			// 执行初始化
			initFunc()
		})

		initFunc()
	})
}

// Get 获取指定的加载器
func (c *loader) Get(name string) []byte {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.set[name]
}

// initLoader 初始化加载器
func (c *loader) initLoader(name string, path string) error {
	// 获取文件内容
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	all, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	defer file.Close()

	c.mux.Lock()
	c.set[name] = all
	c.mux.Unlock()
	return nil
}
