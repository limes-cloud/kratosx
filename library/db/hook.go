package db

import (
	"context"

	"github.com/limes-cloud/kratosx/model/hook"
	"gorm.io/gorm"
)

type hookKey struct{}

var (
	HookKey = hookKey{}
)

// WithHook 启用当前上下文中的 DB hook。
func WithHook(ctx context.Context) context.Context {
	return context.WithValue(ctx, HookKey, true)
}

func registerHook(name string, dbName string, dbClient *gorm.DB, scope hook.ScopeRequestFunc) {
	_ = dbClient.Callback().Query().Before("gorm:query").Register(name+":gorm:query", func(idb *gorm.DB) {
		apply(idb, dbName, hook.Read, scope)
	})

	_ = dbClient.Callback().Create().Before("gorm:create").Register(name+":gorm:create", func(idb *gorm.DB) {
		apply(idb, dbName, hook.Create, scope)
	})

	_ = dbClient.Callback().Update().Before("gorm:update").Register(name+":gorm:update", func(idb *gorm.DB) {
		apply(idb, dbName, hook.Update, scope)
	})

	_ = dbClient.Callback().Delete().Before("gorm:delete").Register(name+":gorm:delete", func(idb *gorm.DB) {
		apply(idb, dbName, hook.Delete, scope)
	})
}

func apply(idb *gorm.DB, dbName, method string, req hook.ScopeRequestFunc) {
	enable, _ := idb.Statement.Context.Value(HookKey).(bool)
	if !enable {
		return
	}
	hook.Apply(idb, dbName, method, req)
}
