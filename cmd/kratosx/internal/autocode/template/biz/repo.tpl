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
	Import{{.Object}}(ctx kratosx.Context, req []*{{.Object}}) (uint32, uint32, error)

	// Export{{.Object}} 导出{{.Title}}
	Export{{.Object}}(ctx kratosx.Context, req *Export{{.Object}}Request) (string, error)

	// Update{{.Object}} 更新{{.Title}}
	Update{{.Object}}(ctx kratosx.Context, req *{{.Object}}) error

	// Delete{{.Object}} 删除{{.Title}}
	Delete{{.Object}}(ctx kratosx.Context, id uint32) error

	// BatchDelete{{.Object}} 批量删除{{.Title}}
	BatchDelete{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error)
}
