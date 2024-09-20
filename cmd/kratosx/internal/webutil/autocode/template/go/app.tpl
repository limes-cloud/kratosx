package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"github.com/go-kratos/kratos/v2/transport/grpc"
    "github.com/go-kratos/kratos/v2/transport/http"

	"{{.Module}}/api/{{.Server}}/errors"
	pb "{{.Module}}/api/{{.Server}}/{{.ClassifyLowerCase}}/v1"
	"{{.Module}}/internal/domain/service"
	"{{.Module}}/internal/conf"
	"{{.Module}}/internal/types"
	"{{.Module}}/internal/domain/entity"
	"{{.Module}}/internal/infra/dbs"
)

type {{.Classify}}App struct {
	pb.Unimplemented{{.Classify}}Server
	srv *service.{{.Classify}}Service
}

func New{{.Classify}}App(conf *conf.Config) *{{.Classify}}App {
	return &{{.Classify}}App{
		srv: service.New{{.Classify}}Service(conf, dbs.New{{.Classify}}Infra()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := New{{.Classify}}App(c)
		pb.Register{{.Classify}}HTTPServer(hs, srv)
		pb.Register{{.Classify}}Server(gs, srv)
	})
}

// Get{{.Object}} 获取指定的{{.Title}}
func (s *{{.Classify}}App) Get{{.Object}}(c context.Context, req *pb.Get{{.Object}}Request) (*pb.Get{{.Object}}Reply, error) {

    {{if gt (len $.GetByCodes) 0}}
	var (
	    ctx = kratosx.MustContext(c)
	    ent *entity.{{.Object}}
	    err error
	)
    switch req.Params.(type) {
    case *pb.Get{{.Object}}Request_Id:
    	ent, err = s.srv.Get{{.Object}}(kratosx.MustContext(c), req.GetId())
    {{- range $val := .GetByCodes}}
    {{- if gt (len $val.Fields) 1}}
    case *pb.Get{{$.Object}}Request_{{$val.Method}}_:
        data := req.Get{{$val.Method}}()
    	ent, err = s.srv.Get{{$.Object}}By{{$val.Method}}(kratosx.MustContext(c) {{- range $p := $val.Fields}}, data.Get{{$p}}(){{- end}})
    {{- else}}
     case *pb.Get{{$.Object}}Request_{{$val.Method}}:
        ent, err = s.srv.Get{{$.Object}}By{{$val.Method}}(kratosx.MustContext(c), req.Get{{$val.Method}}())
    {{- end}}
    {{- end}}
    default:
    	return nil, errors.ParamsError()
    }
    if err != nil {
    	return nil, err
    }
    {{- else}}

	var ctx = kratosx.MustContext(c)

    ent, err := s.srv.Get{{.Object}}(ctx, req.Id)
    if err != nil {
    	return nil, err
    }
    {{- end}}

	reply := pb.Get{{.Object}}Reply{}
	if err := valx.Transform(ent, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// List{{.Object}} 获取{{.Title}}列表
func (s *{{.Classify}}App) List{{.Object}}(c context.Context, req *pb.List{{.Object}}Request) (*pb.List{{.Object}}Reply, error) {
	var ctx = kratosx.MustContext(c)
	result, total, err := s.srv.List{{.Object}}(ctx, &types.List{{.Object}}Request{
        Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
        {{- range $val := .QueryFields}}
        {{$val}}:req.{{$val}},
        {{- end}}
	})
	if err != nil {
		return nil, err
	}

	reply := pb.List{{.Object}}Reply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// Create{{.Object}} 创建{{.Title}}
func (s *{{.Classify}}App) Create{{.Object}}(c context.Context, req *pb.Create{{.Object}}Request) (*pb.Create{{.Object}}Reply, error) {
	var (
		ent  = entity.{{.Object}}{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.srv.Create{{.Object}}(ctx, &ent)
	if err != nil {
		return nil, err
	}

	return &pb.Create{{.Object}}Reply{Id: id}, nil
}


// Update{{.Object}} 更新{{.Title}}
func (s *{{.Classify}}App) Update{{.Object}}(c context.Context, req *pb.Update{{.Object}}Request) (*pb.Update{{.Object}}Reply, error) {
	var (
		ent  = entity.{{.Object}}{}
		ctx  = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.srv.Update{{.Object}}(ctx, &ent); err != nil {
		return nil, err
	}

	return &pb.Update{{.Object}}Reply{}, nil
}

{{- if .EnableBatchDelete}}
// Delete{{.Object}} 删除{{.Title}}
func (s *{{.Classify}}App) Delete{{.Object}}(c context.Context, req *pb.Delete{{.Object}}Request) (*pb.Delete{{.Object}}Reply, error) {
	total, err := s.srv.Delete{{.Object}}(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.Delete{{.Object}}Reply{Total: total}, nil
}
{{- else}}
// Delete{{.Object}} 删除{{.Title}}
func (s *{{.Classify}}App) Delete{{.Object}}(c context.Context, req *pb.Delete{{.Object}}Request) (*pb.Delete{{.Object}}Reply, error) {
	return &pb.Delete{{.Object}}Reply{}, s.srv.Delete{{.Object}}(kratosx.MustContext(c), req.Id)
}
{{- end}}

// GetTrash{{.Object}} 获取回收站指定的{{.Title}}
func (s *{{.Classify}}App) GetTrash{{.Object}}(c context.Context, req *pb.GetTrash{{.Object}}Request) (*pb.GetTrash{{.Object}}Reply, error) {
	var ctx = kratosx.MustContext(c)

	result, err := s.srv.GetTrash{{.Object}}(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	reply := pb.GetTrash{{.Object}}Reply{}
	if err := valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListTrash{{.Object}} 获取回收站{{.Title}}列表
func (s *{{.Classify}}App) ListTrash{{.Object}}(c context.Context, req *pb.ListTrash{{.Object}}Request) (*pb.ListTrash{{.Object}}Reply, error) {
	var ctx = kratosx.MustContext(c)

	result, total, err := s.srv.ListTrash{{.Object}}(ctx, &types.ListTrash{{.Object}}Request{
        Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
	    {{- range $val := .QueryFields}}
        {{$val}}:req.{{$val}},
        {{- end}}
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListTrash{{.Object}}Reply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}


{{- if .EnableBatchDelete}}
// DeleteTrash{{.Object}} 删除回收站{{.Title}}
func (s *{{.Classify}}App) DeleteTrash{{.Object}}(ctx context.Context, req *pb.DeleteTrash{{.Object}}Request) (*pb.DeleteTrash{{.Object}}Reply, error) {
	total, err := s.srv.DeleteTrash{{.Object}}(kratosx.MustContext(ctx), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteTrash{{.Object}}Reply{Total: total}, nil
}
{{- else}}
// DeleteTrash{{.Object}} 删除回收站{{.Title}}
func (s *{{.Classify}}App) DeleteTrash{{.Object}}(ctx context.Context, req *pb.DeleteTrash{{.Object}}Request) (*pb.DeleteTrash{{.Object}}Reply, error) {
	return &pb.DeleteTrash{{.Object}}Reply{}, s.srv.DeleteTrash{{.Object}}(kratosx.MustContext(ctx), req.Id)
}
{{- end}}

// RevertTrash{{.Object}} 还原回收站{{.Title}}
func (s *{{.Classify}}App) RevertTrash{{.Object}}(ctx context.Context, req *pb.RevertTrash{{.Object}}Request) (*pb.RevertTrash{{.Object}}Reply, error) {
	return &pb.RevertTrash{{.Object}}Reply{},s.srv.RevertTrash{{.Object}}(kratosx.MustContext(ctx), req.Id)
}
