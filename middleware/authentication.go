package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"

	ec "github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/authentication"
)

const (
	reason    string = "FORBIDDEN"
	noSupport string = "ONLY_SUPPORT_HTTP"
)

func Authentication(conf *ec.Authentication) middleware.Middleware {
	if conf == nil {
		return nil
	}

	author := authentication.Instance()
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			path, method := "", ""

			if tr, ok := transport.FromServerContext(ctx); ok {
				path = tr.Operation()
			}

			h, is := http.RequestFromServerContext(ctx)
			if is {
				path = h.URL.Path
				method = h.Method
			} else {
				method = "GRPC"
			}

			if author.IsWhitelist(path, method) {
				return handler(ctx, req)
			}

			role, er := author.GetRole(ctx)
			if er != nil {
				return nil, errors.Forbidden(reason, er.Error())
			}

			// 不存在token也跳过鉴权
			if role == "" && er == nil {
				return handler(ctx, req)
			}

			// 如果是跳过白名单内的，也不检验
			if author.IsSkipRole(role) {
				return handler(ctx, req)
			}

			if !author.Auth(role, path, method) {
				return nil, errors.Forbidden(reason, "无接口权限")
			}
			return handler(ctx, req)
		}
	}
}
