package {{.Module}}

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// Get{{.Object}} 获取指定的{{.Title}}
	Get{{.Object}}(ctx kratosx.Context, id uint32) (*{{.Object}}, error)

	// List{{.Object}} 获取{{.Title}}列表
	List{{.Object}}(ctx kratosx.Context, req *List{{.Object}}Request) ([]*{{.Object}}, uint32, error)

	// Create{{.Object}} 创建{{.Title}}
	Create{{.Object}}(ctx kratosx.Context, req *{{.Object}}) (uint32, error)

	// Import{{.Object}} 导入{{.Title}}
	Import{{.Object}}(ctx kratosx.Context, req []*{{.Object}}) (uint32, error)

	// Export{{.Object}} 导出{{.Title}}
	Export{{.Object}}(ctx kratosx.Context, req *Export{{.Object}}Request) (string, error)

	// Update{{.Object}} 更新{{.Title}}
	Update{{.Object}}(ctx kratosx.Context, req *{{.Object}}) error

	// Update{{.Object}}Status 更新{{.Title}}状态
	Update{{.Object}}Status (ctx kratosx.Context, id uint32, status bool) error

	// Delete{{.Object}} 删除{{.Title}}
	Delete{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error)

	// GetTrash{{.Object}} 获取指定的回收站{{.Title}}
	GetTrash{{.Object}}(ctx kratosx.Context, id uint32) (*{{.Object}}, error)

	// ListTrash{{.Object}} 获取回收站{{.Title}}列表
	ListTrash{{.Object}}(ctx kratosx.Context, req *ListTrash{{.Object}}Request) ([]*{{.Object}}, uint32, error)

	// DeleteTrash{{.Object}} 彻底删除{{.Title}}
	DeleteTrash{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error)

	// RevertTrash{{.Object}} 还原{{.Title}}
	RevertTrash{{.Object}}(ctx kratosx.Context, id uint32) error

	// Get{{.Object}}ParentIds 获取父{{.Title}}ID列表
	Get{{.Object}}ParentIds(ctx kratosx.Context, id uint32) ([]uint32, error)

	// Get{{.Object}}ChildrenIds 获取子{{.Title}}ID列表
    Get{{.Object}}ChildrenIds(ctx kratosx.Context, id uint32) ([]uint32, error)
}
