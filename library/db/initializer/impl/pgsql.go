package impl

import (
	"bufio"
	"errors"
	"github.com/limes-cloud/kratosx/library/db/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
	"strings"
)

type pgsqlInitializer struct {
	BaseInitializer
}

func NewPgSQLInitializer(db *gorm.DB, path string, force bool) model.Initializer {
	return &pgsqlInitializer{
		BaseInitializer: BaseInitializer{
			DB:    db,
			Path:  path,
			Force: force,
		},
	}
}

func (p *pgsqlInitializer) SetPath(path string) model.Initializer {
	p.BaseInitializer.SetPath(path)
	return p // 返回具体实现类型以保持链式调用
}

func (p *pgsqlInitializer) SetForce(force bool) model.Initializer {
	p.BaseInitializer.SetForce(force)
	return p
}

func (p *pgsqlInitializer) Exec() error {
	// 1. 检查初始化状态（不中断流程）
	alreadyInit, err := p.CheckInitStatus()
	if err != nil {
		return err
	}
	if alreadyInit && !p.Force {
		return nil // 符合条件时提前返回
	}

	file, err := os.Open(p.Path)
	if err != nil {
		return errors.New("open sql file error " + err.Error())
	}
	defer file.Close()

	tx := p.DB.Begin()
	scanner := bufio.NewScanner(file)
	var sb strings.Builder

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "--") || line == "" {
			continue
		}

		// PostgreSQL-specific processing
		// Remove any MySQL-specific syntax if present in the SQL file
		line = strings.ReplaceAll(line, "`", "\"") // Replace backticks with double quotes
		sb.WriteString(line)
		sb.WriteString(" ") // Add space between lines

		// Check for statement termination (PostgreSQL uses semicolons too)
		if strings.HasSuffix(strings.TrimRight(line, " "), ";") {
			sql := sb.String()
			sb.Reset()

			if err := tx.Exec(sql).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if err := p.MarkInitialized(tx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (p *pgsqlInitializer) MarkInitialized(tx *gorm.DB) error {
	return tx.Model(&model.GormInit{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(map[string]interface{}{
		"id":   1,
		"init": 1, // 直接使用整数 1 代替 true
	}).Error
}
