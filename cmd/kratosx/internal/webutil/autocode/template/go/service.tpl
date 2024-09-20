package service

import (
	"github.com/limes-cloud/kratosx"
    {{if .IsTree}}"github.com/limes-cloud/kratosx/pkg/tree"{{end}}
	"{{.Module}}/api/{{.Server}}/errors"
	"{{.Module}}/internal/conf"
	"{{.Module}}/internal/types"
	"{{.Module}}/internal/domain/repository"
	"{{.Module}}/internal/domain/entity"


)

type {{.Classify}}Service struct {
	conf *conf.Config
	repo repository.{{.Classify}}Repository
}

func New{{.Classify}}Service(conf *conf.Config, repo repository.{{.Classify}}Repository) *{{.Classify}}Service {
	return &{{.Classify}}Service{conf: conf, repo: repo}
}

{{- range $val := .GetByCodes}}
// Get{{$.Object}}By{{$val.Method}} 通过{{$val.Method}}获取指定{{$.Title}}数据
func (srv *{{$.Classify}}Service) Get{{$.Object}}By{{$val.Method}}(ctx kratosx.Context, {{$val.Params}}) (*entity.{{$.Object}}, error) {
	ent, err :=srv.repo.Get{{$.Object}}By{{$val.Method}}(ctx, {{$val.Fields}})
    if err != nil {
    	return nil, errors.GetError(err.Error())
    }
    return ent, nil
}
{{- end}}

// Get{{.Object}} 获取指定的{{.Title}} 
func (srv *{{.Classify}}Service) Get{{.Object}}(ctx kratosx.Context, id uint32) (*entity.{{.Object}}, error) {
    ent, err :=srv.repo.Get{{.Object}}(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return ent, nil
}

{{- if .IsTree}}
// List{{.Object}} 获取{{.Title}}列表树
func (srv *{{.Classify}}Service) List{{.Object}}(ctx kratosx.Context, req *types.List{{.Object}}Request) ([]*entity.{{.Object}}, uint32, error) {
	list, total, err :=srv.repo.List{{.Object}}(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return tree.BuildArrayTree(ts), total, nil
}
{{- else}}

// List{{.Object}} 获取{{.Title}}列表
func (srv *{{.Classify}}Service) List{{.Object}}(ctx kratosx.Context, req *types.List{{.Object}}Request) ([]*entity.{{.Object}}, uint32, error) {
	list, total, err :=srv.repo.List{{.Object}}(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}
{{- end}}

// Create{{.Object}} 创建{{.Title}} 
func (srv *{{.Classify}}Service) Create{{.Object}}(ctx kratosx.Context, req *entity.{{.Object}}) (uint32, error) {
	id, err :=srv.repo.Create{{.Object}}(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// Update{{.Object}} 更新{{.Title}} 
func (srv *{{.Classify}}Service) Update{{.Object}}(ctx kratosx.Context, ent *entity.{{.Object}}) error {
	if err :=srv.repo.Update{{.Object}}(ctx, ent); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

{{- if .EnableBatchDelete}}
// Delete{{.Object}} 删除{{.Title}}
func (srv *{{.Classify}}Service) Delete{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err :=srv.repo.Delete{{.Object}}(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}
{{- else}}
// Delete{{.Object}} 删除{{.Title}}
func (srv *{{.Classify}}Service) Delete{{.Object}}(ctx kratosx.Context, id uint32) error {
	if err :=srv.repo.Delete{{.Object}}(ctx, id); err != nil {
		return  errors.DeleteError(err.Error())
	}
	return  nil
}
{{- end}}

// GetTrash{{.Object}} 获取指定的{{.Title}}
func (srv *{{.Classify}}Service) GetTrash{{.Object}}(ctx kratosx.Context, id uint32) (*entity.{{.Object}}, error) {
	user, err :=srv.repo.GetTrash{{.Object}}(ctx, id)
	if err != nil {
		return nil, errors.GetTrashError(err.Error())
	}
	return user, nil
}

{{- if .IsTree}}
// ListTrash{{.Object}} 获取{{.Title}} 列表树
func (srv *{{.Classify}}Service) ListTrash{{.Object}}(ctx kratosx.Context, req *types.List{{.Object}}Request) ([]*entity.{{.Object}}, uint32, error) {
	list, total, err :=srv.repo.ListTrash{{.Object}}(ctx, req)
	if err != nil {
		return nil, 0, errors.ListTrashError(err.Error())
	}
	return tree.BuildArrayTree(tree.ToTree(list)), total, nil
}
{{- else}}

// ListTrash{{.Object}} 获取{{.Title}}列表
func (srv *{{.Classify}}Service) ListTrash{{.Object}}(ctx kratosx.Context, req *types.ListTrash{{.Object}}Request) ([]*entity.{{.Object}}, uint32, error) {
	list, total, err :=srv.repo.ListTrash{{.Object}}(ctx, req)
	if err != nil {
		return nil, 0, errors.ListTrashError(err.Error())
	}
	return list, total, nil
}
{{- end}}

{{- if .EnableBatchDelete}}
// DeleteTrash{{.Object}} 彻底删除{{.Title}}
func (srv *{{.Classify}}Service) DeleteTrash{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err :=srv.repo.DeleteTrash{{.Object}}(ctx, ids)
	if err != nil {
		return 0, errors.DeleteTrashError(err.Error())
	}
	return total, nil
}
{{- else}}
// DeleteTrash{{.Object}} 彻底删除{{.Title}}
func (srv *{{.Classify}}Service) DeleteTrash{{.Object}}(ctx kratosx.Context, id uint32) error {
	if err :=srv.repo.DeleteTrash{{.Object}}(ctx, id);err != nil {
		return errors.DeleteTrashError(err.Error())
	}
	return nil
}
{{- end}}

// RevertTrash{{.Object}} 还原删除{{.Title}}
func (srv *{{.Classify}}Service) RevertTrash{{.Object}}(ctx kratosx.Context, id uint32) error {
	if err :=srv.repo.RevertTrash{{.Object}}(ctx, id);err != nil {
		return  errors.RevertTrashError(err.Error())
	}
	return nil
}