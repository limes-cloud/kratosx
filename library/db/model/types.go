package model

// DatabaseType 数据库类型枚举
type DatabaseType string

const (
	MySQL    DatabaseType = "mysql"
	Postgres DatabaseType = "postgresql"
)

// Initializer 数据库初始化接口
type Initializer interface {
	Exec() error                     // 执行初始化
	SetForce(force bool) Initializer // 设置强制标志
	SetPath(path string) Initializer // 设置SQL文件路径
}
