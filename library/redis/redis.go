package redis

import (
	"context"
	kratosConfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	goRedis "github.com/go-redis/redis/v8"
	"github.com/limes-cloud/kratosx/config"
	"sync"
)

type Redis interface {
	// Get 获取指定名称的redis实例，如果实例不存在则会nil
	Get(name ...string) *goRedis.Client
}

type redis struct {
	mu  sync.RWMutex
	set map[string]*goRedis.Client
	key string
}

var instance *redis

func Instance() Redis {
	return instance
}

func Init(cfs map[string]*config.Redis, watcher config.Watcher) {
	if len(cfs) == 0 {
		return
	}

	instance = &redis{
		mu:  sync.RWMutex{},
		set: make(map[string]*goRedis.Client),
	}

	for key, conf := range cfs {
		if err := instance.initFactory(key, conf); err != nil {
			panic("init redis error :" + err.Error())
		}

		watcher("redis."+key, func(value kratosConfig.Value) {
			if err := value.Scan(conf); err != nil {
				log.Errorf("Redis配置变更失败：%s", err.Error())
				return
			}
			if err := instance.initFactory(key, conf); err != nil {
				log.Errorf("Redis变更重载失败：%s", err.Error())
			}
		})
	}

	if len(instance.set) != 1 {
		instance.key = ""
	}
}

func (r *redis) initFactory(name string, conf *config.Redis) error {
	if !conf.Enable {
		r.delete(name)
		return nil
	}

	// 连接主数据库
	client := goRedis.NewClient(&goRedis.Options{
		Addr:     conf.Host,
		Username: conf.Username,
		Password: conf.Password,
	})
	if err := client.Ping(context.TODO()).Err(); err != nil {
		return err
	}

	r.mu.Lock()
	r.set[name] = client
	r.key = name
	r.mu.Unlock()
	return nil
}

func (r *redis) delete(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.set, name)
}

func (r *redis) Get(name ...string) *goRedis.Client {
	if r == nil {
		return nil
	}

	if r.key == "" && len(name) == 0 {
		return nil
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	key := r.key
	if len(name) != 0 {
		key = name[0]
	}
	return r.set[key]
}
