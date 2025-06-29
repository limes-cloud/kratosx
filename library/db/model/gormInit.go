package model

/*
GormInit 是一个用于记录系统初始化状态的数据库模型，它的核心作用是避免重复执行初始化 SQL
*/
type GormInit struct {
	Id   uint32 `gorm:"primaryKey"` // 固定主键ID=1
	Init bool   // 是否已初始化
}
