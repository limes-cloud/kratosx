package kratosx

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library"
	"github.com/limes-cloud/kratosx/library/logger"
	"github.com/limes-cloud/kratosx/middleware"
	"os"
)

const (
	AppName    = "APP_NAME"
	AppVersion = "APP_VERSION"
)

var (
	name    = os.Getenv(AppName)
	version = os.Getenv(AppVersion)
	id, _   = os.Hostname()
)

func New(opts ...Option) *kratos.App {
	o := &options{
		config: config.New(file.NewSource("config/config.yaml")),
		loggerFields: logger.LogField{
			"id":      id,
			"name":    name,
			"version": version,
			"trace":   tracing.TraceID(),
			"span":    tracing.SpanID(),
		},
	}

	for _, opt := range opts {
		opt(o)
	}

	// 必须注册服务
	if o.regSrvFn == nil {
		panic("must register server")
	}

	// 加载配置
	if err := o.config.Load(); err != nil {
		panic(err)
	}

	// 插件初始化
	library.Init(o.config, o.loggerFields)

	// 获取中间件
	mds := middleware.New(o.config)

	// 注册服务
	hs := httpServer(o.config.App().Server.Http, mds)
	gs := grpcServer(o.config.App().Server.Grpc, mds)
	o.regSrvFn(o.config, hs, gs)

	defOpts := []kratos.Option{
		kratos.ID(id),
		kratos.Name(name),
		kratos.Version(version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger.Instance()),
		kratos.Server(hs, gs),
	}

	defOpts = append(defOpts, o.kOpts...)

	return kratos.New(
		defOpts...,
	)
}
