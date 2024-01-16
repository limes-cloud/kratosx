package middleware

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"

	"github.com/limes-cloud/kratosx/config"
)

func New(conf config.Config) []middleware.Middleware {
	app := conf.App()

	mds := []middleware.Middleware{
		recovery.Recovery(),
		RateLimit(app.RateLimit),

		Tracer(app.Name, app.Tracing),
		Logging(app.Logging),
		validate.Validator(),
		Signature(app.Signature),

		IP(),
		Jwt(app.JWT),
		JwtBlack(app.JWT),
		JwtUnique(app.JWT),
		Authentication(app.Authentication),

		metadata.Server(),
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
