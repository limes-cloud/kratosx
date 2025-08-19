package initializer

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	Init int    `json:"init"`
}

func New(db *gorm.DB, path string, force bool) Initializer {
	once.Do(func() {
		ins = &itr{
			path:  path,
			force: force,
			db:    db,
		}
		_ = db.AutoMigrate(GormInit{})
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

	// 查询创建状态
	init := GormInit{}
	if err := t.db.FirstOrCreate(&init, GormInit{Init: 0}).Error; err != nil {
		return err
	}

	// 非强制创建则返回
	if init.Init == 1 && !t.force {
		return nil
	}

	file, err := os.Open(t.path)
	if err != nil {
		return errors.New("open sql file error " + err.Error())
	}
	defer file.Close()

	var (
		sb      strings.Builder
		scanner = bufio.NewScanner(file)
		tx      = t.db.Begin()
	)
	skip := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// 忽略事物
		if line == "BEGIN;" || line == "COMMIT;" {
			continue
		}

		if strings.HasPrefix(strings.TrimSpace(line), "/*") {
			skip = true
			continue
		}

		if strings.HasSuffix(strings.TrimSpace(line), "*/") {
			skip = false
			continue
		}

		// 如果是注释或空行则跳过
		if strings.HasPrefix(line, "--") ||
			len(strings.TrimSpace(line)) == 0 ||
			skip {
			continue
		}

		// 替换默认字符集
		line = strings.ReplaceAll(line, "COLLATE=utf8mb4_0900_ai_ci", "")

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
	// 更新初始化标记
	if err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&GormInit{Id: 1, Init: 1}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
