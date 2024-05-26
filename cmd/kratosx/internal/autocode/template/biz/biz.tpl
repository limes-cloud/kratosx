package {{.Module}}

import (
	"github.com/limes-cloud/kratosx"
    {{if .IsTree}}"github.com/limes-cloud/kratosx/pkg/tree"{{end}}
	"{{.Server}}/api/{{.ServerName}}/errors"
	"{{.Server}}/internal/conf"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// Get{{.Object}} 获取指定的{{.Title}} 
func (u *UseCase) Get{{.Object}}(ctx kratosx.Context, req *Get{{.Object}}Request) (*{{.Object}}, error) {
	var (
    	res *Dictionary
    	err error
    )

    if req.Id != nil {
    	res, err = u.repo.Get{{.Object}}(ctx, *req.Id)
    }else {{.GetCodes}} else{
        err = errors.ParamsError()
    }


	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}

{{- if .IsTree}}
// List{{.Object}} 获取{{.Title}} 列表树
func (u *UseCase) List{{.Object}}(ctx kratosx.Context, req *List{{.Object}}Request) ([]tree.Tree, uint32, error) {
	list, total, err := u.repo.List{{.Object}}(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return tree.BuildArrayTree(tree.ToTree(list)), total, nil
}
{{- else}}
// List{{.Object}} 获取{{.Title}} 列表
func (u *UseCase) List{{.Object}}(ctx kratosx.Context, req *List{{.Object}}Request) ([]*{{.Object}}, uint32, error) {
	list, total, err := u.repo.List{{.Object}}(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}
{{- end}}

// Create{{.Object}} 创建{{.Title}} 
func (u *UseCase) Create{{.Object}}(ctx kratosx.Context, user *{{.Object}}) (uint32, error) {
	id, err := u.repo.Create{{.Object}}(ctx, user)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// Import{{.Object}} 导入{{.Title}} 
func (u *UseCase) Import{{.Object}}(ctx kratosx.Context, users []*{{.Object}}) (uint32, error) {
	total, err := u.repo.Import{{.Object}}(ctx, users)
	if err != nil {
		return  0, errors.ImportError(err.Error())
	}
	return total, nil
}

// Export{{.Object}} 导出{{.Title}} 
func (u *UseCase) Export{{.Object}}(ctx kratosx.Context, req *Export{{.Object}}Request) (string, error) {
	id, err := u.repo.Export{{.Object}}(ctx, req)
	if err != nil {
		return "", errors.ExportError(err.Error())
	}
	return id, nil
}

// Update{{.Object}} 更新{{.Title}} 
func (u *UseCase) Update{{.Object}}(ctx kratosx.Context, user *{{.Object}}) error {
	if err := u.repo.Update{{.Object}}(ctx, user); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// Update{{.Object}}Status 更新{{.Title}}状态
func (u *UseCase) Update{{.Object}}Status(ctx kratosx.Context, id uint32, status bool) error {
	if err := u.repo.Update{{.Object}}Status(ctx, id, status); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// Delete{{.Object}} 删除{{.Title}}
func (u *UseCase) Delete{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.Delete{{.Object}}(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}

// GetTrash{{.Object}} 获取指定的{{.Title}}
func (u *UseCase) GetTrash{{.Object}}(ctx kratosx.Context, id uint32) (*{{.Object}}, error) {
	user, err := u.repo.GetTrash{{.Object}}(ctx, id)
	if err != nil {
		return nil, errors.GetTrashError(err.Error())
	}
	return user, nil
}

{{- if .IsTree}}
// ListTrash{{.Object}} 获取{{.Title}} 列表树
func (u *UseCase) ListTrash{{.Object}}(ctx kratosx.Context, req *List{{.Object}}Request) ([]tree.Tree, uint32, error) {
	list, total, err := u.repo.ListTrash{{.Object}}(ctx, req)
	if err != nil {
		return nil, 0, errors.ListTrashError(err.Error())
	}
	return tree.BuildArrayTree(tree.ToTree(list)), total, nil
}
{{- else}}
// ListTrash{{.Object}} 获取{{.Title}} 列表
func (u *UseCase) ListTrash{{.Object}}(ctx kratosx.Context, req *ListTrash{{.Object}}Request) ([]*{{.Object}}, uint32, error) {
	list, total, err := u.repo.ListTrash{{.Object}}(ctx, req)
	if err != nil {
		return nil, 0, errors.ListTrashError(err.Error())
	}
	return list, total, nil
}
{{- end}}

// DeleteTrash{{.Object}} 彻底删除{{.Title}}
func (u *UseCase) DeleteTrash{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteTrash{{.Object}}(ctx, ids)
	if err != nil {
		return 0, errors.DeleteTrashError(err.Error())
	}
	return total, nil
}

// RevertTrash{{.Object}} 还原删除{{.Title}}
func (u *UseCase) RevertTrash{{.Object}}(ctx kratosx.Context, id uint32) error {
	if err := u.repo.RevertTrash{{.Object}}(ctx, id);err != nil {
		return  errors.RevertTrashError(err.Error())
	}
	return nil
}