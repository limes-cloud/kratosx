package handler

import (
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/service"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/types"
)

// ListModel 获取全部的数据库模型
func ListModel(ctx http.Context) error {
	return nil
}

// ConnectDatabase 连接db数据库
func ConnectDatabase(ctx http.Context) error {
	var req types.ConnectDatabaseRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	return service.ConnectDatabase(&req)
}
