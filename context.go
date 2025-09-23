package kratosx

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/limes-cloud/kratosx/library/tasker"
	"github.com/limes-cloud/kratosx/pkg/ua"
	"time"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/metadata"
	md "github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/limes-cloud/kratosx/library/env"
	"github.com/limes-cloud/kratosx/library/request"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/captcha"
	"github.com/limes-cloud/kratosx/library/client"
	"github.com/limes-cloud/kratosx/library/db"
	"github.com/limes-cloud/kratosx/library/email"
	"github.com/limes-cloud/kratosx/library/ip"
	"github.com/limes-cloud/kratosx/library/jwt"
	"github.com/limes-cloud/kratosx/library/loader"
	"github.com/limes-cloud/kratosx/library/logger"
	"github.com/limes-cloud/kratosx/library/pool"
	"github.com/limes-cloud/kratosx/library/prometheus"
	rd "github.com/limes-cloud/kratosx/library/redis"
)

type Context interface {
	// Env 获取环境变量
	Env() env.Env

	// Logger 获取链路日志器
	Logger() logger.Logger

	// DB 获取数据库
	DB(name ...string) *gorm.DB

	// Database 获取系统的数据库
	Database() db.DB

	// Transaction 获取DB事务
	Transaction(fn func(ctx Context) error, name ...string) error

	// Redis 获取Redis客户端
	Redis(name ...string) *redis.Client

	// Pool 获取全局协程池
	Pool() pool.Pool

	// Loader 获取文件加载器
	Loader(name string) []byte

	// Email 获取邮件服务
	Email() email.Email

	// ClientIP 获取客户端IP库
	ClientIP() string

	// Captcha 获取验证码服务
	Captcha() captcha.PCaptcha

	// JWT 获取Jwt服务
	JWT() jwt.Jwt

	// Token 获取JwtToken
	Token() string

	// Ctx 获取行下文ctx
	Ctx() context.Context

	// GetMetadata 获取元数据
	GetMetadata(string) string

	// SetMetadata 设置元数据
	SetMetadata(key, value string)

	// Request 获取http请求工具
	Request() request.Request

	// GrpcConn 获取grpc连接
	GrpcConn(srvName string) (*grpc.ClientConn, error)

	// BeforeStart 注册关闭前执行函数
	BeforeStart(name string, fn func())

	// AfterStart 注册关闭后执行函数
	AfterStart(name string, fn func())

	// BeforeStop 注册关闭前执行函数
	BeforeStop(name string, fn func())

	// AfterStop 注册关闭后执行函数
	AfterStop(name string, fn func())

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

// Database 获取数据库
func (c *ctx) Database() db.DB {
	return db.Instance()
}

// Ctx 获取行下文ctx
func (c *ctx) Ctx() context.Context {
	return c.Context
}

// Logger 获取链路日志器
func (c *ctx) Logger() logger.Logger {
	return logger.Instance()
}

// Transaction 开启层级事物
func (c *ctx) Transaction(fn func(ctx Context) error, name ...string) error {
	dbi := db.Instance()
	if _, ok := c.Value(dbi.TxKey(name...)).(*gorm.DB); ok {
		return fn(MustContext(c.Ctx()))
	}

	return dbi.Get(name...).WithContext(c.Ctx()).Transaction(func(tx *gorm.DB) error {
		// nolint
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
	// 追加数据库名称
	return dbi.Get(name...).WithContext(c.Ctx())
}

// Prometheus 监控
func (c *ctx) Prometheus() prometheus.Prometheus {
	return prometheus.Instance()
}

// Redis 获取缓存实例
func (c *ctx) Redis(name ...string) *redis.Client {
	return rd.Instance().Get(name...)
}

// Pool 获取并发池实例
func (c *ctx) Pool() pool.Pool {
	return pool.Instance()
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
func (c *ctx) Captcha() captcha.PCaptcha {
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

// Request 带链路日志的请求工具
func (c *ctx) Request() request.Request {
	return request.Instance(c)
}

// Token 获取令牌验证器
func (c *ctx) Token() string {
	return jwt.Instance().GetToken(c.Context)
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

// UserAgent 用户请求头
func (c *ctx) UserAgent() ua.UserAgent {
	header, ok := transport.FromServerContext(c.Context)
	if !ok {
		return ua.UserAgent{}
	}
	return ua.Parse(header.RequestHeader().Get("User-Agent"))
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
	return MustContext(context.WithoutCancel(c.Context))
}

// BeforeStop 注册服务关闭回调
func (c *ctx) BeforeStop(name string, fn func()) {
	tasker.Instance().BeforeStop(name, fn)
}

// AfterStop 注册服务关闭回调
func (c *ctx) AfterStop(name string, fn func()) {
	tasker.Instance().AfterStop(name, fn)
}

// BeforeStart 注册服务启动回调
func (c *ctx) BeforeStart(name string, fn func()) {
	tasker.Instance().BeforeStart(name, fn)
}

// AfterStart 注册服务启动回调
func (c *ctx) AfterStart(name string, fn func()) {
	tasker.Instance().AfterStart(name, fn)
}

// Trace 获取trace id
func (c *ctx) Trace() string {
	t, _ := tracing.TraceID()(c.Context).(string)
	return t
}

// Span 获取span id
func (c *ctx) Span() string {
	t, _ := tracing.SpanID()(c.Context).(string)
	return t
}

// Env 获取配置环境
func (c *ctx) Env() env.Env {
	return env.Instance()
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

type ContextOption struct {
	Trace      string
	Span       string
	SkipDBHook bool
}

type ContextOptionFunc func(*ContextOption)

// WithTrace 主动设置trace信息
func WithTrace(trace string, span string) ContextOptionFunc {
	return func(o *ContextOption) {
		o.Trace = trace
		o.Span = span
	}
}

func WithSkipDBHook() ContextOptionFunc {
	return func(o *ContextOption) {
		o.SkipDBHook = true
	}
}

// MustContext returns the Transport value stored in ctx, if any.
func MustContext(c context.Context, opts ...ContextOptionFunc) Context {
	o := &ContextOption{}
	for _, opt := range opts {
		opt(o)
	}

	if o.Trace != "" {
		c = withTraceContext(c, o.Trace, o.Span)
	}

	if o.SkipDBHook {
		c = context.WithValue(c, db.SkipHookKey, true)
	}

	app, _ := kratos.FromContext(c)
	return &ctx{
		Context: c,
		AppInfo: app,
	}
}

func withTraceContext(ctx context.Context, t string, s string) context.Context {
	tid, err := trace.TraceIDFromHex(t)
	if err != nil {
		return ctx
	}

	sid, err := trace.SpanIDFromHex(s)
	if err != nil {
		return ctx
	}

	// 创建一个新的SpanContext，使用已知的Trace ID和Span ID
	sc := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: tid,
		SpanID:  sid,
		Remote:  true,
	})

	// 创建一个包含新SpanContext的context
	return trace.ContextWithSpanContext(ctx, sc)
}
