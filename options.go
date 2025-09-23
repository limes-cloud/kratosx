package kratosx

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	kconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library"
	"github.com/limes-cloud/kratosx/library/env"
	kmid "github.com/limes-cloud/kratosx/middleware"
)

type Option func(o *options)

type RegistrarServerFn func(hs *http.Server, gs *grpc.Server)

type options struct {
	// 监听配置
	watch func(config.Watcher)
	// 配置服务
	config config.Config
	// 配置钩子
	configHooks func(*config.App)
	// 注册函数
	regSrvFn RegistrarServerFn
	// 组件配置
	libOpts []library.Option
	// kratos 配置
	kOpts []kratos.Option
	// http 服务
	httpSrvOptions []http.ServerOption
	// grpc 服务
	grpcSrvOptions []grpc.ServerOption
	// 中间件
	midOpts []middleware.Middleware
	// 中间件hook
	midHooks kmid.MidHook
}

func WithValidateErrHook(hook func(ctx context.Context, err error) error) Option {
	return func(o *options) {
		o.midHooks.ValidateErrHook = hook
	}
}

// WithRegistrarServer 服务注册
func WithRegistrarServer(fn RegistrarServerFn) Option {
	return func(o *options) {
		o.regSrvFn = fn
	}
}

// WithConfigSource 配置接入
func WithConfigSource(source ...kconfig.Source) Option {
	return func(o *options) {
		o.config = config.New(source...)
	}
}

// WithConfigWatch 配置监听
func WithConfigWatch(watch func(config.Watcher)) Option {
	return func(o *options) {
		o.watch = watch
	}
}

// WithKratosOptions kratos option
func WithKratosOptions(opts ...kratos.Option) Option {
	return func(o *options) {
		o.kOpts = opts
	}
}

// WithLibraryOptions 组件服务option
func WithLibraryOptions(opts ...library.Option) Option {
	return func(o *options) {
		o.libOpts = opts
	}
}

// WithHttpServerOptions http server option
func WithHttpServerOptions(opts ...http.ServerOption) Option {
	return func(o *options) {
		o.httpSrvOptions = opts
	}
}

// WithGrpcServerOptions grpc server option
func WithGrpcServerOptions(opts ...grpc.ServerOption) Option {
	return func(o *options) {
		o.grpcSrvOptions = opts
	}
}

// WithMiddleware middleware option
func WithMiddleware(opts ...middleware.Middleware) Option {
	return func(o *options) {
		o.midOpts = opts
	}
}

// WithUnitTest 单元测试
func WithUnitTest() Option {
	return func(o *options) {
		env.Instance().SetRunUnitTest()
	}
}

// WithHookSystemConfig 劫持修改系统配置
func WithHookSystemConfig(fn func(*config.App)) Option {
	return func(o *options) {
		o.configHooks = fn
	}
}
