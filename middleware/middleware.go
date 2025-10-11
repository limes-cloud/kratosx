package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/signature"
)

type MidHook struct {
	ValidateErrHook func(ctx context.Context, err error) error
}

func New(conf config.Config, hook MidHook) []middleware.Middleware {
	app := conf.App()

	mds := []middleware.Middleware{
		Recovery(),
		// 超时
		Timeout(app.Server.Grpc, app.Server.Http),
		// 限流
		RateLimit(app.RateLimit),
		// 监控
		Metrics(app.Metrics),
		// 元数据
		metadata.Server(),
		// 签名
		signature.Instance().Server(),
		// 链路
		Tracer(app.Name, app.Tracing),
		// 请求日志
		Logging(app.Logging),
		// 参数校验
		//nolint
		Validator(hook.ValidateErrHook),
		// ip
		IP(),
		// jwt
		Jwt(app.JWT),
		// jwt token
		JwtToken(),
		// jwt黑名单
		JwtBlack(app.JWT),
		// jwt 唯一设备
		JwtUnique(app.JWT),
	}
	// 原地删除不启用的中间件
	return removeDisableMiddleware(mds)
}

func removeDisableMiddleware(slice []middleware.Middleware) []middleware.Middleware {
	fast, slow := 0, 0
	for fast < len(slice) {
		if slice[fast] != nil {
			slice[slow] = slice[fast]
			slow++
		}
		fast++
	}
	return slice[:slow]
}
