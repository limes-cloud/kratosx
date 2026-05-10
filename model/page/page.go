package page

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Scopes func(db *gorm.DB) *gorm.DB

type Search struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
}

// SearchScopes 搜索排序
func SearchScopes(db *gorm.DB, req *Search) *gorm.DB {
	// 分页查询数据
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 10
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	var (
		orderBy = "id"
		isDesc  = false
	)

	if req.OrderBy != nil && *req.OrderBy != "" && isValidColumn(*req.OrderBy) {
		orderBy = *req.OrderBy
	}

	if req.Order != nil && *req.Order != "" {
		isDesc = *req.Order == "desc"
	}

	return db.Order(clause.OrderByColumn{Column: clause.Column{Name: orderBy}, Desc: isDesc})
}

func isValidColumn(name string) bool {
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
