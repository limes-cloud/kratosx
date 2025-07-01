package impl

import (
	"bufio"
	"errors"
	"github.com/limes-cloud/kratosx/library/db/model"
	"gorm.io/gorm"
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
		line := scanner.Text()

		// 忽略事务
		if line == "BEGIN;" || line == "COMMIT;" {
			continue
		}
		// 如果是注释或空行则跳过
		if strings.HasPrefix(line, "--") || strings.HasPrefix(line, "/*") || len(strings.TrimSpace(line)) == 0 {
			continue
		}

		sb.WriteString(line + "\n")

		// 如果SQL语句以分号结尾，则执行
		if strings.HasSuffix(line, ";") {
			sql := sb.String()
			sb.Reset()

			if strings.Contains(sql, "gorm_init") {
				continue
			}

			if err = tx.Exec(sql).Error; err != nil {
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
