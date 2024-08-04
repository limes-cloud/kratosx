package data

import (
    "fmt"
	{{if .IsTree}} "errors" {{end}}

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"

	"{{.Module}}/internal/domain/entity"
	"{{.Module}}/internal/types"
)

type {{.Classify}}Infra struct {
}

func New{{.Classify}}Infra() *{{$.Classify}}Infra {
	return &{{$.Classify}}Infra{}
}


{{- range $val := .GetByCodes}}
// Get{{$.Object}}By{{$val.Method}} 通过{{$val.Method}}获取指定{{$.Title}}数据
func (r *{{$.Classify}}Infra) Get{{$.Object}}By{{$val.Method}}(ctx kratosx.Context, {{$val.Params}}) (*entity.{{$.Object}}, error) {
	var (
		ent  = entity.{{$.Object}}{}
		fs = []string{ {{$.GetFields}} }
	)
    db := ctx.DB().Select(fs){{- if $.HasGetPreload}}.{{$.GetPreload}}{{- end}}
	return &ent, db.{{$val.Where}}.First(&ent).Error
}
{{- end}}

// Get{{.Object}} 获取指定的{{.Title}}数据
func (r *{{$.Classify}}Infra) Get{{.Object}}(ctx kratosx.Context, id uint32) (*entity.{{.Object}}, error) {
	var (
		ent  = entity.{{.Object}}{}
		fs = []string{ {{.GetFields}} }
	)
    db := ctx.DB().Select(fs){{- if .HasGetPreload}}.{{.GetPreload}}{{- end}}
	return &ent, db.First(&ent, id).Error
}

// List{{.Object}} 获取{{.Title}}列表
func (r *{{$.Classify}}Infra) List{{.Object}}(ctx kratosx.Context, req *types.List{{.Object}}Request) ([]*entity.{{.Object}}, uint32, error) {
	var (
		list    []*entity.{{.Object}}
		fs    = []string{ {{.ListFields}} }
		total int64
	)

	db := ctx.DB().Model(entity.{{.Object}}{})
	db = db.Select(fs)
	{{- if .HasGetPreload}}
	db = db.{{.ListPreload}}
	{{- end}}

	{{.QueryCodes}}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

    {{- if not .IsTree}}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))
	{{- end}}

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

	return list, uint32(total), db.Find(&list).Error
}

{{- if not .IsTree}}
// Create{{.Object}} 创建{{.Title}}数据
func (r *{{$.Classify}}Infra) Create{{.Object}}(ctx kratosx.Context, ent *entity.{{.Object}}) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}
{{- else}}
// Create{{.Object}} 创建{{.Title}}数据
func (r *{{$.Classify}}Infra) Create{{.Object}}(ctx kratosx.Context, ent *entity.{{.Object}}) (uint32, error) {
    return m.Id, ctx.Transaction(func(ctx kratosx.Context) error {
    	if err := ctx.DB().Create(ent).Error; err != nil {
    		return err
    	}
    	return r.append{{.Object}}Children(ctx, ent.ParentId, ent.Id)
   })
}
{{- end }}


{{- if not .IsTree}}
// Update{{.Object}} 更新{{.Title}}数据
func (r *{{$.Classify}}Infra) Update{{.Object}}(ctx kratosx.Context, ent *entity.{{.Object}}) error {
	return ctx.DB().Updates(ent).Error
}
{{- else}}
// Update{{.Object}} 更新{{.Title}}数据
func (r *{{$.Classify}}Infra) Update{{.Object}}(ctx kratosx.Context, ent *entity.{{.Object}}) error {
	if ent.Id == ent.ParentId {
        return errors.New("父级不能为自己")
    }
    old, err := r.Get{{.Object}}(ctx, ent.Id)
	if err != nil {
		return err
	}

	return ctx.Transaction(func(ctx kratosx.Context) error {
		if old.ParentId != ent.ParentId {
			if err := r.remove{{.Object}}Parent(ctx, ent.Id); err != nil {
				return err
			}
			if err := r.append{{.Object}}Children(ctx, ent.ParentId, ent.Id); err != nil {
				return err
			}
		}
		return ctx.DB().Updates(ent).Error
	})
}
{{- end }}

{{- if not .IsTree}}
{{- if .EnableBatchDelete}}
// Delete{{.Object}} 删除{{.Title}}数据
func (r *{{$.Classify}}Infra) Delete{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&entity.{{.Object}}{})
	return uint32(db.RowsAffected), db.Error
}
{{- else}}
// Delete{{.Object}} 删除{{.Title}}数据
func (r *{{$.Classify}}Infra) Delete{{.Object}}(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&entity.{{.Object}}{},id).Error
}
{{- end}}
{{- else}}
{{- if .EnableBatchDelete}}
// Delete{{.Object}} 删除{{.Title}}数据
func (r *{{$.Classify}}Infra) Delete{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error) {
    var del []uint32
	for _, id := range ids {
		del = append(del, id)
		childrenIds, err := r.Get{{.Object}}ChildrenIds(ctx, id)
		if err != nil {
			return 0, err
		}
		del = append(del, childrenIds...)
	}
	db := ctx.DB().Where("id in ?", del).Delete(&entity.{{.Object}}{})
	return uint32(db.RowsAffected), db.Error
}
{{- else}}
// Delete{{.Object}} 删除{{.Title}}数据
func (r *{{$.Classify}}Infra) Delete{{.Object}}(ctx kratosx.Context, id uint32) error {
    del, err := r.Get{{.Object}}ChildrenIds(ctx, id)
    if err != nil {
    	return 0, err
    }
	del = append(del, id)
	return ctx.DB().Where("id in ?", del).Delete(&entity.{{.Object}}{}).Error
}
{{- end}}
{{- end}}

// GetTrash{{.Object}} 获取垃圾桶指定{{.Title}}数据
func (r *{{$.Classify}}Infra) GetTrash{{.Object}}(ctx kratosx.Context, id uint32) (*entity.{{.Object}}, error) {
	var (
		ent  = entity.{{.Object}}{}
		fs = []string{ {{.GetTrashFields}} }
	)

	return &ent, ctx.DB().Unscoped().Where("deleted_at != 0").Select(fs).First(&ent, "id = ?",id).Error
}

// ListTrash{{.Object}} 获取垃圾桶{{.Title}}列表
func (r *{{$.Classify}}Infra) ListTrash{{.Object}}(ctx kratosx.Context, req *types.ListTrash{{.Object}}Request) ([]*entity.{{.Object}}, uint32, error) {
	var (
		list    []*entity.{{.Object}}
		fs    = []string{ {{.ListFields}} }
		total int64
	)

	db := ctx.DB().Model(entity.{{.Object}}{}).Unscoped()
	db = db.Select(fs)
	db = db.Where("deleted_at != 0")
	{{- if .HasGetPreload}}
	db = db.{{.ListPreload}}
	{{- end}}

	{{.QueryCodes}}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

    {{- if not .IsTree}}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))
	{{- end}}

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

	return list, uint32(total), db.Find(&list).Error
}

{{- if .EnableBatchDelete}}
// DeleteTrash{{.Object}} 彻底删除{{.Title}}数据
func (r *{{$.Classify}}Infra) DeleteTrash{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Unscoped().Delete(entity.{{.Object}}{}, "id in ?", ids)
	return uint32(db.RowsAffected), db.Error
}
{{- else}}
// DeleteTrash{{.Object}} 彻底删除{{.Title}}数据
func (r *{{$.Classify}}Infra) DeleteTrash{{.Object}}(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Unscoped().Delete(entity.{{.Object}}{}, id).Error
}
{{- end}}

// RevertTrash{{.Object}} 还原指定的{{.Title}}数据
func (r *{{$.Classify}}Infra) RevertTrash{{.Object}}(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Unscoped().Model(entity.{{.Object}}{}).Where("id=?", id).Update("deleted_at", 0).Error
}

{{- if  .IsTree}}
// Get{{.Object}}ChildrenIds 获取{{.Title}}指定id的所有子id
func (r *{{$.Classify}}Infra) Get{{.Object}}ChildrenIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids,ctx.DB().Model(entity.{{.Object}}Closure{}).
		Select("children").
		Where("parent=?", id).
		Scan(&ids).Error
}

// Get{{.Object}}ParentIds 获取{{.Title}}指定id的所有父id
func (r *{{$.Classify}}Infra) Get{{.Object}}ParentIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.{{.Object}}Closure{}).
		Select("parent").
		Where("children=?", id).
		Scan(&ids).Error
}

// append{{.Object}}Children 添加{{.Title}}id到指定的父id下
func (r *{{$.Classify}}Infra) append{{.Object}}Children(ctx kratosx.Context, pid uint32, id uint32) error {
	list := []*entity.{{.Object}}Closure{
        {
            Parent:   pid,
            Children: id,
        },
	}
	ids, _ := r.Get{{.Object}}ParentIds(ctx, pid)
	for _, item := range ids {
		list = append(list, &entity.{{.Object}}Closure{
			Parent:   item,
			Children: id,
		})
	}
	return ctx.DB().Create(&list).Error
}

// remove{{.Object}}Parent 删除{{.Title}}指定id的所有父层级
func (r *{{$.Classify}}Infra) remove{{.Object}}Parent(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&entity.{{.Object}}Closure{}, "children=?", id).Error
}
{{- end }}