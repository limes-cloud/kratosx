package db

import (
	"fmt"
	"sort"
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

type Entity struct {
	Database string  `json:"database" gorm:"-"`
	Name     string  `json:"name"`
	Comment  string  `json:"comment"`
	Fields   []Field `json:"fields" gorm:"-"`
}

type Field struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type DB interface {
	// Get 获取指定名称的db实例，不指定名称则返回第一个如果实例不存在则返回nil
	Get(name ...string) *gorm.DB

	// List 获取gorm列表
	List() []*gorm.DB

	// TxKey 获取指定的事务key
	TxKey(name ...string) string

	// Entities 获取全部实体信息
	Entities() []*Entity
}

type db struct {
	cfs    map[string]*config.Database
	set    map[string]*gorm.DB
	drives map[string]string
	key    string
}

// List 数据数据库列表
func (d *db) List() []*gorm.DB {
	var (
		keys []string
		list []*gorm.DB
	)
	for key, _ := range d.set {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		if d.set[key] == nil {
			continue
		}
		list = append(list, d.set[key])
	}
	return list
}

func (d *db) Entities() []*Entity {
	var list []*Entity
	var keys []string
	for key := range d.set {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		if d.set[key] == nil {
			continue
		}
		list = append(list, d.loadEntities(d.set[key], d.drives[key])...)
	}
	return list
}

const (
	_mysql      = "mysql"
	_postgresql = "postgres"
	_sqlServer  = "sqlserver"
	_tidb       = "tidb"
)

var (
	ins *db

	once sync.Once
)

func Instance() DB {
	return ins
}

// Init 初始化全局db
func Init(cfs []*config.Database, opts ...Option) {
	if len(cfs) == 0 {
		return
	}

	once.Do(func() {
		o := &options{}
		for _, opt := range opts {
			opt(o)
		}

		ins = &db{
			set:    make(map[string]*gorm.DB),
			drives: make(map[string]string),
		}

		// 遍历配置连接数据库
		for ind, conf := range cfs {
			if err := ins.initFactory(conf, o); err != nil {
				panic("database init error :" + err.Error())
			}
			if ind == 0 {
				ins.key = conf.Name
			}
		}
	})
}

func (d *db) initFactory(conf *config.Database, opt *options) error {
	if conf.AutoCreate {
		if err := d.create(conf); err != nil {
			panic("auto create database error:" + err.Error())
		}
	}

	// 连接主数据库
	dialector, err := d.open(conf)
	if err != nil {
		return err
	}
	client, err := gorm.Open(dialector, &gorm.Config{
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

	sdb, err := client.DB()
	if err != nil || sdb == nil {
		return fmt.Errorf("get database instance error: %v", err)
	}
	sdb.SetConnMaxLifetime(conf.Config.MaxLifetime)
	sdb.SetMaxOpenConns(conf.Config.MaxOpenConn)
	sdb.SetMaxIdleConns(conf.Config.MaxIdleConn)

	// 注册hook
	registerHook(conf.Name, conf.Connect.DBName, client, opt.hook)

	d.set[conf.Name] = client
	d.drives[conf.Name] = conf.Drive
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

func (d *db) open(conf *config.Database) (gorm.Dialector, error) {
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
		return mysql.Open(dsn), nil
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
		return postgres.Open(dsn), nil
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

		return sqlserver.Open(dsn), nil
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", conf.Drive)
	}
}

func (d *db) create(conf *config.Database) error {
	if !isValidDBName(conf.Connect.DBName) {
		return fmt.Errorf("invalid database name: %s", conf.Connect.DBName)
	}

	copyConf := *conf
	copyConf.Connect.DBName = ""
	copyConf.Connect.Option = ""

	dialector, err := d.open(&copyConf)
	if err != nil {
		return err
	}
	connect, err := gorm.Open(dialector)
	if err != nil {
		return err
	}

	var sql string
	switch conf.Drive {
	case _mysql, _tidb:
		sql = fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", conf.Connect.DBName)
	case _postgresql:
		sql = fmt.Sprintf(`SELECT 'CREATE DATABASE "%s"' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '%s')`, conf.Connect.DBName, conf.Connect.DBName)
	case _sqlServer:
		sql = fmt.Sprintf("IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = '%s') CREATE DATABASE [%s]", conf.Connect.DBName, conf.Connect.DBName)
	default:
		return fmt.Errorf("unsupported driver for auto create database: %s", conf.Drive)
	}

	_ = connect.Session(&gorm.Session{
		Logger: glogger.Default.LogMode(glogger.Silent),
	}).Exec(sql).Error

	return nil
}

func isValidDBName(name string) bool {
	if name == "" {
		return false
	}
	for _, c := range name {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_') {
			return false
		}
	}
	return true
}

// loadEntities 加载实体
func (d *db) loadEntities(db *gorm.DB, drive string) []*Entity {
	var (
		tables []string
		list   []*Entity
	)

	database := db.Migrator().CurrentDatabase()

	switch drive {
	case _mysql, _tidb:
		db.Raw("SHOW TABLES").Scan(&tables)
		for _, table := range tables {
			entity := Entity{Database: database}
			db.Raw("SELECT table_name AS name, table_comment AS comment FROM information_schema.tables WHERE table_schema = ? AND table_name = ?", database, table).Scan(&entity)

			var columns []Field
			db.Raw("SELECT column_name AS name, column_comment AS comment FROM information_schema.columns WHERE table_schema = ? AND table_name = ?", database, table).Scan(&columns)
			entity.Fields = columns
			list = append(list, &entity)
		}
	case _postgresql:
		db.Raw("SELECT tablename FROM pg_tables WHERE schemaname = 'public'").Scan(&tables)
		for _, table := range tables {
			entity := Entity{Database: database, Name: table}

			var columns []Field
			db.Raw("SELECT column_name AS name, col_description((table_schema||'.'||table_name)::regclass, ordinal_position) AS comment FROM information_schema.columns WHERE table_schema = 'public' AND table_name = ?", table).Scan(&columns)
			entity.Fields = columns
			list = append(list, &entity)
		}
	case _sqlServer:
		db.Raw("SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_TYPE = 'BASE TABLE'").Scan(&tables)
		for _, table := range tables {
			entity := Entity{Database: database, Name: table}

			var columns []Field
			db.Raw("SELECT c.COLUMN_NAME AS name, CAST(ep.value AS NVARCHAR(500)) AS comment FROM INFORMATION_SCHEMA.COLUMNS c LEFT JOIN sys.extended_properties ep ON ep.major_id = OBJECT_ID(c.TABLE_SCHEMA + '.' + c.TABLE_NAME) AND ep.minor_id = c.ORDINAL_POSITION AND ep.name = 'MS_Description' WHERE c.TABLE_NAME = ?", table).Scan(&columns)
			entity.Fields = columns
			list = append(list, &entity)
		}
	}

	return list
}
