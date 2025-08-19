package model

type CreateModel struct {
	Id        uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	CreatedAt int64  `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
}

type BaseModel struct {
	Id        uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	CreatedAt int64  `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64  `json:"updatedAt,omitempty" gorm:"index;comment:修改时间"`
}

type DeleteModel struct {
	Id        uint32    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	CreatedAt int64     `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64     `json:"updatedAt,omitempty" gorm:"index;comment:修改时间"`
	DeletedAt NullInt64 `json:"deletedAt" gorm:"index;comment:删除时间"`
}
