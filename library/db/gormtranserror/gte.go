package gormtranserror

import "sync"

var (
	ins  GormErrorPlugin
	once sync.Once
)

func NewGlobalGormErrorPlugin(opts ...Option) GormErrorPlugin {
	once.Do(func() {
		ins = NewGormErrorPlugin(opts...)
		if ins.options().db != nil {
			_ = ins.options().db.Use(ins)
		}
	})
	return ins
}
