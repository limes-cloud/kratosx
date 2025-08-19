package db

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
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
	set map[string]*gorm.DB
	key string
}

const (
	_mysql      = "mysql"
	_postgresql = "postgres"
	_sqlServer  = "sqlserver"
	_tidb       = "tidb"
	_clickhouse = "clickhouse"
)

var (
	ins *db

	once sync.Once
)

func Instance() DB {
	return ins
}

// Init 初始化全局db
func Init(cfs []*config.Database) {
	if len(cfs) == 0 {
		return
	}

	once.Do(func() {
		ins = &db{
			set: make(map[string]*gorm.DB),
		}

		// 遍历配置连接数据库
		for ind, conf := range cfs {
			if err := ins.initFactory(conf); err != nil {
				panic("database init error :" + err.Error())
			}
			if ind == 0 {
				ins.key = conf.Name
			}
		}
	})
}

func (d *db) initFactory(conf *config.Database) error {
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
		if err := initializer.New(
			client,
			conf.Config.Initializer.Path,
			conf.Config.Initializer.Force,
		).Exec(); err != nil {
			panic("db init error:" + err.Error())
		}
	}

	sdb, _ := client.DB()
	sdb.SetConnMaxLifetime(conf.Config.MaxLifetime)
	sdb.SetMaxOpenConns(conf.Config.MaxOpenConn)
	sdb.SetMaxIdleConns(conf.Config.MaxIdleConn)

	d.set[conf.Name] = client
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

func (d *db) Get(name ...string) *gorm.DB {
	if d.key == "" && len(name) == 0 {
		return nil
	}

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
		dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d",
			conf.Connect.Host,
			conf.Connect.Username,
			conf.Connect.Password,
			conf.Connect.Port,
		)
		if conf.Connect.DBName != "" {
			dsn += fmt.Sprintf(" dbname=%s %s", conf.Connect.DBName, conf.Connect.Option)
		}
		return postgres.Open(dsn)
	case _sqlServer:
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d",
			conf.Connect.Username,
			conf.Connect.Password,
			conf.Connect.Host,
			conf.Connect.Port,
		)
		if conf.Connect.DBName != "" {
			dsn += fmt.Sprintf("?database=%s", conf.Connect.DBName)
		}

		return sqlserver.Open(dsn)
	default:
		return nil
	}
}

func (d *db) create(conf *config.Database) error {
	copyConf := *conf
	copyConf.Connect.DBName = ""
	copyConf.Connect.Option = ""

	connect, err := gorm.Open(d.open(&copyConf))
	if err != nil {
		return err
	}

	_ = connect.Session(&gorm.Session{
		Logger: glogger.Default.LogMode(glogger.Silent),
	}).Exec(fmt.Sprintf("CREATE DATABASE %s", conf.Connect.DBName)).Error

	return nil
}
