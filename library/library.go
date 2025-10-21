package library

import (
	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/captcha"
	"github.com/limes-cloud/kratosx/library/client"
	"github.com/limes-cloud/kratosx/library/db"
	"github.com/limes-cloud/kratosx/library/email"
	"github.com/limes-cloud/kratosx/library/jwt"
	"github.com/limes-cloud/kratosx/library/loader"
	"github.com/limes-cloud/kratosx/library/logger"
	"github.com/limes-cloud/kratosx/library/logging"
	"github.com/limes-cloud/kratosx/library/pool"
	"github.com/limes-cloud/kratosx/library/prometheus"
	"github.com/limes-cloud/kratosx/library/redis"
	"github.com/limes-cloud/kratosx/library/request"
	"github.com/limes-cloud/kratosx/library/signature"
)

func Init(conf config.Config, opts ...Option) {
	o := option{}
	for _, opt := range opts {
		opt(&o)
	}

	// 初始化全局日志
	logger.Init(conf.App().Logger, o.loggerOpts...)

	// 初始化数据库
	db.Init(conf.App().Database, o.dbOpts...)

	// 初始化缓存
	redis.Init(conf.App().Redis)

	// 初始化证书
	loader.Init(conf.App().Loader, conf.Watch)

	// 并发池初始化
	pool.Init(conf.App().Pool)

	// 邮箱初始化
	email.Init(conf.App().Email, conf.Watch)

	// 验证码初始化
	captcha.Init(conf.App().Captcha, conf.Watch)

	// jwt初始化
	jwt.Init(conf.App().JWT, conf.Watch)

	// logging 初始化
	logging.Init(conf.App().Logging, conf.Watch)

	// grpc 客户端初始化
	client.Init(conf.App().Server.Registry, conf.App().Client, conf.Watch)

	// 签名验证器初始化
	signature.Init(conf.App().Signature, conf.Watch)

	// 初始化监控
	prometheus.Init(conf.App().Prometheus, conf.Watch)

	// request工具初始化
	request.Init(conf.App().Request, conf.Watch)
}
