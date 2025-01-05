package email

import (
	"io"
	"os"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/limes-cloud/kratosx/config"
)

type Email interface {
	Template(name string) Sender
}

type email struct {
	mu   sync.RWMutex
	set  map[string][]byte
	conf *config.Email
}

var instance *email

// Instance 获取email对象实例
func Instance() Email {
	return instance
}

func Init(conf *config.Email, watcher config.Watcher) {
	// 没有模板则跳过初始化
	if conf == nil || len(conf.Template) == 0 {
		return
	}

	instance = &email{
		mu:   sync.RWMutex{},
		set:  make(map[string][]byte),
		conf: conf,
	}

	// 遍历初始化模板
	for key, tpc := range conf.Template {
		if err := instance.initFactory(key, tpc); err != nil {
			panic("Email 初始化失败 :" + err.Error())
		}

		watcher("email.template."+key, func(value config.Value) {
			if err := value.Scan(&tpc); err != nil {
				log.Errorf("Email Template 配置变更失败：%s", err.Error())
				return
			}
			if err := instance.initFactory(key, tpc); err != nil {
				log.Errorf("Email Template 变更重载失败:%s", err.Error())
			}
		})
	}
}

func (c *email) Template(name string) Sender {
	return &sender{
		stp:  name,
		set:  c.set,
		conf: c.conf,
	}
}

func (c *email) initFactory(name string, et config.EmailTemplate) error {
	if et.Enable != nil && !*et.Enable {
		c.delete(name)
		return nil
	}

	// 获取文件内容
	file, err := os.Open(et.Path)
	if err != nil {
		return err
	}
	all, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	defer file.Close()

	c.mu.Lock()
	c.set[name] = all
	c.mu.Unlock()
	return nil
}

func (c *email) delete(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.set, name)
}
