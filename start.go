package kratosx

import (
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library"
	"github.com/limes-cloud/kratosx/library/logger"
	"github.com/limes-cloud/kratosx/library/pprof"
	"github.com/limes-cloud/kratosx/library/registry"
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
		config: config.New(file.NewSource("internal/conf/conf.yaml")),
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
	defOpts := []kratos.Option{
		kratos.ID(o.config.App().ID),
		kratos.Name(o.config.App().Name),
		kratos.Version(o.config.App().Version),
		kratos.Metadata(map[string]string{}),
	}

	// 必注册服务
	if o.regSrvFn != nil {
		gsOpts, hsOpts := serverOptions(o.config)
		srv := o.config.App().Server
		gs := grpcServer(srv.Grpc, srv.Count, gsOpts)
		hs := httpServer(srv.Http, srv.Count, hsOpts)
		o.regSrvFn(o.config, hs, gs)

		var srvList []transport.Server
		if srv.Http != nil {
			srvList = append(srvList, hs)
			// 监控
			if o.config.App().Metrics {
				hs.Handle("/metrics", promhttp.Handler())
			}
			// pprof
			if o.config.App().Server.Http.Pprof != nil {
				pprof.PprofServer(o.config.App().Server.Http.Pprof, hs)
			}
		}
		if srv.Grpc != nil {
			srvList = append(srvList, gs)
		}
		defOpts = append(defOpts, kratos.Server(srvList...))

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
