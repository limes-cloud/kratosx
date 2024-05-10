package client

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	midmetadata "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/logger"
	"github.com/limes-cloud/kratosx/library/signature"
)

func Middlewares(conf *config.Client) []middleware.Middleware {
	mds := []middleware.Middleware{
		midmetadata.Client(),
		logging.Client(logger.Instance(logger.AddCallerSkip(-1))),
		circuitbreaker.Client(),
		tracing.Client(),
		signature.Instance().Client(conf.Signature),
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
