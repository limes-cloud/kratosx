package db

import (
	kratosConfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/limes-cloud/kratosx/config"
	"sync"

	"github.com/glebarez/sqlite"
	"github.com/limes-cloud/library/gte"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DB interface {
	// Get 获取指定名称的db实例，不指定名称则返回第一个如果实例不存在则返回nil
	Get(name ...string) *gorm.DB
}

type db struct {
	mu  sync.RWMutex
	set map[string]*gorm.DB
	key string
}

const (
	_mysql      = "mysql"
	_postgresql = "postgresql"
	_sqlite     = "sqlite"
	_sqlServer  = "sqlServer"
	_tidb       = "tidb"
)

var instance *db

func Instance() DB {
	return instance
}

// Init 初始化全局db
func Init(cfs map[string]*config.Database, watcher config.Watcher) {
	if len(cfs) == 0 {
		return
	}

	instance = &db{
		mu:  sync.RWMutex{},
		set: make(map[string]*gorm.DB),
	}

	// 遍历配置连接数据库
	for key, conf := range cfs {
		if conf == nil {
			continue
		}

		if err := instance.initFactory(key, conf); err != nil {
			panic("database init error :" + err.Error())
		}

		watcher("database."+key, func(value kratosConfig.Value) {
			if err := value.Scan(conf); err != nil {
				log.Error("Database配置变更失败：%s", err.Error())
				return
			}
			if err := instance.initFactory(key, conf); err != nil {
				log.Errorf("Database变更重载失败：%s", err.Error())
			}
		})

	}

	// 如果配置了多个库，则不能启用快速获取
	if len(instance.set) != 1 {
		instance.key = ""
	}
}

func (d *db) initFactory(name string, conf *config.Database) error {
	if !conf.Enable {
		d.delete(name)
		return nil
	}

	// 连接主数据库
	client, err := gorm.Open(d.open(conf.Drive, conf.Dsn), &gorm.Config{
		Logger: newLog(conf.LogLevel, conf.SlowThreshold),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.TablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}

	if conf.TransformError != nil {
		opts := []gte.Option{
			gte.WithEnableLoad(),
		}
		format := conf.TransformError.Format
		if format != nil {
			if format.AddForeign != nil {
				opts = append(opts, gte.WithAddForeignKeyFormat(*format.AddForeign))
			}
			if format.DelForeign != nil {
				opts = append(opts, gte.WithDelForeignKeyFormat(*format.DelForeign))
			}
			if format.Duplicated != nil {
				opts = append(opts, gte.WithDuplicatedKeyFormat(*format.Duplicated))
			}
		}
		gte.NewGormErrorPlugin(opts...)
	}

	sdb, _ := client.DB()
	sdb.SetConnMaxLifetime(conf.MaxLifetime)
	sdb.SetMaxOpenConns(conf.MaxOpenConn)
	sdb.SetMaxIdleConns(conf.MaxIdleConn)

	d.mu.Lock()
	d.set[name] = client
	d.key = name
	d.mu.Unlock()
	return nil
}

func (d *db) delete(name string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.set, name)
}

func (d *db) Get(name ...string) *gorm.DB {
	if d.key == "" && len(name) == 0 {
		return nil
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	key := d.key
	if len(name) != 0 {
		key = name[0]
	}
	return d.set[key]
}

func (d *db) open(drive, dsn string) gorm.Dialector {
	switch drive {
	case _mysql, _tidb:
		return mysql.Open(dsn)
	case _postgresql:
		return postgres.Open(dsn)
	case _sqlite:
		return sqlite.Open(dsn)
	case _sqlServer:
		return sqlserver.Open(dsn)
	default:
		return nil
	}
}
