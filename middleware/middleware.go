package middleware

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/validate"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/signature"
)

func New(conf config.Config) []middleware.Middleware {
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
		// ip
		IP(),
		// jwt
		Jwt(app.JWT),
		// jwt黑名单
		JwtBlack(app.JWT),
		// jwt 唯一设备
		JwtUnique(app.JWT),
		// 参数校验
		//nolint
		validate.Validator(),
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
