package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc/peer"
	"strings"
)

func IP() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			ip := ""
			if p, is := peer.FromContext(ctx); is { //grpc
				if strings.Contains(p.Addr.String(), "::1") {
					ip = "localhost"
				} else {
					ip = strings.Split(p.Addr.String(), ":")[0]
				}
			}
			if h, is := http.RequestFromServerContext(ctx); is {
				if strings.Contains(h.RemoteAddr, "::1") {
					ip = "localhost"
				} else {
					ip = strings.Split(h.RemoteAddr, ":")[0]
				}
				if h.Header.Get("x-real-ip") != "" {
					ip = h.Header.Get("x-real-ip")
				}
			}
			ctx = context.WithValue(ctx, "ClientIP", ip)
			return handler(ctx, req)
		}
	}
}
