package db

import "github.com/limes-cloud/kratosx/model/hook"

type Option func(o *options)

type options struct {
	hook hook.ScopeRequestFunc
}

// WithHookScope 设置hook 权限
func WithHookScope(scope hook.ScopeRequestFunc) Option {
	return func(o *options) {
		o.hook = scope
	}
}
