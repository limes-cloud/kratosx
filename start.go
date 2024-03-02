package kratosx

import (
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library"
	"github.com/limes-cloud/kratosx/library/logger"
	"github.com/limes-cloud/kratosx/library/registry"
	"github.com/limes-cloud/kratosx/middleware"
)

const (
	AppName    = "APP_NAME"
	AppVersion = "APP_VERSION"
)

var (
	envName    = os.Getenv(AppName)
	envVersion = os.Getenv(AppVersion)
	id, _      = os.Hostname()
)

func New(opts ...Option) *kratos.App {
	o := &options{
		config: config.New(file.NewSource("internal/config/config.yaml")),
	}

	for _, opt := range opts {
		opt(o)
	}

	// 加载配置
	if err := o.config.Load(); err != nil {
		panic(err)
	}

	// 初始化服务信息
	if o.config.App().Name == "" {
		o.config.SetAppInfo(id, envName, envVersion)
	}

	// 插件初始化
	if o.loggerFields == nil {
		o.loggerFields = logger.LogField{
			"id":      o.config.App().ID,
			"name":    o.config.App().Name,
			"version": o.config.App().Version,
			"trace":   tracing.TraceID(),
			"span":    tracing.SpanID(),
		}
	}

	library.Init(o.config, o.loggerFields)

	// 获取中间件
	mds := middleware.New(o.config)
	defOpts := []kratos.Option{
		kratos.ID(o.config.App().ID),
		kratos.Name(o.config.App().Name),
		kratos.Version(o.config.App().Version),
		kratos.Metadata(map[string]string{}),
	}

	// 必注册服务
	if o.regSrvFn != nil {
		srv := o.config.App().Server
		hs := httpServer(srv.Http, srv.Count, mds)
		gs := grpcServer(srv.Grpc, srv.Count, mds)
		o.regSrvFn(o.config, hs, gs)
		defOpts = append(defOpts, kratos.Server(hs, gs))

		if srv.Registry != nil {
			reg, err := registry.Create(*srv.Registry)
			if err != nil {
				panic(err)
			}
			defOpts = append(defOpts, kratos.Registrar(reg))
		}
	}

	// 日志
	if o.config.App().Log != nil {
		defOpts = append(defOpts, kratos.Logger(logger.Instance()))
	}

	defOpts = append(defOpts, o.kOpts...)

	return kratos.New(
		defOpts...,
	)
}
