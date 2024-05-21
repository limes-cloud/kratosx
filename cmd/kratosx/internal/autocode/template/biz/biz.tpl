package {{.Module}}

import (
	"github.com/limes-cloud/kratosx"

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

// GetNotice 获取指定的通知
func (u *UseCase) Get{{.Object}}(ctx kratosx.Context, id uint32) (*{{.Object}}, error) {
	user, err := u.repo.Get{{.Object}}(ctx, id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get error", "err", err.Error())
		return nil, errors.NotFound()
	}
	return user, nil
}

// ListNotice 获取通知列表
func (u *UseCase) List{{.Object}}(ctx kratosx.Context, req *List{{.Object}}Request) ([]*{{.Object}}, uint32, error) {
	list, total, err := u.repo.List{{.Object}}(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list error", "err", err.Error())
		return nil, 0, errors.List()
	}
	return list, total, nil
}

// CreateNotice 创建通知
func (u *UseCase) Create{{.Object}}(ctx kratosx.Context, user *{{.Object}}) (uint32, error) {
	id, err := u.repo.Create{{.Object}}(ctx, user)
	if err != nil {
		ctx.Logger().Warnw("msg", "create error", "err", err.Error())
		return 0, errors.Create()
	}
	return id, nil
}

// ImportNotice 导入通知
func (u *UseCase) Import{{.Object}}(ctx kratosx.Context, users []*{{.Object}}) (uint32, uint32, error) {
	it, ut, err := u.repo.Import{{.Object}}(ctx, users)
	if err != nil {
		ctx.Logger().Warnw("msg", "import error", "err", err.Error())
		return 0, 0, errors.Import()
	}
	return it, ut, nil
}

// ExportNotice 导出通知
func (u *UseCase) Export{{.Object}}(ctx kratosx.Context, req *Export{{.Object}}Request) (string, error) {
	id, err := u.repo.Export{{.Object}}(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "export error", "err", err.Error())
		return "", errors.Export()
	}
	return id, nil
}

// UpdateNotice 更新通知
func (u *UseCase) Update{{.Object}}(ctx kratosx.Context, user *{{.Object}}) error {
	if err := u.repo.Update{{.Object}}(ctx, user); err != nil {
		ctx.Logger().Warnw("msg", "update error", "err", err.Error())
		return errors.Update()
	}
	return nil
}

// DeleteNotice 删除通知
func (u *UseCase) Delete{{.Object}}(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete{{.Object}}(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete error", "err", err.Error())
		return errors.Delete()
	}
	return nil
}

// BatchDeleteNotice 批量删除通知
func (u *UseCase) BatchDelete{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.BatchDelete{{.Object}}(ctx, ids)
	if err != nil {
		ctx.Logger().Warnw("msg", "batch delete error", "err", err.Error())
		return 0, errors.BatchDelete()
	}
	return total, nil
}
