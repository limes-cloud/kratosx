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

type CreateTenantModel struct {
	Id        uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	TenantId  uint32 `json:"tenantId" gorm:"index;comment:租户ID;hook:tenant"`
	CreatedAt int64  `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
}

type BaseTenantModel struct {
	Id        uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	TenantId  uint32 `json:"tenantId" gorm:"index;comment:租户ID;hook:tenant"`
	CreatedAt int64  `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64  `json:"updatedAt,omitempty" gorm:"index;comment:修改时间"`
}

type DeleteTenantModel struct {
	Id        uint32    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	TenantId  uint32    `json:"tenantId" gorm:"index;comment:租户ID;hook:tenant"`
	CreatedAt int64     `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64     `json:"updatedAt,omitempty" gorm:"index;comment:修改时间"`
	DeletedAt NullInt64 `json:"deletedAt,omitempty" gorm:"index;comment:删除时间"`
}

type CreateTenantDeptModel struct {
	Id        uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	TenantId  uint32 `json:"tenantId" gorm:"index;comment:租户ID;hook:tenant"`
	DeptId    uint32 `json:"deptId" gorm:"index;comment:部门ID;hook:dept"`
	CreatedAt int64  `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
}

type BaseTenantDeptModel struct {
	Id        uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	TenantId  uint32 `json:"tenantId" gorm:"index;comment:租户ID;hook:tenant"`
	DeptId    uint32 `json:"deptId" gorm:"index;comment:部门ID;hook:dept"`
	CreatedAt int64  `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64  `json:"updatedAt,omitempty" gorm:"index;comment:修改时间"`
}

type DeleteTenantDeptModel struct {
	Id        uint32    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	TenantId  uint32    `json:"tenantId" gorm:"index;comment:租户ID;hook:tenant"`
	DeptId    uint32    `json:"deptId" gorm:"index;comment:部门ID;hook:dept"`
	CreatedAt int64     `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64     `json:"updatedAt,omitempty" gorm:"index;comment:修改时间"`
	DeletedAt NullInt64 `json:"deletedAt,omitempty" gorm:"index;comment:删除时间"`
}

type CreateTenantUserModel struct {
	Id        uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	TenantId  uint32 `json:"tenantId" gorm:"index;comment:租户ID;hook:tenant"`
	DeptId    uint32 `json:"deptId" gorm:"index;comment:部门ID;hook:dept"`
	UserId    uint32 `json:"userId" gorm:"index;comment:用户ID;hook:user"`
	CreatedAt int64  `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
}

type BaseTenantUserModel struct {
	Id        uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	TenantId  uint32 `json:"tenantId" gorm:"index;comment:租户ID;hook:tenant"`
	DeptId    uint32 `json:"deptId" gorm:"index;comment:部门ID;hook:dept"`
	UserId    uint32 `json:"userId" gorm:"index;comment:用户ID;hook:user"`
	CreatedAt int64  `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64  `json:"updatedAt,omitempty" gorm:"index;comment:修改时间"`
}

type DeleteTenantUserModel struct {
	Id        uint32    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键ID"`
	TenantId  uint32    `json:"tenantId" gorm:"index;comment:租户ID;hook:tenant"`
	DeptId    uint32    `json:"deptId" gorm:"index;comment:部门ID;hook:dept"`
	UserId    uint32    `json:"userId" gorm:"index;comment:用户ID;hook:user"`
	CreatedAt int64     `json:"createdAt,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64     `json:"updatedAt,omitempty" gorm:"index;comment:修改时间"`
	DeletedAt NullInt64 `json:"deletedAt,omitempty" gorm:"index;comment:删除时间"`
}
