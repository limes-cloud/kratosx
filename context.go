package kratosx

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	md "github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/authentication"
	"github.com/limes-cloud/kratosx/library/captcha"
	"github.com/limes-cloud/kratosx/library/client"
	"github.com/limes-cloud/kratosx/library/db"
	"github.com/limes-cloud/kratosx/library/email"
	"github.com/limes-cloud/kratosx/library/http"
	"github.com/limes-cloud/kratosx/library/ip"
	"github.com/limes-cloud/kratosx/library/jwt"
	"github.com/limes-cloud/kratosx/library/loader"
	"github.com/limes-cloud/kratosx/library/logger"
	"github.com/limes-cloud/kratosx/library/pool"
	"github.com/limes-cloud/kratosx/library/prometheus"
	rd "github.com/limes-cloud/kratosx/library/redis"
	"github.com/limes-cloud/kratosx/library/stop"
)

type Context interface {
	// Env 获取环境变量
	Env() string

	// Logger 获取链路日志器
	Logger() *log.Helper

	// DB 获取数据库
	DB(name ...string) *gorm.DB

	// Transaction 获取DB事务
	Transaction(fn func(ctx Context) error, name ...string) error

	// Redis 获取Redis客户端
	Redis(name ...string) *redis.Client

	// Go 获取全局协程池
	Go(runner pool.Runner) error

	// Loader 获取文件加载器
	Loader(name string) []byte

	// Email 获取邮件服务
	Email() email.Email

	// ClientIP 获取客户端IP库
	ClientIP() string

	// Captcha 获取验证码服务
	Captcha() captcha.Captcha

	// JWT 获取Jwt服务
	JWT() jwt.Jwt

	// Token 获取JwtToken
	Token() string

	// Authentication 获取认证服务
	Authentication() authentication.Authentication

	// Ctx 获取行下文ctx
	Ctx() context.Context

	// GetMetadata 获取元数据
	GetMetadata(string) string

	// SetMetadata 设置元数据
	SetMetadata(key, value string)

	// WaitRunner 获取等待协程池
	WaitRunner(opts ...pool.WaitRunnerOptionFunc) pool.WaitRunner

	// Http 获取http请求
	Http() http.Request

	// GrpcConn 获取grpc连接
	GrpcConn(srvName string) (*grpc.ClientConn, error)

	// RegisterBeforeStop 注册关闭前执行函数
	RegisterBeforeStop(name string, fn func())

	// RegisterAfterStop 注册关闭后执行函数
	RegisterAfterStop(name string, fn func())

	// ID 获取ID
	ID() string

	// Name 获取服务名称
	Name() string

	// Version 获取服务版本
	Version() string

	// Metadata 获取元数据
	Metadata() map[string]string

	// Config 获取配置
	Config() config.Config

	// Endpoint 获取服务地址
	Endpoint() []string

	// Clone 获取行下文ctx
	Clone() Context

	// Deadline 获取截止时间
	Deadline() (deadline time.Time, ok bool)

	// Done 获取关闭通道
	Done() <-chan struct{}

	// Err 获取错误
	Err() error

	// Value 获取值
	Value(key any) any
}

type ctx struct {
	context.Context
	kratos.AppInfo
}

type cloneCtx struct {
	child  context.Context
	parent context.Context
}

// MustContext returns the Transport value stored in ctx, if any.
func MustContext(c context.Context) Context {
	app, _ := kratos.FromContext(c)
	return &ctx{
		Context: c,
		AppInfo: app,
	}
}

// Ctx 获取行下文ctx
func (c *ctx) Ctx() context.Context {
	return c.Context
}

// Logger 获取链路日志器
func (c *ctx) Logger() *log.Helper {
	if !c.Config().IsInit() {
		return log.NewHelper(log.DefaultLogger)
	}
	return logger.Helper().WithContext(c)
}

func (c *ctx) Transaction(fn func(ctx Context) error, name ...string) error {
	dbi := db.Instance()
	if _, ok := c.Value(dbi.TxKey(name...)).(*gorm.DB); ok {
		return fn(MustContext(c.Ctx()))
	}

	return dbi.Get(name...).WithContext(c.Ctx()).Transaction(func(tx *gorm.DB) error {
		cc := context.WithValue(c.Ctx(), dbi.TxKey(name...), tx)
		return fn(MustContext(cc))
	})
}

// DB 数据库实例
func (c *ctx) DB(name ...string) *gorm.DB {
	dbi := db.Instance()
	tx, ok := c.Value(dbi.TxKey(name...)).(*gorm.DB)
	if ok {
		return tx
	}
	return dbi.Get(name...).WithContext(c.Ctx())
}

// Prometheus 监控
func (c *ctx) Prometheus() prometheus.Prometheus {
	return c.Prometheus()
}

// Redis 获取缓存实例
func (c *ctx) Redis(name ...string) *redis.Client {
	return rd.Instance().Get(name...).WithContext(c.Context)
}

// Go 获取并发池实例
func (c *ctx) Go(runner pool.Runner) error {
	return pool.Instance().Go(runner)
}

// WaitRunner 获取并发池等待实例
func (c *ctx) WaitRunner(opts ...pool.WaitRunnerOptionFunc) pool.WaitRunner {
	return pool.NewWaitRunner(opts...)
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

// Http 带链路日志的请求工具
func (c *ctx) Http() http.Request {
	if !c.Config().IsInit() || c.Config().App().Http == nil {
		return http.NewDefault(c.Logger())
	}
	return http.New(c.Config().App().Http, c.Logger())
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

// GrpcConn 获取grpc 连接具柄
func (c *ctx) GrpcConn(srvName string) (*grpc.ClientConn, error) {
	cli := client.Get(srvName)
	if cli == nil {
		return nil, errors.New("not exist server " + srvName)
	}
	return cli.Conn(c.Ctx())
}

// Clone 克隆上下文，上下文中继承了被克隆的值，但是并不会收到上下文的timeout和cancel信号。
func (c *ctx) Clone() Context {
	return MustContext(&cloneCtx{
		child:  context.Background(),
		parent: c.Context,
	})
}

// RegisterBeforeStop 注册服务关闭回调
func (c *ctx) RegisterBeforeStop(name string, fn func()) {
	stop.Instance().RegisterBefore(name, fn)
}

// RegisterAfterStop 注册服务关闭回调
func (c *ctx) RegisterAfterStop(name string, fn func()) {
	stop.Instance().RegisterAfter(name, fn)
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

func (c *cloneCtx) Ctx() context.Context {
	return c
}

func (c *cloneCtx) Deadline() (deadline time.Time, ok bool) {
	return c.child.Deadline()
}

func (c *cloneCtx) Done() <-chan struct{} {
	return c.child.Done()
}

func (c *cloneCtx) Err() error {
	return c.child.Err()
}

func (c *cloneCtx) Value(key any) any {
	return c.parent.Value(key)
}
