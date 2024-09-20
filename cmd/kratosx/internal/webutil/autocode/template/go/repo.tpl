package repository

import (
	"github.com/limes-cloud/kratosx"
	"{{.Module}}/internal/domain/entity"
    "{{.Module}}/internal/types"
)

type {{.ClassifyUpper}}Repository interface {
	// Get{{.Object}} 获取指定的{{.Title}}
	Get{{.Object}}(ctx kratosx.Context, id uint32) (*entity.{{.Object}}, error)

    {{- range $val := .GetByCodes}}
    // Get{{$.Object}}By{{$val.Method}} 通过{{$val.Method}}获取指定{{$.Title}}数据
     Get{{$.Object}}By{{$val.Method}}(ctx kratosx.Context,  {{$val.Params}}) (*entity.{{$.Object}}, error)
    {{- end}}

	// List{{.Object}} 获取{{.Title}}列表
	List{{.Object}}(ctx kratosx.Context, req *types.List{{.Object}}Request) ([]*entity.{{.Object}}, uint32, error)

	// Create{{.Object}} 创建{{.Title}}
	Create{{.Object}}(ctx kratosx.Context, req *entity.{{.Object}}) (uint32, error)

	// Update{{.Object}} 更新{{.Title}}
	Update{{.Object}}(ctx kratosx.Context, req *entity.{{.Object}}) error

    {{- if .EnableBatchDelete}}
	// Delete{{.Object}} 删除{{.Title}}
	Delete{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error)
    {{- else}}
    // Delete{{.Object}} 删除{{.Title}}
    Delete{{.Object}}(ctx kratosx.Context, id uint32) error
    {{- end}}

	// GetTrash{{.Object}} 获取指定的回收站{{.Title}}
	GetTrash{{.Object}}(ctx kratosx.Context, id uint32) (*entity.{{.Object}}, error)

	// ListTrash{{.Object}} 获取回收站{{.Title}}列表
	ListTrash{{.Object}}(ctx kratosx.Context, req *types.ListTrash{{.Object}}Request) ([]*entity.{{.Object}}, uint32, error)

    {{- if .EnableBatchDelete}}
	// DeleteTrash{{.Object}} 彻底删除{{.Title}}
	DeleteTrash{{.Object}}(ctx kratosx.Context, ids []uint32) (uint32, error)
    {{- else}}
    // DeleteTrash{{.Object}} 彻底删除{{.Title}}
	DeleteTrash{{.Object}}(ctx kratosx.Context, id uint32) error
    {{- end}}

	// RevertTrash{{.Object}} 还原{{.Title}}
	RevertTrash{{.Object}}(ctx kratosx.Context, id uint32) error
}
