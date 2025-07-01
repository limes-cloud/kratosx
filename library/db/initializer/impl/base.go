package impl

import (
	"errors"
	"github.com/limes-cloud/kratosx/library/db/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sync"
)

type BaseInitializer struct {
	mu    sync.Mutex // 保证线程安全
	DB    *gorm.DB
	Path  string
	Force bool
}

// SetPath 设置SQL文件路径（线程安全）
func (b *BaseInitializer) SetPath(path string) *BaseInitializer {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Path = path
	return b // 返回自身以支持链式调用
}

func (b *BaseInitializer) SetForce(force bool) *BaseInitializer {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Force = force
	return b // 返回自身以支持链式调用
}

// 检查初始化状态（使用GormInit）
func (b *BaseInitializer) CheckInitStatus() (bool, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	var initRecord model.GormInit
	err := b.DB.First(&initRecord, 1).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil // 未初始化
	}
	if err != nil {
		return false, err // 查询出错
	}
	return initRecord.Init, nil // 返回当前初始化状态
}

// 更新初始化状态
func (b *BaseInitializer) MarkInitialized(tx *gorm.DB) error {
	return tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&model.GormInit{Id: 1, Init: true}).Error
}

// GetPath 获取当前路径（内部使用）
func (b *BaseInitializer) getPath() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.Path
}

// IsForce 检查是否强制初始化（内部使用）
func (b *BaseInitializer) isForce() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.Force
}
