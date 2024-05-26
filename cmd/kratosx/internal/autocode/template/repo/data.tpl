package data

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm/clause"

	biz "{{.Server}}/internal/biz/{{.Module}}"
	"{{.Server}}/internal/data/model"
)

type {{.ModuleLower}}Repo struct {
}

func New{{.ModuleUpper}}Repo() biz.Repo {
	return &{{.ModuleLower}}Repo{}
}

func (r {{.ModuleLower}}Repo) To{{.Object}}Entity(m *model.{{.Object}}) *biz.{{.Object}} {
	e := &biz.{{.Object}}{}
    _ = valx.Transform(m, e)
    return e
}

func (r {{.ModuleLower}}Repo) To{{.Object}}Model(e *biz.{{.Object}}) *model.{{.Object}} {
	m := &model.{{.Object}}{}
    _ = valx.Transform(e, m)
    return m
}

{{range $val := .ByCodes}}
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
{{end}}

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

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
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

func (r {{.ModuleLower}}Repo) Create{{.Object}}(ctx kratosx.Context, req *biz.{{.Object}}) (uint32, error) {
	m := r.To{{.Object}}Model(req)
	return m.Id, ctx.DB().Create(m).Error
}

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

func (r {{.ModuleLower}}Repo) Export{{.Object}}(ctx kratosx.Context, req *biz.Export{{.Object}}Request) (string, error) {
	return "", nil
}

func (r {{.ModuleLower}}Repo) Update{{.Object}}(ctx kratosx.Context, req *biz.{{.Object}}) error {
	return ctx.DB().Updates(r.To{{.Object}}Model(req)).Error
}

func (r {{.ModuleLower}}Repo) Update{{.Object}}Status(ctx kratosx.Context, id uint32, status bool) error {
	return ctx.DB().Model(model.{{.Object}}{}).Where("id=?", id).Update("status", status).Error
}

func (r {{.ModuleLower}}Repo) Delete{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&model.{{.Object}}{})
	return uint32(db.RowsAffected), db.Error
}

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

func (r {{.ModuleLower}}Repo) DeleteTrash{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Unscoped().Delete(model.{{.Object}}{}, "id in ?", ids)
	return uint32(db.RowsAffected), db.Error
}

func (r {{.ModuleLower}}Repo) RevertTrash{{.Object}}(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Unscoped().Model(model.{{.Object}}{}).Where("id=?", id).Update("deleted_at", 0).Error
}
