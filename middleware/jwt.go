package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	kratosJwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv5 "github.com/golang-jwt/jwt/v5"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/jwt"
)

const (
	tokenKey = "x-md-global-token"
)

// Jwt jwt验证
func Jwt(conf *config.JWT) middleware.Middleware {
	if conf == nil {
		return nil
	}

	keyFunc := func(token *jwtv5.Token) (any, error) {
		return []byte(conf.Secret), nil
	}

	whitelist := func(ctx context.Context) bool {
		jwtIns := jwt.Instance()
		path, method := "", ""

		if tr, ok := transport.FromServerContext(ctx); ok {
			path = tr.Operation()
			method = "GRPC"
		}
		if h, is := http.RequestFromServerContext(ctx); is {
			path = h.URL.Path
			method = h.Method
		}

		return jwtIns.IsWhitelist(path, method)
	}

	return selector.Server(
		kratosJwt.Server(keyFunc),
	).Match(func(ctx context.Context, operation string) bool {
		return !whitelist(ctx)
	}).Build()
}

func JwtToken() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			header, ok := transport.FromServerContext(ctx)
			if !ok {
				return handler(ctx, req)
			}

			auths := strings.SplitN(header.RequestHeader().Get("Authorization"), " ", 2)
			if len(auths) != 2 || !strings.EqualFold(auths[0], "Bearer") {
				return handler(ctx, req)
			}

			token := auths[1]
			if md, ok := metadata.FromServerContext(ctx); ok {
				md.Set(tokenKey, token)
			}
			return handler(ctx, req)
		}
	}
}

// JwtBlack jwt黑名单
func JwtBlack(conf *config.JWT) middleware.Middleware {
	if conf == nil {
		return nil
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			md, ok := metadata.FromServerContext(ctx)
			if !ok {
				return handler(ctx, req)
			}

			token := md.Get(tokenKey)
			// 判断token是否在黑名单内
			jwtIns := jwt.Instance()
			if jwtIns.IsBlacklist(token) {
				return nil, kratosJwt.ErrTokenInvalid
			}

			ctx = jwtIns.SetToken(ctx, token)
			return handler(ctx, req)
		}
	}
}

// JwtUnique jwt唯一校验
func JwtUnique(conf *config.JWT) middleware.Middleware {
	if conf == nil || !conf.Unique {
		return nil
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			claims, ok := kratosJwt.FromContext(ctx)
			if !ok {
				return handler(ctx, req)
			}

			mapClaims, ok := claims.(jwtv5.MapClaims)
			if !ok {
				return handler(ctx, req)
			}

			// 获取uniqueKey
			jwtIns := jwt.Instance()
			uk := fmt.Sprint(mapClaims[conf.UniqueKey])

			// 对比token
			if !jwtIns.CompareUniqueToken(uk, jwtIns.GetToken(ctx)) {
				return nil, kratosJwt.ErrTokenInvalid
			}
			return handler(ctx, req)
		}
	}
}
