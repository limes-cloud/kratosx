package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"github.com/go-kratos/kratos/v2/transport/grpc"
    "github.com/go-kratos/kratos/v2/transport/http"

	"{{.Server}}/api/{{.ServerName}}/errors"
	pb "{{.Server}}/api/{{.ServerName}}/{{.Module}}/v1"
	"{{.Server}}/internal/biz/{{.Module}}"
	"{{.Server}}/internal/conf"
	"{{.Server}}/internal/data"
)

type {{.ModuleUpper}}Service struct {
	pb.Unimplemented{{.ModuleUpper}}Server
	uc *{{.Module}}.UseCase
}

func New{{.ModuleUpper}}Service(conf *conf.Config) *{{.ModuleUpper}}Service {
	return &{{.ModuleUpper}}Service{
		uc: {{.Module}}.NewUseCase(conf, data.New{{.ModuleUpper}}Repo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := New{{.ModuleUpper}}Service(c)
		pb.Register{{.ModuleUpper}}HTTPServer(hs, srv)
		pb.Register{{.ModuleUpper}}Server(gs, srv)
	})
}

// Get{{.Object}} 获取指定的{{.Title}}
func (s *{{.ModuleUpper}}Service) Get{{.Object}}(c context.Context, req *pb.Get{{.Object}}Request) (*pb.Get{{.Object}}Reply, error) {
	var (
    	in  = {{.Module}}.Get{{.Object}}Request{}
    	ctx = kratosx.MustContext(c)
   )

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, err := s.uc.Get{{.Object}}(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.Get{{.Object}}Reply{}
	if err := valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// List{{.Object}} 获取{{.Title}}列表
func (s *{{.ModuleUpper}}Service) List{{.Object}}(c context.Context, req *pb.List{{.Object}}Request) (*pb.List{{.Object}}Reply, error) {
	var (
		in  = {{.Module}}.List{{.Object}}Request{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.List{{.Object}}(ctx, &in)
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
func (s *{{.ModuleUpper}}Service) Create{{.Object}}(c context.Context, req *pb.Create{{.Object}}Request) (*pb.Create{{.Object}}Reply, error) {
	var (
		in  = {{.Module}}.{{.Object}}{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.Create{{.Object}}(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.Create{{.Object}}Reply{Id: id}, nil
}

// Import{{.Object}} 导入{{.Title}}
func (s *{{.ModuleUpper}}Service) Import{{.Object}}(c context.Context, req *pb.Import{{.Object}}Request) (*pb.Import{{.Object}}Reply, error) {
	var (
		in  []*{{.Module}}.{{.Object}}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req.List, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	total, err := s.uc.Import{{.Object}}(ctx, in)
	if err != nil {
		return nil, err
	}

	return &pb.Import{{.Object}}Reply{Total: total}, nil
}

// Export{{.Object}} 导出{{.Title}}
func (s *{{.ModuleUpper}}Service) Export{{.Object}}(c context.Context, req *pb.Export{{.Object}}Request) (*pb.Export{{.Object}}Reply, error) {
	var (
		in  = {{.Module}}.Export{{.Object}}Request{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	src, err := s.uc.Export{{.Object}}(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.Export{{.Object}}Reply{Src: src}, nil
}

// Update{{.Object}} 更新{{.Title}}
func (s *{{.ModuleUpper}}Service) Update{{.Object}}(c context.Context, req *pb.Update{{.Object}}Request) (*pb.Update{{.Object}}Reply, error) {
	var (
		in  = {{.Module}}.{{.Object}}{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.Update{{.Object}}(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.Update{{.Object}}Reply{}, nil
}

// Update{{.Object}}Status 更新{{.Title}}状态
func (s *{{.ModuleUpper}}Service) Update{{.Object}}Status(c context.Context, req *pb.Update{{.Object}}StatusRequest) (*pb.Update{{.Object}}StatusReply, error) {
	return &pb.Update{{.Object}}StatusReply{}, s.uc.Update{{.Object}}Status(kratosx.MustContext(c), req.Id, req.Status)
}


// Delete{{.Object}} 删除{{.Title}}
func (s *{{.ModuleUpper}}Service) Delete{{.Object}}(c context.Context, req *pb.Delete{{.Object}}Request) (*pb.Delete{{.Object}}Reply, error) {
	total, err := s.uc.Delete{{.Object}}(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.Delete{{.Object}}Reply{Total: total}, nil
}

// GetTrash{{.Object}} 获取回收站指定的{{.Title}}
func (s *{{.ModuleUpper}}Service) GetTrash{{.Object}}(c context.Context, req *pb.GetTrash{{.Object}}Request) (*pb.GetTrash{{.Object}}Reply, error) {
	var ctx = kratosx.MustContext(c)

	result, err := s.uc.GetTrash{{.Object}}(ctx, req.Id)
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
func (s *{{.ModuleUpper}}Service) ListTrash{{.Object}}(c context.Context, req *pb.ListTrash{{.Object}}Request) (*pb.ListTrash{{.Object}}Reply, error) {
	var (
		in  = {{.Module}}.ListTrash{{.Object}}Request{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListTrash{{.Object}}(ctx, &in)
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

// DeleteTrash{{.Object}} 删除回收站{{.Title}}
func (s *{{.ModuleUpper}}Service) DeleteTrash{{.Object}}(ctx context.Context, req *pb.DeleteTrash{{.Object}}Request) (*pb.DeleteTrash{{.Object}}Reply, error) {
	total, err := s.uc.DeleteTrash{{.Object}}(kratosx.MustContext(ctx), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteTrash{{.Object}}Reply{Total: total}, nil
}

// RevertTrash{{.Object}} 还原回收站{{.Title}}
func (s *{{.ModuleUpper}}Service) RevertTrash{{.Object}}(ctx context.Context, req *pb.RevertTrash{{.Object}}Request) (*pb.RevertTrash{{.Object}}Reply, error) {
	return &pb.RevertTrash{{.Object}}Reply{},s.uc.RevertTrash{{.Object}}(kratosx.MustContext(ctx), req.Id)
}
