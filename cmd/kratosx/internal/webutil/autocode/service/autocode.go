package service

import (
	"fmt"

	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/core"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/types"
)

const (
	_mysql      = "mysql"
	_postgresql = "postgresql"
	_sqlServer  = "sqlServer"
	_tidb       = "tidb"
	_clickhouse = "clickhouse"
)

func ConnectDatabase(req *types.ConnectDatabaseRequest) error {
	dial := open(req)
	if err := core.ConnectDatabase(dial); err != nil {
		return err
	}
	return nil
}

func open(conf *types.ConnectDatabaseRequest) gorm.Dialector {
	switch conf.Drive {
	case _mysql, _tidb:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.DBName,
		)
		return mysql.Open(dsn)
	case _postgresql:
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
			conf.Host,
			conf.Username,
			conf.Password,
			conf.DBName,
			conf.Port,
		)
		return postgres.Open(dsn)
	case _sqlServer:
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.DBName,
		)
		return sqlserver.Open(dsn)
	case _clickhouse:
		dsn := fmt.Sprintf("tcp://%s:%d?database=%s&username=%s&password=%s",
			conf.Host,
			conf.Port,
			conf.DBName,
			conf.Username,
			conf.Password,
		)
		return clickhouse.Open(dsn)
	default:
		return nil
	}
}
