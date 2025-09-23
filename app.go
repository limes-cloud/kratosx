package kratosx

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/limes-cloud/kratosx/library/tasker"
	"net"
	"time"

	"github.com/go-kratos/kratos/v2"
	kenv "github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/encoding/json"
	kmid "github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library"
	"github.com/limes-cloud/kratosx/library/env"
	"github.com/limes-cloud/kratosx/library/logger"
	"github.com/limes-cloud/kratosx/library/pprof"
	"github.com/limes-cloud/kratosx/library/registry"
	"github.com/limes-cloud/kratosx/library/web"
	"github.com/limes-cloud/kratosx/middleware"
	"github.com/limes-cloud/kratosx/server/cors"
	"github.com/limes-cloud/kratosx/server/response"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/protobuf/encoding/protojson"
)

type App struct {
	*options
}

func New(opts ...Option) *App {
	// 加载环境变量
	env.Load()

	// 默认配置项
	o := &options{
		config: config.New(
			kenv.NewSource("APP_"),
			file.NewSource("conf/"),
		),
	}

	for _, opt := range opts {
		opt(o)
	}

	// 加载配置
	if err := o.config.Load(); err != nil {
		panic(err)
	}

	// 配置监听
	if o.watch != nil {
		o.watch(o.config.ScanWatch)
	}

	library.Init(o.config, o.libOpts...)
	return &App{
		options: o,
	}
}

func (app *App) App() *kratos.App {
	if app.regSrvFn == nil {
		panic("must set registrarServerFn")
	}

	// 组装kratos默认参数
	defOpts := []kratos.Option{
		kratos.ID(app.config.App().ID),
		kratos.Name(app.config.App().Name),
		kratos.Version(app.config.App().Version),
		kratos.Metadata(map[string]string{}),
		kratos.BeforeStop(func(ctx context.Context) error {
			tasker.Instance().WaitBeforeStop()
			return nil
		}),
		kratos.AfterStop(func(ctx context.Context) error {
			tasker.Instance().WaitAfterStop()
			return nil
		}),
		kratos.BeforeStart(func(ctx context.Context) error {
			tasker.Instance().WaitBeforeStart()
			return nil
		}),
		kratos.AfterStart(func(ctx context.Context) error {
			tasker.Instance().WaitAfterStart()
			return nil
		}),
		kratos.Logger(logger.Instance()),
	}

	// 获取中间件
	gsOpts, hsOpts := serverOptions(app.config, app.midOpts, app.midHooks)

	// 获取grpc/http
	gsOpts = append(gsOpts, app.grpcSrvOptions...)
	hsOpts = append(hsOpts, app.httpSrvOptions...)

	// 获取http/grpc服务
	srv := app.config.App().Server
	gs := grpcServer(srv.Grpc, srv.Count, gsOpts)
	hs := httpServer(srv.Http, srv.Count, hsOpts)

	// 注册api服务
	app.options.regSrvFn(hs, gs)

	// 合并服务
	var srvList []transport.Server
	if srv.Http != nil {
		srvList = append(srvList, hs)
		// 监控
		if app.config.App().Metrics {
			hs.Handle("/metrics", promhttp.Handler())
		}
		// pprof
		if app.config.App().Server.Http.Pprof != nil {
			pprof.Server(app.config.App().Server.Http.Pprof, hs)
		}
		// webServer
		if app.config.App().Server.Http.WebServerDir != "" {
			web.Server(app.config.App().Server.Http.WebServerDir, hs)
		}
	}
	if srv.Grpc != nil {
		srvList = append(srvList, gs)
	}

	defOpts = append(defOpts, kratos.Server(srvList...))

	// 是否启用服务注册
	if srv.Registry != nil {
		reg, err := registry.Create(*srv.Registry)
		if err != nil {
			panic(err)
		}
		defOpts = append(defOpts, kratos.Registrar(reg))
	}

	// 注册关闭回调
	tasker.Instance().AfterStop("logger sync", func() {
		_ = logger.Instance().Sync()
	})

	return kratos.New(
		defOpts...,
	)
}

func serverOptions(conf config.Config, midOpts []kmid.Middleware, hook middleware.MidHook) ([]grpc.ServerOption, []http.ServerOption) {
	var gs []grpc.ServerOption
	var hs []http.ServerOption

	// 中间件
	mds := middleware.New(conf, hook)
	mds = append(mds, midOpts...)
	gs = append(gs, grpc.Middleware(mds...))
	hs = append(hs, http.Middleware(mds...))

	// tls
	if conf.App().Server.TLS != nil {
		cert, err := tls.X509KeyPair([]byte(conf.App().Server.TLS.Pem), []byte(conf.App().Server.TLS.Key))
		if err != nil {
			panic(err)
		}
		tlsConf := &tls.Config{Certificates: []tls.Certificate{cert}}
		gs = append(gs, grpc.TLSConfig(tlsConf))
		hs = append(hs, http.TLSConfig(tlsConf))
	}

	return gs, hs
}

func grpcServer(c *config.GrpcService, count int, so []grpc.ServerOption) *grpc.Server {
	if c == nil {
		return grpc.NewServer()
	}
	var opts []grpc.ServerOption
	opts = append(opts, so...)
	if c.Network != "" {
		opts = append(opts, grpc.Network(c.Network))
	}
	if c.Host != "" {
		if count > 1 {
			port, err := getUsablePort(c.Port, count)
			if err != nil {
				panic(err)
			}
			c.Port = port
		}
		addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
		opts = append(opts, grpc.Address(addr))
	}

	if c.Timeout != 0 {
		opts = append(opts, grpc.Timeout(-1))
	}
	return grpc.NewServer(opts...)
}

func httpServer(c *config.HttpService, count int, so []http.ServerOption) *http.Server {
	if c == nil {
		return http.NewServer()
	}
	var opts []http.ServerOption
	opts = append(opts, so...)
	if c.Network != "" {
		opts = append(opts, http.Network(c.Network))
	}
	if c.Host != "" {
		if count > 1 {
			port, err := getUsablePort(c.Port, count)
			if err != nil {
				panic(err)
			}
			c.Port = port
		}
		addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
		opts = append(opts, http.Address(addr))
	}

	if c.Timeout != 0 {
		opts = append(opts, http.Timeout(c.Timeout))
	}
	if c.FormatResponse {
		opts = append(opts, response.HTTPEncoder())
	}
	if c.Cors != nil {
		opts = append(opts, cors.Cors(c.Cors))
	}
	if c.Marshal != nil {
		json.MarshalOptions = protojson.MarshalOptions{
			EmitUnpopulated: c.Marshal.EmitUnpopulated, // 默认值不忽略
			UseProtoNames:   c.Marshal.UseProtoNames,   // 使用proto name返回http字段
		}
	}
	return http.NewServer(opts...)
}

// getUsablePort 查询可用的端口
func getUsablePort(base int, count int) (int, error) {
	// 获取的列表
	usablePort := func(port int) bool {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", port), timeout)
		if err == nil {
			// 能够连接说明端口确实在使用中
			conn.Close()
			return false
		}

		// 如果连接失败，尝试绑定监听
		addr := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			// 监听失败说明端口可能被占用
			return false
		}
		listener.Close()
		return true
	}

	for i := base; i < base+count; i++ {
		if usablePort(i) {
			return i, nil
		}
	}
	return 0, errors.New("not port usable")
}
