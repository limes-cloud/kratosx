package logger

type Option func(o *options)

type options struct {
	callerSkip int
}

// AddCallerSkip 服务注册
func AddCallerSkip(skip int) Option {
	return func(o *options) {
		o.callerSkip = skip
	}
}
