package router

import (
	"github.com/go-kratos/kratos/v2/transport/http"

	autocode "github.com/limes-cloud/kratosx/cmd/kratosx/internal/webutil/autocode/router"
)

func NewRouter(srv *http.Server) {
	r := srv.Route("/autocode")

	// autocode 注册
	autocode.Register(r)
}
