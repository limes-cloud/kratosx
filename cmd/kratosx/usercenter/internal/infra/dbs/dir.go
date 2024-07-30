package data

import (
	"fmt"
	"github.com/limescloud/usercenter/internal/types"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"

	"github.com/limescloud/usercenter/internal/entity"
)

type DictionaryInfra struct {
}

func NewDictionaryInfra() *DictionaryInfra {
	return &DictionaryInfra{}
}
// GetDictionaryByKeyword 获取指定数据
func (r *DictionaryInfra) GetDictionaryByKeyword(ctx kratosx.Context, keyword string) (*entity.Dictionary, error) {
	var (
		ent  = entity.Dictionary{}
		fs = []string{ "*" }
	)
	db := ctx.DB().Select(fs).Preload("Users")
	return &ent, db.Where("keyword = ?",keyword).First(&m).Error
}

// GetDictionary 获取指定的数据
func (r *DictionaryInfra) GetDictionary(ctx kratosx.Context, id uint32) (*entity.Dictionary, error) {
	var (
		m  = entity.Dictionary{}
		fs = []string{ "*" }
	)
	db := ctx.DB().Select(fs).Preload("Users")
	return r.ToDictionaryEntity(&m), db.First(&m, id).Error
}

// ListDictionary 获取列表
func (r *DictionaryInfra) ListDictionary(ctx kratosx.Context, req *types.ListDictionaryRequest) ([]*entity.Dictionary, uint32, error) {
	var (
		list    []*entity.Dictionary
		fs    = []string{ "*" }
		total int64
	)

	db := ctx.DB().Model(entity.Dictionary{})
	db = db.Select(fs)db = db.Preload("Users")

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id asc")
	}

	return &list, uint32(total), db.Find(&ms).Error
}
// CreateDictionary 创建数据
func (r *DictionaryInfra) CreateDictionary(ctx kratosx.Context, ent *entity.Dictionary) (uint32, error) {
	return m.Id, ctx.DB().Create(ent).Error
}
// UpdateDictionary 更新数据
func (r *DictionaryInfra) UpdateDictionary(ctx kratosx.Context, ent *entity.Dictionary) error {
	return ctx.DB().Updates(ent).Error
}
// DeleteDictionary 删除数据
func (r *DictionaryInfra) DeleteDictionary(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&entity.Dictionary{})
	return uint32(db.RowsAffected), db.Error
}

// GetTrashDictionary 获取垃圾桶指定数据
func (r *DictionaryInfra) GetTrashDictionary(ctx kratosx.Context, id uint32) (*entity.Dictionary, error) {
	var (
		ent  = entity.Dictionary{}
		fs = []string{ "*" }
	)

	return &ent, ctx.DB().Unscoped().Where("deleted_at != 0").Select(fs).First(&m, "id = ?",id).Error
}

// ListTrashDictionary 获取垃圾桶列表
func (r *DictionaryInfra) ListTrashDictionary(ctx kratosx.Context, req *entity.ListTrashDictionaryRequest) ([]*entity.Dictionary, uint32, error) {
	var (
		list    []*entity.Dictionary
		fs    = []string{ "*" }
		total int64
	)

	db := ctx.DB().Model(entity.Dictionary{}).Unscoped()
	db = db.Select(fs)
	db = db.Where("deleted_at != 0")
	db = db.Preload("Users")

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id asc")
	}

	return &list, uint32(total), db.Find(&ms).Error
}
// DeleteTrashDictionary 彻底删除数据
func (r *DictionaryInfra) DeleteTrashDictionary(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Unscoped().Delete(entity.Dictionary{}, "id in ?", ids)
	return uint32(db.RowsAffected), db.Error
}

// RevertTrashDictionary 还原指定的数据
func (r *DictionaryInfra) RevertTrashDictionary(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Unscoped().Model(entity.Dictionary{}).Where("id=?", id).Update("deleted_at", 0).Error
}