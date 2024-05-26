package kratosx

import (
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/cors"
	"github.com/limes-cloud/kratosx/library/httpencoder"
)

func grpcServer(c *config.GrpcService, count int, so []grpc.ServerOption) *grpc.Server {
	if c == nil {
		return grpc.NewServer()
	}
	var opts []grpc.ServerOption
	for _, o := range so {
		opts = append(opts, o)
	}
	if c.Network != "" {
		opts = append(opts, grpc.Network(c.Network))
	}
	if c.Addr != "" {
		opts = append(opts, grpc.Address(c.Addr))
	} else {
		c.Addr = fmt.Sprintf("%s:%d", c.Host, c.Port)
		if count > 1 {
			port, err := getPort(c.Port, count)
			if err != nil {
				panic(err)
			}
			c.Addr = fmt.Sprintf("%s:%d", c.Host, port)
		}
		opts = append(opts, grpc.Address(c.Addr))
	}

	if c.Timeout != 0 {
		opts = append(opts, grpc.Timeout(c.Timeout))
	}
	return grpc.NewServer(opts...)
}

func httpServer(c *config.HttpService, count int, so []http.ServerOption) *http.Server {
	if c == nil {
		return http.NewServer()
	}
	var opts []http.ServerOption
	for _, o := range so {
		opts = append(opts, o)
	}
	if c.Network != "" {
		opts = append(opts, http.Network(c.Network))
	}
	if c.Addr != "" {
		opts = append(opts, http.Address(c.Addr))
	} else {
		c.Addr = fmt.Sprintf("%s:%d", c.Host, c.Port)
		if count > 1 {
			port, err := getPort(c.Port, count)
			if err != nil {
				panic(err)
			}
			c.Addr = fmt.Sprintf("%s:%d", c.Host, port)
		}
		opts = append(opts, http.Address(c.Addr))
	}
	if c.Timeout != 0 {
		opts = append(opts, http.Timeout(c.Timeout))
	}
	if c.FormatResponse {
		opts = append(opts, httpencoder.HttpEncoder()...)
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

func usablePort(port int) bool {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", strconv.Itoa(port)))
	if err != nil {
		return false
	}
	defer l.Close()
	return true
}

func getPort(base int, count int) (int, error) {
	for i := base; i < base+count; i++ {
		if usablePort(i) {
			return i, nil
		}
	}
	return 0, errors.New("not usable port")
}
