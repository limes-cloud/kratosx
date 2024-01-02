package kratosx

import (
	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/server"
)

func grpcServer(c *config.GrpcService, md []middleware.Middleware) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(md...),
	}
	if c.Network != "" {
		opts = append(opts, grpc.Network(c.Network))
	}
	if c.Addr != "" {
		opts = append(opts, grpc.Address(c.Addr))
	}
	if c.Timeout != 0 {
		opts = append(opts, grpc.Timeout(c.Timeout))
	}
	return grpc.NewServer(opts...)
}

func httpServer(c *config.HttpService, md []middleware.Middleware) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(md...),
	}
	if c.Network != "" {
		opts = append(opts, http.Network(c.Network))
	}
	if c.Addr != "" {
		opts = append(opts, http.Address(c.Addr))
	}
	if c.Timeout != 0 {
		opts = append(opts, http.Timeout(c.Timeout))
	}
	if c.FormatResponse {
		opts = append(opts, server.HttpEncoder()...)
	}
	if c.Cors != nil {
		opts = append(opts, server.Cors(c.Cors))
	}
	if c.Marshal != nil {
		json.MarshalOptions = protojson.MarshalOptions{
			EmitUnpopulated: c.Marshal.EmitUnpopulated, // 默认值不忽略
			UseProtoNames:   c.Marshal.UseProtoNames,   // 使用proto name返回http字段
		}
	}
	return http.NewServer(opts...)
}
