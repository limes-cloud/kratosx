package kratosx

import (
	"crypto/tls"

	kmid "github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/middleware"
)

func serverOptions(conf config.Config, midOpts []kmid.Middleware) ([]grpc.ServerOption, []http.ServerOption) {
	var gs []grpc.ServerOption
	var hs []http.ServerOption

	// 中间件
	mds := middleware.New(conf)
	mds = append(mds, midOpts...)
	gs = append(gs, grpc.Middleware(mds...))
	hs = append(hs, http.Middleware(mds...))

	// tls
	if conf.App().Server.Tls != nil {
		cert, err := tls.X509KeyPair([]byte(conf.App().Server.Tls.Pem), []byte(conf.App().Server.Tls.Key))
		if err != nil {
			panic(err)
		}
		tlsConf := &tls.Config{Certificates: []tls.Certificate{cert}}

		gs = append(gs, grpc.TLSConfig(tlsConf))
		hs = append(hs, http.TLSConfig(tlsConf))
	}

	return gs, hs
}
