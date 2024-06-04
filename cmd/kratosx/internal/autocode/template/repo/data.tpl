package data

import (
    "fmt"
	{{if .IsTree}} "errors" {{end}}

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"google.golang.org/protobuf/proto"

	biz "{{.Server}}/internal/biz/{{.Module}}"
	"{{.Server}}/internal/data/model"

)

type {{.ModuleLower}}Repo struct {
}

func New{{.ModuleUpper}}Repo() biz.Repo {
	return &{{.ModuleLower}}Repo{}
}

// To{{.Object}}Entity model转entity
func (r {{.ModuleLower}}Repo) To{{.Object}}Entity(m *model.{{.Object}}) *biz.{{.Object}} {
	e := &biz.{{.Object}}{}
    _ = valx.Transform(m, e)
    return e
}

// To{{.Object}}Model entity转model
func (r {{.ModuleLower}}Repo) To{{.Object}}Model(e *biz.{{.Object}}) *model.{{.Object}} {
	m := &model.{{.Object}}{}
    _ = valx.Transform(e, m)
    return m
}

{{- range $val := .ByCodes}}
// Get{{$.Object}}By{{$val.Method}} 获取指定数据
func (r {{$.ModuleLower}}Repo) Get{{$.Object}}By{{$val.Method}}(ctx kratosx.Context, {{$val.Params}}) (*biz.{{$.Object}}, error) {
	var (
		m  = model.{{$.Object}}{}
		fs = []string{ {{$.GetFields}} }
	)
    db := ctx.DB().Select(fs)
    {{- if $.HasGetPreload}}
    db = db.{{$.GetPreload}}
    {{- end}}
	if err := db.{{$val.Where}}.First(&m).Error; err != nil {
		return nil, err
	}

	return r.To{{$.Object}}Entity(&m), nil
}
{{- end}}

// Get{{.Object}} 获取指定的数据
func (r {{.ModuleLower}}Repo) Get{{.Object}}(ctx kratosx.Context, id uint32) (*biz.{{.Object}}, error) {
	var (
		m  = model.{{.Object}}{}
		fs = []string{ {{.GetFields}} }
	)
    db := ctx.DB().Select(fs)
    {{- if .HasGetPreload}}
    db = db.{{.GetPreload}}
    {{- end}}
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return r.To{{.Object}}Entity(&m), nil
}

// List{{.Object}} 获取列表
func (r {{.ModuleLower}}Repo) List{{.Object}}(ctx kratosx.Context, req *biz.List{{.Object}}Request) ([]*biz.{{.Object}}, uint32, error) {
	var (
		bs    []*biz.{{.Object}}
		ms    []*model.{{.Object}}
		total int64
		fs    = []string{ {{.ListFields}} }
	)

	db := ctx.DB().Model(model.{{.Object}}{}).Select(fs)

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

	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.To{{.Object}}Entity(m))
	}
	return bs, uint32(total), nil
}

{{- if not .IsTree}}
// Create{{.Object}} 创建数据
func (r {{.ModuleLower}}Repo) Create{{.Object}}(ctx kratosx.Context, req *biz.{{.Object}}) (uint32, error) {
	m := r.To{{.Object}}Model(req)
	return m.Id, ctx.DB().Create(m).Error
}
{{- else}}
// Create{{.Object}} 创建数据
func (r {{.ModuleLower}}Repo) Create{{.Object}}(ctx kratosx.Context, req *biz.{{.Object}}) (uint32, error) {
	m := r.To{{.Object}}Model(req)
    return m.Id, ctx.Transaction(func(ctx kratosx.Context) error {
    	if err := ctx.DB().Create(m).Error; err != nil {
    		return err
    	}
    	return r.append{{.Object}}Children(ctx, req.ParentId, m.Id)
   })
}
{{- end }}

// Import{{.Object}} 导入数据
func (r {{.ModuleLower}}Repo) Import{{.Object}}(ctx kratosx.Context, req []*biz.{{.Object}}) (uint32, error) {
	var (
		ms []*model.{{.Object}}
	)

	for _, item := range req {
		ms = append(ms, r.To{{.Object}}Model(item))
	}

	db := ctx.DB().Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(ms, 1000)
	return uint32(len(req)), db.Error
}

// Export{{.Object}} 导出数据
func (r {{.ModuleLower}}Repo) Export{{.Object}}(ctx kratosx.Context, req *biz.Export{{.Object}}Request) (string, error) {
	return "", nil
}

{{- if not .IsTree}}
// Update{{.Object}} 更新数据
func (r {{.ModuleLower}}Repo) Update{{.Object}}(ctx kratosx.Context, req *biz.{{.Object}}) error {
	return ctx.DB().Updates(r.To{{.Object}}Model(req)).Error
}
{{- else}}
// Update{{.Object}} 更新数据
func (r {{.ModuleLower}}Repo) Update{{.Object}}(ctx kratosx.Context, req *biz.{{.Object}}) error {
	if req.Id == req.ParentId {
        return errors.New("父级不能为自己")
    }
    old, err := r.Get{{.Object}}(ctx, req.Id)
	if err != nil {
		return err
	}

	return ctx.Transaction(func(ctx kratosx.Context) error {
		if old.ParentId != req.ParentId {
			if err := r.remove{{.Object}}Parent(ctx, req.Id); err != nil {
				return err
			}
			if err := r.append{{.Object}}Children(ctx, req.ParentId, req.Id); err != nil {
				return err
			}
		}
		return ctx.DB().Updates(r.To{{.Object}}Model(req)).Error
	})
}
{{- end }}
// Update{{.Object}}Status 更新数据状态
func (r {{.ModuleLower}}Repo) Update{{.Object}}Status(ctx kratosx.Context, id uint32, status bool) error {
	return ctx.DB().Model(model.{{.Object}}{}).Where("id=?", id).Update("status", status).Error
}

{{- if not .IsTree}}
// Delete{{.Object}} 删除数据
func (r {{.ModuleLower}}Repo) Delete{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&model.{{.Object}}{})
	return uint32(db.RowsAffected), db.Error
}
{{- else}}
// Delete{{.Object}} 删除数据
func (r {{.ModuleLower}}Repo) Delete{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error) {
    var del []uint32
	for _, id := range ids {
		del = append(del, id)
		childrenIds, err := r.Get{{.Object}}ChildrenIds(ctx, id)
		if err != nil {
			return 0, err
		}
		del = append(del, childrenIds...)
	}
	db := ctx.DB().Where("id in ?", del).Delete(&model.{{.Object}}{})
	return uint32(db.RowsAffected), db.Error
}
{{- end}}

// GetTrash{{.Object}} 获取垃圾桶指定数据
func (r {{.ModuleLower}}Repo) GetTrash{{.Object}}(ctx kratosx.Context, id uint32) (*biz.{{.Object}}, error) {
	var (
		m  = model.{{.Object}}{}
		fs = []string{ {{.GetTrashFields}} }
	)

	if err := ctx.DB().Unscoped().Select(fs).First(&m, "id=? and deleted_at != 0",id).Error; err != nil {
		return nil, err
	}

	return r.To{{.Object}}Entity(&m), nil
}

// ListTrash{{.Object}} 获取垃圾桶列表
func (r {{.ModuleLower}}Repo) ListTrash{{.Object}}(ctx kratosx.Context, req *biz.ListTrash{{.Object}}Request) ([]*biz.{{.Object}}, uint32, error) {
	var (
		bs    []*biz.{{.Object}}
		ms    []*model.{{.Object}}
		total int64
		fs    = []string{ {{.ListTrashFields}} }
	)

	db := ctx.DB().Unscoped().Model(model.{{.Object}}{}).Select(fs)
    db = db.Where("deleted_at != 0")
	{{.QueryCodes}}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}else{
     	*req.OrderBy = *req.OrderBy + ",id"
    }
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}

	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))
	if err := db.Order(clause.OrderByColumn{
		Column: clause.Column{Name: *req.OrderBy},
		Desc:   *req.Order == "desc",
	}).Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.To{{.Object}}Entity(m))
	}
	return bs, uint32(total), nil
}

// DeleteTrash{{.Object}} 彻底删除数据
func (r {{.ModuleLower}}Repo) DeleteTrash{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Unscoped().Delete(model.{{.Object}}{}, "id in ?", ids)
	return uint32(db.RowsAffected), db.Error
}

// RevertTrash{{.Object}} 还原指定的数据
func (r {{.ModuleLower}}Repo) RevertTrash{{.Object}}(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Unscoped().Model(model.{{.Object}}{}).Where("id=?", id).Update("deleted_at", 0).Error
}

{{- if  .IsTree}}
// Get{{.Object}}ChildrenIds 获取指定id的所有子id
func (r {{.ModuleLower}}Repo) Get{{.Object}}ChildrenIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids,ctx.DB().Model(model.{{.Object}}Closure{}).
		Select("children").
		Where("parent=?", id).
		Scan(&ids).Error
}

// Get{{.Object}}ParentIds 获取指定id的所有父id
func (r {{.ModuleLower}}Repo) Get{{.Object}}ParentIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(model.{{.Object}}Closure{}).
		Select("parent").
		Where("children=?", id).
		Scan(&ids).Error
}

// append{{.Object}}Children 添加id到指定的父id下
func (r {{.ModuleLower}}Repo) append{{.Object}}Children(ctx kratosx.Context, pid uint32, id uint32) error {
	list := []*model.{{.Object}}Closure{
        {
            Parent:   pid,
            Children: id,
        },
	}
	ids, _ := r.Get{{.Object}}ParentIds(ctx, pid)
	for _, item := range ids {
		list = append(list, &model.{{.Object}}Closure{
			Parent:   item,
			Children: id,
		})
	}
	return ctx.DB().Create(&list).Error
}

// remove{{.Object}}Parent 删除指定id的所有父层级
func (r {{.ModuleLower}}Repo) remove{{.Object}}Parent(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&model.{{.Object}}Closure{}, "children=?", id).Error
}
{{- end }}