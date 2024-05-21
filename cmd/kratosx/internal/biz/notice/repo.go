package notice

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// GetNotice 获取指定的通知
	GetNotice(ctx kratosx.Context, id uint32) (*Notice, error)

	// ListNotice 获取通知列表
	ListNotice(ctx kratosx.Context, req *ListNoticeRequest) ([]*Notice, uint32, error)

	// CreateNotice 创建通知
	CreateNotice(ctx kratosx.Context, req *Notice) (uint32, error)

	// ImportNotice 导入通知
	ImportNotice(ctx kratosx.Context, req []*Notice) (uint32, uint32, error)

	// ExportNotice 导出通知
	ExportNotice(ctx kratosx.Context, req *ExportNoticeRequest) (string, error)

	// UpdateNotice 更新通知
	UpdateNotice(ctx kratosx.Context, req *Notice) error

	// DeleteNotice 删除通知
	DeleteNotice(ctx kratosx.Context, id uint32) error

	// BatchDeleteNotice 批量删除通知
	BatchDeleteNotice(ctx kratosx.Context, ids []uint32) (uint32, error)
}
