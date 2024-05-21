package notice

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/conf"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// GetNotice 获取指定的通知
func (u *UseCase) GetNotice(ctx kratosx.Context, id uint32) (*Notice, error) {
	user, err := u.repo.GetNotice(ctx, id)
	if err != nil {
		ctx.Logger().Warnw("msg", "get error", "err", err.Error())
		return nil, errors.NotFound()
	}
	return user, nil
}

// ListNotice 获取通知列表
func (u *UseCase) ListNotice(ctx kratosx.Context, req *ListNoticeRequest) ([]*Notice, uint32, error) {
	list, total, err := u.repo.ListNotice(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list error", "err", err.Error())
		return nil, 0, errors.List()
	}
	return list, total, nil
}

// CreateNotice 创建通知
func (u *UseCase) CreateNotice(ctx kratosx.Context, user *Notice) (uint32, error) {
	id, err := u.repo.CreateNotice(ctx, user)
	if err != nil {
		ctx.Logger().Warnw("msg", "create error", "err", err.Error())
		return 0, errors.Create()
	}
	return id, nil
}

// ImportNotice 导入通知
func (u *UseCase) ImportNotice(ctx kratosx.Context, users []*Notice) (uint32, uint32, error) {
	it, ut, err := u.repo.ImportNotice(ctx, users)
	if err != nil {
		ctx.Logger().Warnw("msg", "import error", "err", err.Error())
		return 0, 0, errors.Import()
	}
	return it, ut, nil
}

// ExportNotice 导出通知
func (u *UseCase) ExportNotice(ctx kratosx.Context, req *ExportNoticeRequest) (string, error) {
	id, err := u.repo.ExportNotice(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "export error", "err", err.Error())
		return "", errors.Export()
	}
	return id, nil
}

// UpdateNotice 更新通知
func (u *UseCase) UpdateNotice(ctx kratosx.Context, user *Notice) error {
	if err := u.repo.UpdateNotice(ctx, user); err != nil {
		ctx.Logger().Warnw("msg", "update error", "err", err.Error())
		return errors.Update()
	}
	return nil
}

// DeleteNotice 删除通知
func (u *UseCase) DeleteNotice(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteNotice(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete error", "err", err.Error())
		return errors.Delete()
	}
	return nil
}

// BatchDeleteNotice 批量删除通知
func (u *UseCase) BatchDeleteNotice(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.BatchDeleteNotice(ctx, ids)
	if err != nil {
		ctx.Logger().Warnw("msg", "batch delete error", "err", err.Error())
		return 0, errors.BatchDelete()
	}
	return total, nil
}
