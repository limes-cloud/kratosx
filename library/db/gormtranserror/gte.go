package gormtranserror

import "sync"

var (
	_ins GormErrorPlugin
	once sync.Once
)

func NewGlobalGormErrorPlugin(opts ...Option) GormErrorPlugin {
	once.Do(func() {
		_ins = NewGormErrorPlugin(opts...)
		if _ins.options().db != nil {
			_ins.options().db.Use(_ins)
		}
	})
	return _ins
}
