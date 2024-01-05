package kratosx

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	md "github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/authentication"
	"github.com/limes-cloud/kratosx/library/captcha"
	"github.com/limes-cloud/kratosx/library/db"
	"github.com/limes-cloud/kratosx/library/email"
	"github.com/limes-cloud/kratosx/library/ip"
	"github.com/limes-cloud/kratosx/library/jwt"
	"github.com/limes-cloud/kratosx/library/loader"
	"github.com/limes-cloud/kratosx/library/logger"
	"github.com/limes-cloud/kratosx/library/pool"
	rd "github.com/limes-cloud/kratosx/library/redis"
)

type Context interface {
	Env() string
	Logger() *log.Helper
	DB(name ...string) *gorm.DB
	Go(runner pool.Runner) error
	Loader(name string) []byte
	Email() email.Email
	ClientIP() string
	Captcha() captcha.Captcha
	JWT() jwt.Jwt
	Token() string
	Authentication() authentication.Authentication
	Ctx() context.Context
	GetMetadata(string) string
	SetMetadata(key, value string)

	ID() string
	Name() string
	Version() string
	Metadata() map[string]string
	Config() config.Config
	Endpoint() []string

	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}

type ctx struct {
	context.Context
	kratos.AppInfo
}

// MustContext returns the Transport value stored in ctx, if any.
func MustContext(c context.Context) Context {
	app, _ := kratos.FromContext(c)
	return &ctx{
		Context: c,
		AppInfo: app,
	}
}

func (c *ctx) Ctx() context.Context {
	return c.Context
}

// Logger 获取链路日志器
func (c *ctx) Logger() *log.Helper {
	return logger.Helper().WithContext(c.Context)
}

// DB 数据库实例
func (c *ctx) DB(name ...string) *gorm.DB {
	return db.Instance().Get(name...).WithContext(c.Context)
}

// Redis 获取缓存实例
func (c *ctx) Redis(name ...string) *redis.Client {
	return rd.Instance().Get(name...).WithContext(c.Context)
}

// Go 获取并发池实例
func (c *ctx) Go(runner pool.Runner) error {
	return pool.Instance().Go(runner)
}

// Loader 获加载器实例
func (c *ctx) Loader(name string) []byte {
	return loader.Instance().Get(name)
}

// Email 获取邮箱实例
func (c *ctx) Email() email.Email {
	return email.Instance()
}

// Captcha 获取图形验证器
func (c *ctx) Captcha() captcha.Captcha {
	return captcha.Instance()
}

// ClientIP 获取客户端IP地址
func (c *ctx) ClientIP() string {
	return ip.ClientIP(c.Context)
}

// JWT 获取令牌验证器
func (c *ctx) JWT() jwt.Jwt {
	return jwt.Instance()
}

// Token 获取令牌验证器
func (c *ctx) Token() string {
	return jwt.Instance().GetToken(c.Context)
}

// Authentication 获取权限验证器
func (c *ctx) Authentication() authentication.Authentication {
	return authentication.Instance()
}

// GetMetadata 获取元数据信息
func (c *ctx) GetMetadata(key string) string {
	if values, ok := metadata.FromServerContext(c.Context); ok {
		return values.Get(key)
	}
	return ""
}

// SetMetadata 设置元数据信息
func (c *ctx) SetMetadata(key, value string) {
	c.Context = md.AppendToClientContext(context.Background(), key, value)
}

// Config 获取配置对象
func (c *ctx) Config() config.Config {
	return config.Instance()
}

// Env 获取配置环境
func (c *ctx) Env() string {
	return c.Config().App().Env
}

func (c *ctx) Deadline() (deadline time.Time, ok bool) {
	return c.Context.Deadline()
}

func (c *ctx) Done() <-chan struct{} {
	return c.Context.Done()
}

func (c *ctx) Err() error {
	return c.Context.Err()
}

func (c *ctx) Value(key any) any {
	return c.Context.Value(key)
}
