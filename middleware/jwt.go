package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	kratosJwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/jwt"
	"strings"
)

func Jwt(conf *config.JWT) middleware.Middleware {
	if conf == nil {
		return nil
	}

	keyFunc := func(token *jwtv4.Token) (any, error) {
		return []byte(conf.Secret), nil
	}

	whitelist := func(ctx context.Context) bool {
		operation, path := "", ""
		if tr, ok := transport.FromServerContext(ctx); ok {
			operation = tr.Operation()
		}
		if h, is := http.RequestFromServerContext(ctx); is {
			path = h.Method + ":" + h.URL.Path
		}

		jwtIns := jwt.Instance()
		return jwtIns.IsWhitelist(operation) || jwtIns.IsWhitelist(path)
	}

	return selector.Server(
		kratosJwt.Server(keyFunc),
	).Match(func(ctx context.Context, operation string) bool {
		return !whitelist(ctx)
	}).Build()
}

func JwtBlack(conf *config.JWT) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {

			header, ok := transport.FromServerContext(ctx)
			if !ok {
				return handler(ctx, req)
			}

			auths := strings.SplitN(header.RequestHeader().Get("Authorization"), " ", 2)
			if len(auths) != 2 || !strings.EqualFold(auths[0], "Bearer") {
				return handler(ctx, req)
			}

			token := auths[1]

			// 判断token是否在黑名单内
			jwtIns := jwt.Instance()
			if jwtIns.IsBlacklist(token) {
				return nil, errors.Unauthorized("UNAUTHORIZED", "JWT token is lose efficacy")
			}

			ctx = jwtIns.SetToken(ctx, token)
			return handler(ctx, req)
		}
	}

}
