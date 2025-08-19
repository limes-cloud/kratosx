package redis

import (
	"context"
	"sync"

	goredis "github.com/redis/go-redis/v9"

	"github.com/limes-cloud/kratosx/config"
)

type Redis interface {
	// Get 获取指定名称的redis实例，如果实例不存在则会nil
	Get(name ...string) *goredis.Client
}

type redis struct {
	set map[string]*goredis.Client
	key string
}

var (
	// redis 实例
	ins *redis

	once sync.Once
)

// Instance 获取redis实例
func Instance() Redis {
	return ins
}

// Init 初始化redis连接池
func Init(cfs []*config.Redis) {
	if len(cfs) == 0 {
		return
	}

	once.Do(func() {
		ins = &redis{
			set: make(map[string]*goredis.Client),
		}

		for ind, conf := range cfs {
			if err := ins.initRedis(conf); err != nil {
				panic("init redis error :" + err.Error())
			}
			if ind == 0 {
				ins.key = conf.Name
			}
		}
	})
}

// initRedis 初始化redis
func (r *redis) initRedis(conf *config.Redis) error {
	if !conf.Enable {
		return nil
	}
	// 连接主数据库
	client := goredis.NewClient(&goredis.Options{
		Addr:     conf.Host,
		Username: conf.Username,
		Password: conf.Password,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return err
	}

	r.set[conf.Name] = client
	return nil
}

// Get 获取指定名称的redis实例，如果实例不存在则会nil
// 当只存在一个redis配置时，name默认可不传入
func (r *redis) Get(name ...string) *goredis.Client {
	if r == nil {
		return nil
	}

	if r.key == "" && len(name) == 0 {
		return nil
	}

	key := r.key
	if len(name) != 0 {
		key = name[0]
	}
	return r.set[key]
}
