package db

import (
	"context"
	"testing"

	"github.com/limes-cloud/kratosx/model/hook"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type hookGateModel struct {
	ID uint32 `gorm:"column:id;primaryKey"`
}

func (hookGateModel) TableName() string {
	return "hook_gate_models"
}

func newHookGateDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@tcp(127.0.0.1:3306)/kratosx?charset=utf8mb4&parseTime=True&loc=Local",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{DryRun: true, DisableAutomaticPing: true, SkipDefaultTransaction: true})
	if err != nil {
		t.Fatalf("open dry-run db: %v", err)
	}
	return db
}

func TestApplyRequiresWithHookContext(t *testing.T) {
	var calls int
	scope := func(context.Context, string, string, string) (hook.ScopeResponse, error) {
		calls++
		return nil, nil
	}
	db := newHookGateDB(t)
	registerHook("test:hook:gate", "db", db, scope)

	var rows []hookGateModel
	if err := db.Model(&hookGateModel{}).Find(&rows).Error; err != nil {
		t.Fatalf("query without hook: %v", err)
	}
	if calls != 0 {
		t.Fatalf("expected hook to be disabled by default, got %d calls", calls)
	}

	if err := db.WithContext(WithHook(context.Background())).Model(&hookGateModel{}).Find(&rows).Error; err != nil {
		t.Fatalf("query with hook: %v", err)
	}
	if calls != 1 {
		t.Fatalf("expected hook to run once, got %d calls", calls)
	}
}
