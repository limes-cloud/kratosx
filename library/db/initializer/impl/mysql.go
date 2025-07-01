package impl

import (
	"bufio"
	"errors"
	"github.com/limes-cloud/kratosx/library/db/model"
	"gorm.io/gorm"
	"os"
	"strings"
)

type mysqlInitializer struct {
	BaseInitializer
}

func NewMySQLInitializer(db *gorm.DB, path string, force bool) model.Initializer {
	return &mysqlInitializer{
		BaseInitializer: BaseInitializer{
			DB:    db,
			Path:  path,
			Force: force,
		},
	}
}

func (m *mysqlInitializer) SetPath(path string) model.Initializer {
	m.BaseInitializer.SetPath(path)
	return m // 返回具体实现类型以保持链式调用
}

func (m *mysqlInitializer) SetForce(force bool) model.Initializer {
	m.BaseInitializer.SetForce(force)
	return m
}

func (m *mysqlInitializer) Exec() error {
	// 1. 检查初始化状态（不中断流程）
	alreadyInit, err := m.CheckInitStatus()
	if err != nil {
		return err
	}
	if alreadyInit && !m.Force {
		return nil // 符合条件时提前返回
	}
	// 2. 打开SQL文件
	file, err := os.Open(m.Path)
	if err != nil {
		return errors.New("open sql file error " + err.Error())
	}
	defer file.Close()

	// 3. 开始事务
	tx := m.DB.Begin()
	scanner := bufio.NewScanner(file)
	var sb strings.Builder

	// 4. 逐行处理SQL
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "--") || line == "" {
			continue
		}

		// MySQL特定处理
		line = strings.ReplaceAll(line, "COLLATE=utf8mb4_0900_ai_ci", "")
		sb.WriteString(line)

		if strings.HasSuffix(line, ";") {
			sql := sb.String()
			sb.Reset()

			if err := tx.Exec(sql).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// 5. 更新初始化状态
	if err := m.MarkInitialized(tx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
