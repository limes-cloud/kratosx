package initializer

import (
	"errors"
	"io"
	"os"
	"strings"
	"sync"

	"gorm.io/gorm"
)

type Initializer interface {
	Exec() error
	SetForce(force bool) Initializer
	SetPath(path string) Initializer
}

var (
	ins  *itr
	once sync.Once
)

type itr struct {
	path  string
	force bool
	db    *gorm.DB
}

type GormInit struct {
	Id   uint32 `json:"id"`
	Init bool   `json:"init"`
}

func Instance() Initializer {
	return ins
}

func New(db *gorm.DB, path string, force bool) Initializer {
	once.Do(func() {
		ins = &itr{
			path:  path,
			force: force,
			db:    db,
		}
		db.AutoMigrate(GormInit{})
	})
	return ins
}

func (t *itr) SetForce(force bool) Initializer {
	t.force = force
	return t
}

func (t *itr) SetPath(path string) Initializer {
	t.path = path
	return t
}

func (t *itr) Exec() error {
	if t == nil {
		return errors.New("not enable initializer")
	}

	init := GormInit{}
	if err := t.db.FirstOrCreate(&init, GormInit{Init: false}).Error; err != nil {
		return err
	}

	if init.Init && !t.force {
		return nil
	}

	file, err := os.Open(t.path)
	if err != nil {
		return errors.New("open sql file error " + err.Error())
	}

	byteData, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	str := string(byteData)
	str = strings.ReplaceAll(str, "BEGIN;", "")
	str = strings.ReplaceAll(str, "COMMIT;", "")
	str = strings.ReplaceAll(str, "COLLATE=utf8mb4_0900_ai_ci", "")

	tx := t.db.Begin()
	sqlList := strings.Split(str, ";\n")
	for _, item := range sqlList {
		sql := strings.TrimSpace(item)
		if sql == "" {
			continue
		}
		if err := tx.Exec(sql + ";").Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Model(GormInit{}).Where("id=1").UpdateColumn("init", 1)
	tx.Commit()
	return nil
}
