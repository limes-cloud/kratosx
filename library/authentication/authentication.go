package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"

	rediswatcher "github.com/billcobbler/casbin-redis-watcher/v2"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/redis/go-redis/v9"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/db"
	rd "github.com/limes-cloud/kratosx/library/redis"
)

type Authentication interface {
	AddWhitelist(path string, method string)
	RemoveWhitelist(path, method string)
	IsWhitelist(path string, method string) bool
	Auth(role, path, method string) bool
	Enforce() *casbin.Enforcer
	IsSkipRole(role string) bool
	SetAuth(req *http.Request, data string)
	ParseAuth(req context.Context, dst any) error
}

type authentication struct {
	conf     *config.Authentication
	enforcer *casbin.Enforcer
	redis    *redis.Client
	roleKey  string
	skipRole map[string]struct{}
	mutex    sync.RWMutex
}

var ins *authentication

const (
	redisKey  = "rbac_authentication"
	authMdKey = "x-md-global-auth"
)

func Instance() Authentication {
	return ins
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

	// 初始化监听器
	w, err := rediswatcher.NewWatcher(
		rdi.Options().Addr,
		rediswatcher.Username(rdi.Options().Username),
		rediswatcher.Password(rdi.Options().Password),
	)
	if err != nil {
		panic("authentication init watcher error:" + err.Error())
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

	// 设置监听器
	_ = object.SetWatcher(w)
	_ = w.SetUpdateCallback(func(s string) {
		_ = object.LoadPolicy()
		log.Errorf("casbin watch load policy")
	})

	ins = &authentication{
		enforcer: object,
		redis:    rdi,
		roleKey:  conf.RoleKey,
		skipRole: make(map[string]struct{}),
		conf:     conf,
	}
	ins.initSkipRole(conf.SkipRole)
	ins.initWhitelist(conf.Whitelist)

	whs := map[string]bool{}
	watcher("authentication.whitelist", func(value config.Value) {
		if err := value.Scan(&whs); err != nil {
			log.Errorf("Authentication Whitelist 配置变更失败：%s", err.Error())
			return
		}
		ins.initWhitelist(whs)
	})

	skips := make([]string, 0)
	watcher("authentication.whitelist", func(value config.Value) {
		if err := value.Scan(&skips); err != nil {
			log.Errorf("Authentication SkipRole 配置变更失败：%s", err.Error())
			return
		}

		ins.initSkipRole(skips)
	})
}

func (a *authentication) initWhitelist(whs map[string]bool) {
	for path, is := range whs {
		arr := strings.Split(path, ":")
		if len(arr) != 2 {
			continue
		}
		if is {
			ins.AddWhitelist(arr[1], arr[0])
		} else {
			ins.RemoveWhitelist(arr[1], arr[0])
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

func (a *authentication) SetAuth(req *http.Request, data string) {
	if data == "" {
		return
	}
	req.Header.Set(authMdKey, data)
}

func (a *authentication) ParseAuth(ctx context.Context, dst any) error {
	if md, ok := metadata.FromServerContext(ctx); ok {
		body := md.Get(authMdKey)
		if err := json.Unmarshal([]byte(body), dst); err != nil {
			return errors.New("auth info format error:" + err.Error())
		}
		return nil
	}
	return errors.New("not exist auth info")
}

func (a *authentication) IsSkipRole(role string) bool {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	_, is := a.skipRole[role]
	return is
}

func (a *authentication) path(path, method string) string {
	return fmt.Sprintf("%s:%s", path, method)
}

func (a *authentication) AddWhitelist(path string, method string) {
	a.redis.HSet(context.Background(), redisKey, a.path(path, method), 1)
}

func (a *authentication) RemoveWhitelist(path, method string) {
	a.redis.HDel(context.Background(), redisKey, a.path(path, method))
}

func (a *authentication) IsWhitelist(path, method string) bool {
	if !a.conf.EnableGrpc && method == "GRPC" {
		return true
	}
	is, _ := a.redis.HGet(context.Background(), redisKey, a.path(path, method)).Bool()
	return is
}

func (a *authentication) Auth(role, path, method string) bool {
	if a.IsWhitelist(path, method) {
		return true
	}

	// 进行鉴权
	is, _ := a.enforcer.Enforce(role, path, method)
	return is
}

func (a *authentication) Enforce() *casbin.Enforcer {
	return a.enforcer
}
