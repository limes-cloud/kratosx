package authentication

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter "github.com/casbin/gorm-adapter/v3"
	kratosConfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/db"
	"github.com/limes-cloud/kratosx/library/jwt"
	rd "github.com/limes-cloud/kratosx/library/redis"
	"strings"
	"sync"
)

type Authentication interface {
	AddWhitelist(path string, method string)
	RemoveWhitelist(path, method string)
	IsWhitelist(path string, method string) bool
	Auth(role, path, method string) bool
	GetRole(ctx context.Context) (string, error)
	Enforce() *casbin.Enforcer
	IsSkipRole(role string) bool
}

type authentication struct {
	enforcer *casbin.Enforcer
	redis    *redis.Client
	roleKey  string
	skipRole map[string]struct{}
	mutex    sync.RWMutex
	prefix   string
}

var instance *authentication

const redisKey = "rbac_authentication"

func Instance() Authentication {
	return instance
}

func Init(conf *config.Authentication, watcher config.Watcher) {
	if conf == nil {
		return
	}

	dbi := db.Instance().Get(conf.DB)
	rdi := rd.Instance().Get(conf.Redis)

	if dbi == nil {
		panic("authentication init error not exist database " + conf.DB)
	}

	if rdi == nil {
		panic("authentication init error not exist redis " + conf.DB)
	}

	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act")

	a, err := adapter.NewAdapterByDB(db.Instance().Get(conf.DB))
	if err != nil {
		panic("authentication init error:" + err.Error())
	}
	object, _ := casbin.NewEnforcer(m, a)
	if err = object.LoadPolicy(); err != nil {
		panic("authentication init error:" + err.Error())
	}

	instance = &authentication{
		enforcer: object,
		redis:    rdi,
		roleKey:  conf.RoleKey,
		skipRole: make(map[string]struct{}),
		prefix:   conf.Prefix,
	}
	instance.initSkipRole(conf.SkipRole)
	instance.initWhitelist(conf.Whitelist)

	whs := map[string]bool{}
	watcher("authentication.whitelist", func(value kratosConfig.Value) {
		if err := value.Scan(&whs); err != nil {
			log.Errorf("Authentication Whitelist 配置变更失败：%s", err.Error())
			return
		}
		instance.initWhitelist(whs)
	})

	skips := make([]string, 0)
	watcher("authentication.whitelist", func(value kratosConfig.Value) {
		if err := value.Scan(&skips); err != nil {
			log.Errorf("Authentication SkipRole 配置变更失败：%s", err.Error())
			return
		}
		instance.initSkipRole(skips)
	})

}

func (a *authentication) initWhitelist(whs map[string]bool) {
	for path, is := range whs {
		arr := strings.Split(path, ":")
		if len(arr) != 2 {
			continue
		}
		if is {
			instance.AddWhitelist(arr[1], arr[0])
		} else {
			instance.RemoveWhitelist(arr[1], arr[0])
		}
	}
}

func (a *authentication) initSkipRole(skips []string) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	for _, role := range skips {
		a.skipRole[role] = struct{}{}
	}
}

func (a *authentication) IsSkipRole(role string) bool {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	_, is := a.skipRole[role]
	return is
}

func (a *authentication) path(path, method string) string {
	return fmt.Sprintf("%s:%s", a.prefix+path, method)
}

func (a *authentication) AddWhitelist(path string, method string) {
	a.redis.HSet(context.Background(), redisKey, a.path(path, method), 1)
}

func (a *authentication) RemoveWhitelist(path, method string) {
	a.redis.HDel(context.Background(), redisKey, a.path(path, method))
}

func (a *authentication) IsWhitelist(path, method string) bool {
	is, _ := a.redis.HGet(context.Background(), redisKey, a.path(path, method)).Bool()
	return is
}

func (a *authentication) Auth(role, path, method string) bool {
	if a.IsWhitelist(path, method) {
		return true
	}

	// 进行鉴权
	is, _ := a.enforcer.Enforce(role, a.prefix+path, method)
	return is
}

func (a *authentication) GetRole(ctx context.Context) (string, error) {
	claims, _ := jwt.Instance().ParseMapClaims(ctx)
	if claims == nil {
		return "", nil
	}

	role, is := claims[a.roleKey].(string)
	if !is {
		return "", fmt.Errorf("not exist role field %v", a.roleKey)
	}
	return role, nil
}

func (a *authentication) Enforce() *casbin.Enforcer {
	return a.enforcer
}
