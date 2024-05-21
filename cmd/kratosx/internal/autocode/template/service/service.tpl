package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/value"

	"{{.Server}}/api/{{.ServerName}}/errors"
	pb "{{.Server}}/api/{{.ServerName}}/{{.Module}}/v1"
	"{{.Server}}/internal/biz/notice"
	"{{.Server}}/internal/conf"
	"{{.Server}}/internal/data"
)

type {{.ModuleUpper}}Service struct {
	pb.Unimplemented{{.ModuleUpper}}Server
	uc *notice.UseCase
}

func New{{.ModuleUpper}}(conf *conf.Config) *{{.ModuleUpper}}Service {
	return &{{.ModuleUpper}}Service{
		uc: notice.NewUseCase(conf, data.New{{.ModuleUpper}}Repo()),
	}
}

// Get{{.Object}} 获取指定的{{.Title}}
func (s *{{.ModuleUpper}}Service) Get{{.Object}}(c context.Context, req *pb.Get{{.Object}}Request) (*pb.Get{{.Object}}Reply, error) {
	var ctx = kratosx.MustContext(c)

	result, err := s.uc.Get{{.Object}}(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	reply := pb.Get{{.Object}}Reply{}
	if err := value.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.Transform()
	}
	return &reply, nil
}

// List{{.Object}} 获取{{.Title}}列表
func (s *{{.ModuleUpper}}Service) List{{.Object}}(c context.Context, req *pb.List{{.Object}}Request) (*pb.List{{.Object}}Reply, error) {
	var (
		in  = notice.List{{.Object}}Request{}
		ctx = kratosx.MustContext(c)
	)

	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.Transform()
	}
	result, total, err := s.uc.List{{.Object}}(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.List{{.Object}}Reply{Total: total}
	if err := value.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.Transform()
	}
	return &reply, nil
}

// Create{{.Object}} 创建{{.Title}}
func (s *{{.ModuleUpper}}Service) Create{{.Object}}(c context.Context, req *pb.Create{{.Object}}Request) (*pb.Create{{.Object}}Reply, error) {
	var (
		in  = notice.{{.Object}}{}
		ctx = kratosx.MustContext(c)
	)

	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.Transform()
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
		in  []*notice.{{.Object}}
		ctx = kratosx.MustContext(c)
	)

	if err := value.Transform(req.List, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.Transform()
	}
	ct, ut, err := s.uc.Import{{.Object}}(ctx, in)
	if err != nil {
		return nil, err
	}

	return &pb.Import{{.Object}}Reply{CreateTotal: ct, UpdateTotal: ut}, nil
}

// Export{{.Object}} 导出{{.Title}}
func (s *{{.ModuleUpper}}Service) Export{{.Object}}(c context.Context, req *pb.Export{{.Object}}Request) (*pb.Export{{.Object}}Reply, error) {
	var (
		in  = notice.Export{{.Object}}Request{}
		ctx = kratosx.MustContext(c)
	)

	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.Transform()
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
		in  = notice.{{.Object}}{}
		ctx = kratosx.MustContext(c)
	)

	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.Transform()
	}
	if err := s.uc.Update{{.Object}}(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.Update{{.Object}}Reply{}, nil
}

// Delete{{.Object}} 删除{{.Title}}
func (s *{{.ModuleUpper}}Service) Delete{{.Object}}(ctx context.Context, req *pb.Delete{{.Object}}Request) (*pb.Delete{{.Object}}Reply, error) {
	return &pb.Delete{{.Object}}Reply{}, s.uc.Delete{{.Object}}(kratosx.MustContext(ctx), req.Id)
}

// BatchDelete{{.Object}} 批量删除{{.Title}}
func (s *{{.ModuleUpper}}Service) BatchDelete{{.Object}}(ctx context.Context, req *pb.BatchDelete{{.Object}}Request) (*pb.BatchDelete{{.Object}}Reply, error) {
	total, err := s.uc.BatchDelete{{.Object}}(kratosx.MustContext(ctx), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.BatchDelete{{.Object}}Reply{Total: total}, nil
}
