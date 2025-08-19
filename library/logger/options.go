package logger

type Option func(o *options)

type options struct {
	callerSkip int
	fields     map[string]any
}

// GetKV 获取kv对
func (o *options) GetKV() []any {
	// 转换kv
	var kvs []any
	if len(o.fields) != 0 {
		for k, v := range o.fields {
			kvs = append(kvs, k, v)
		}
	}
	return kvs
}

// WithCallerSkip 设置skip的数量
func WithCallerSkip(skip int) Option {
	return func(o *options) {
		o.callerSkip = skip
	}
}

// WithLogFields 设置日志字段
func WithLogFields(fs map[string]any) Option {
	return func(o *options) {
		o.fields = fs
	}
}
