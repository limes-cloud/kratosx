package db

import (
	"fmt"
	"github.com/limes-cloud/kratosx/library/db/model"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/limes-cloud/kratosx/config"
	gte "github.com/limes-cloud/kratosx/library/db/gormtranserror"
	"github.com/limes-cloud/kratosx/library/db/initializer"
)

type DB interface {
	// Get 获取指定名称的db实例，不指定名称则返回第一个如果实例不存在则返回nil
	Get(name ...string) *gorm.DB

	TxKey(name ...string) string
}

type db struct {
	mu  sync.RWMutex
	set map[string]*gorm.DB
	key string
}

const (
	_mysql      = "mysql"
	_postgresql = "postgresql"
	_sqlServer  = "sqlServer"
	_tidb       = "tidb"
	_clickhouse = "clickhouse"
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

		watcher("database."+key, func(value config.Value) {
			if err := value.Scan(conf); err != nil {
				log.Errorf("Database配置变更失败：%s", err.Error())
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

	if conf.AutoCreate {
		if err := d.create(conf); err != nil {
			panic("auto create database error:" + err.Error())
		}
	}

	// 连接主数据库
	client, err := gorm.Open(d.open(conf), &gorm.Config{
		Logger: newLog(conf.Config.LogLevel, conf.Config.SlowThreshold),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.Config.TablePrefix,
			SingularTable: true,
		},
		PrepareStmt: conf.Config.PrepareStmt,
		DryRun:      conf.Config.DryRun,
	})
	if err != nil {
		return err
	}

	if conf.Config.TransformError != nil && conf.Config.TransformError.Enable {
		opts := []gte.Option{
			gte.WithEnableLoad(),
		}

		format := conf.Config.TransformError.Format
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
		if err := gte.NewGlobalGormErrorPlugin(opts...).Initialize(client); err != nil {
			panic("gorm transform error:" + err.Error())
		}
	}

	if conf.Config.Initializer != nil && conf.Config.Initializer.Enable {
		if err := initializer.New(model.DatabaseType(conf.Drive), client, conf.Config.Initializer.Path, conf.Config.Initializer.Force).Exec(); err != nil {
			panic("db init error:" + err.Error())
		}
	}

	sdb, _ := client.DB()
	sdb.SetConnMaxLifetime(conf.Config.MaxLifetime)
	sdb.SetMaxOpenConns(conf.Config.MaxOpenConn)
	sdb.SetMaxIdleConns(conf.Config.MaxIdleConn)

	d.mu.Lock()
	d.set[name] = client
	d.key = name
	d.mu.Unlock()
	return nil
}

func (d *db) TxKey(name ...string) string {
	if d.key == "" && len(name) == 0 {
		return ""
	}
	key := d.key
	if len(name) != 0 {
		key = name[0]
	}
	return "db_tx_" + key
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

func (d *db) open(conf *config.Database) gorm.Dialector {
	switch conf.Drive {
	case _mysql, _tidb:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
			conf.Connect.Username,
			conf.Connect.Password,
			conf.Connect.Host,
			conf.Connect.Port,
			conf.Connect.DBName,
			conf.Connect.Option,
		)
		return mysql.Open(dsn)
	case _postgresql:
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d %s",
			conf.Connect.Host,
			conf.Connect.Username,
			conf.Connect.Password,
			conf.Connect.DBName,
			conf.Connect.Port,
			conf.Connect.Option,
		)
		return postgres.Open(dsn)
	case _sqlServer:
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			conf.Connect.Username,
			conf.Connect.Password,
			conf.Connect.Host,
			conf.Connect.Port,
			conf.Connect.DBName,
		)
		return sqlserver.Open(dsn)
	case _clickhouse:
		dsn := fmt.Sprintf("tcp://%s:%d?database=%s&username=%s&password=%s%s",
			conf.Connect.Host,
			conf.Connect.Port,
			conf.Connect.DBName,
			conf.Connect.Username,
			conf.Connect.Password,
			conf.Connect.Option,
		)
		return clickhouse.Open(dsn)
	default:
		return nil
	}
}

func (d *db) create(conf *config.Database) error {
	copyConf := *conf
	//copyConf.Connect.DBName = ""
	//copyConf.Connect.Option = ""

	connect, err := gorm.Open(d.open(&copyConf))
	if err != nil {
		return err
	}
	_ = connect.Session(&gorm.Session{
		Logger: logger.Default.LogMode(logger.Silent),
	}).Exec(fmt.Sprintf("CREATE DATABASE %s", conf.Connect.DBName))
	return nil
}
