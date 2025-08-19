package library

import "github.com/limes-cloud/kratosx/library/logger"

type option struct {
	loggerOpts []logger.Option
}

type Option func(*option)

// WithLoggerOptions 添加日志配置选项
func WithLoggerOptions(opts ...logger.Option) Option {
	return func(o *option) {
		o.loggerOpts = opts
	}
}
