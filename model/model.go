package model

type CreateModel struct {
	Id        uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	CreatedAt int64  `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
}

func (CreateModel) Name() string {
	return "CreateModel"
}

type BaseModel struct {
	Id        uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	CreatedAt int64  `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64  `json:"updatedAt,omitempty" gorm:"index;comment:修改时间"`
}

func (BaseModel) Name() string {
	return "BaseModel"
}

type DeleteModel struct {
	Id        uint32    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	CreatedAt int64     `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64     `json:"updatedAt,omitempty" gorm:"index;comment:修改时间"`
	DeletedAt NullInt64 `json:"deletedAt" gorm:"index;comment:删除时间"`
}

type CreateDeptModel struct {
	Id        uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	DeptId    uint32 `json:"deptId" gorm:"index;comment:部门ID;hook:dept"`
	CreatedAt int64  `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
}

type BaseDeptModel struct {
	Id        uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	DeptId    uint32 `json:"deptId" gorm:"index;comment:部门ID;hook:dept"`
	CreatedAt int64  `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64  `json:"updatedAt,omitempty" gorm:"index;comment:修改时间"`
}

type DeleteDeptModel struct {
	Id        uint32    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	DeptId    uint32    `json:"deptId" gorm:"index;comment:部门ID;hook:dept"`
	CreatedAt int64     `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64     `json:"updatedAt,omitempty" gorm:"index;comment:修改时间"`
	DeletedAt NullInt64 `json:"deletedAt,omitempty" gorm:"index;comment:删除时间"`
}

type CreateUserModel struct {
	Id        uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	DeptId    uint32 `json:"deptId" gorm:"index;comment:部门ID;hook:dept"`
	UserId    uint32 `json:"userId" gorm:"index;comment:用户ID;hook:user"`
	CreatedAt int64  `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
}

type BaseUserModel struct {
	Id        uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	DeptId    uint32 `json:"deptId" gorm:"index;comment:部门ID;hook:dept"`
	UserId    uint32 `json:"userId" gorm:"index;comment:用户ID;hook:user"`
	CreatedAt int64  `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64  `json:"updatedAt,omitempty" gorm:"index;comment:修改时间"`
}

type DeleteUserModel struct {
	Id        uint32    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	DeptId    uint32    `json:"deptId" gorm:"index;comment:部门ID;hook:dept"`
	UserId    uint32    `json:"userId" gorm:"index;comment:用户ID;hook:user"`
	CreatedAt int64     `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64     `json:"updatedAt,omitempty" gorm:"index;comment:修改时间"`
	DeletedAt NullInt64 `json:"deletedAt,omitempty" gorm:"index;comment:删除时间"`
}
