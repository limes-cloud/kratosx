package library

import (
	"github.com/limes-cloud/kratosx/library/db"
	"github.com/limes-cloud/kratosx/library/logger"
)

type option struct {
	loggerOpts []logger.Option
	dbOpts     []db.Option
}

type Option func(*option)

// WithLoggerOptions 添加日志配置选项
func WithLoggerOptions(opts ...logger.Option) Option {
	return func(o *option) {
		o.loggerOpts = opts
	}
}

// WithDBOptions 添加DB配置选项
func WithDBOptions(opts ...db.Option) Option {
	return func(o *option) {
		o.dbOpts = opts
	}
}
