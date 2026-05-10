package middleware

import (
	"context"
	"net"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc/peer"

	ip2 "github.com/limes-cloud/kratosx/library/ip"
)

func IP() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			ip := ""
			if p, is := peer.FromContext(ctx); is { // grpc
				if host, _, err := net.SplitHostPort(p.Addr.String()); err == nil {
					if host == "::1" {
						ip = "localhost"
					} else {
						ip = host
					}
				}
			}
			if h, is := http.RequestFromServerContext(ctx); is {
				if host, _, err := net.SplitHostPort(h.RemoteAddr); err == nil {
					if host == "::1" {
						ip = "localhost"
					} else {
						ip = host
					}
				}
				if h.Header.Get("x-real-ip") != "" {
					ip = h.Header.Get("x-real-ip")
				}
			}
			ctx = context.WithValue(ctx, ip2.Key{}, ip)
			return handler(ctx, req)
		}
	}
}
