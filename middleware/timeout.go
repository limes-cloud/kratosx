package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"

	ec "github.com/limes-cloud/kratosx/config"
)

func Timeout(gs *ec.GrpcService, hs *ec.HttpService) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			var timeout time.Duration
			if tr, ok := transport.FromServerContext(ctx); ok {
				key := fmt.Sprintf("%s:%s", "GRPC", tr.Operation())
				t, ok := gs.TimeoutSpecial[key]
				if ok {
					timeout = t
				} else {
					timeout = gs.Timeout
				}
			}
			if h, is := http.RequestFromServerContext(ctx); is {
				key := fmt.Sprintf("%s:%s", h.Method, h.URL.Path)
				t, ok := hs.TimeoutSpecial[key]
				if ok {
					timeout = t
				} else {
					timeout = hs.Timeout
				}
			}

			var (
				cctx   context.Context
				cancel context.CancelFunc
			)
			if timeout > 0 {
				cctx, cancel = context.WithTimeout(ctx, timeout)
			} else {
				cctx, cancel = context.WithCancel(ctx)
			}
			defer func() {
				cancel()
			}()

			return handler(cctx, req)
		}
	}
}
