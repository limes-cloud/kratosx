package kratosx

import (
	"github.com/go-kratos/kratos/v2"
	kratosConfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/logger"
)

type Option func(o *options)

type RegistrarServerFn func(config config.Config, hs *http.Server, gs *grpc.Server)

type options struct {
	regSrvFn       RegistrarServerFn
	loggerFields   logger.LogField
	config         config.Config
	kOpts          []kratos.Option
	httpSrvOptions []http.ServerOption
	grpcSrvOptions []grpc.ServerOption
	midOpts        []middleware.Middleware
}

// RegistrarServer 服务注册
func RegistrarServer(fn RegistrarServerFn) Option {
	return func(o *options) {
		o.regSrvFn = fn
	}
}

// LoggerWith 自定义字段
func LoggerWith(fields logger.LogField) Option {
	// var fs []any
	// for key, val := range fields {
	//	fs = append(fs, key, val)
	// }
	return func(o *options) { o.loggerFields = fields }
}

// Config 配置接入
func Config(source kratosConfig.Source) Option {
	return func(o *options) {
		o.config = config.New(source)
	}
}

// Options kratos option
func Options(opts ...kratos.Option) Option {
	return func(o *options) {
		o.kOpts = opts
	}
}

// HttpServerOptions http server option
func HttpServerOptions(opts ...http.ServerOption) Option {
	return func(o *options) {
		o.httpSrvOptions = opts
	}
}

// GrpcServerOptions grpc server option
func GrpcServerOptions(opts ...grpc.ServerOption) Option {
	return func(o *options) {
		o.grpcSrvOptions = opts
	}
}

// MiddlewareOptions middleware option
func MiddlewareOptions(opts ...middleware.Middleware) Option {
	return func(o *options) {
		o.midOpts = opts
	}
}
