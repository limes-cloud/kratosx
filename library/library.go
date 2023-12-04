package library

import (
	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/library/authentication"
	"github.com/limes-cloud/kratosx/library/captcha"
	"github.com/limes-cloud/kratosx/library/db"
	"github.com/limes-cloud/kratosx/library/email"
	"github.com/limes-cloud/kratosx/library/jwt"
	"github.com/limes-cloud/kratosx/library/loader"
	"github.com/limes-cloud/kratosx/library/logger"
	"github.com/limes-cloud/kratosx/library/logging"
	"github.com/limes-cloud/kratosx/library/pool"
	"github.com/limes-cloud/kratosx/library/redis"
)

func Init(conf config.Config, fs logger.LogField) {
	// 初始化全局日志
	logger.Init(conf.App().Log, conf.Watch, fs)

	// 初始化数据库
	db.Init(conf.App().Database, conf.Watch)

	// 初始化缓存
	redis.Init(conf.App().Redis, conf.Watch)

	// 初始化证书
	loader.Init(conf.App().Loader, conf.Watch)

	// 并发池初始化
	pool.Init(conf.App().Pool, conf.Watch)

	// 邮箱初始化
	email.Init(conf.App().Email, conf.Watch)

	// 验证码初始化
	captcha.Init(conf.App().Captcha, conf.Watch)

	// jwt初始化
	jwt.Init(conf.App().JWT, conf.Watch)

	// logging 初始化
	logging.Init(conf.App().Logging, conf.Watch)

	// authentication 鉴权器初始化
	authentication.Init(conf.App().Authentication, conf.Watch)
}
