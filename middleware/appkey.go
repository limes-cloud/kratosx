package middleware

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/limes-cloud/kratosx/pkg/appkey"
)

// AppKey extracts the current app key from the request header and stores it in metadata.
func AppKey() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			appKey := ""
			if tr, ok := transport.FromServerContext(ctx); ok {
				appKey = strings.TrimSpace(tr.RequestHeader().Get(appkey.HeaderKey))
			}
			if appKey == "" {
				if md, ok := metadata.FromServerContext(ctx); ok {
					appKey = strings.TrimSpace(md.Get(appkey.MetadataKey))
				}
			}
			if appKey != "" {
				if md, ok := metadata.FromServerContext(ctx); ok {
					md.Set(appkey.MetadataKey, appKey)
				}
				ctx = metadata.AppendToClientContext(ctx, appkey.MetadataKey, appKey)
			}
			return handler(ctx, req)
		}
	}
}
