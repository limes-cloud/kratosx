package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valutil"

	"github.com/limes-cloud/kratosx/cmd/kratosx/api/kratosx/errors"
	pb "github.com/limes-cloud/kratosx/cmd/kratosx/api/kratosx/notice/v1"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/biz/notice"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/conf"
	"github.com/limes-cloud/kratosx/cmd/kratosx/internal/data"
)

type NoticeService struct {
	pb.UnimplementedNoticeServer
	uc *notice.UseCase
}

func NewNotice(conf *conf.Config) *NoticeService {
	return &NoticeService{
		uc: notice.NewUseCase(conf, data.NewNoticeRepo()),
	}
}

// GetNotice 获取指定的通知
func (s *NoticeService) GetNotice(c context.Context, req *pb.GetNoticeRequest) (*pb.GetNoticeReply, error) {
	var ctx = kratosx.MustContext(c)

	result, err := s.uc.GetNotice(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	reply := pb.GetNoticeReply{}
	if err := valutil.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.Transform()
	}
	return &reply, nil
}

// ListNotice 获取通知列表
func (s *NoticeService) ListNotice(c context.Context, req *pb.ListNoticeRequest) (*pb.ListNoticeReply, error) {
	var (
		in  = notice.ListNoticeRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.Transform()
	}
	result, total, err := s.uc.ListNotice(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListNoticeReply{Total: total}
	if err := value.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.Transform()
	}
	return &reply, nil
}

// CreateNotice 创建通知
func (s *NoticeService) CreateNotice(c context.Context, req *pb.CreateNoticeRequest) (*pb.CreateNoticeReply, error) {
	var (
		in  = notice.Notice{}
		ctx = kratosx.MustContext(c)
	)

	if err := valuer.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.Transform()
	}
	id, err := s.uc.CreateNotice(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateNoticeReply{Id: id}, nil
}

// ImportNotice 导入通知
func (s *NoticeService) ImportNotice(c context.Context, req *pb.ImportNoticeRequest) (*pb.ImportNoticeReply, error) {
	var (
		in  []*notice.Notice
		ctx = kratosx.MustContext(c)
	)

	if err := value.Transform(req.List, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.Transform()
	}
	ct, ut, err := s.uc.ImportNotice(ctx, in)
	if err != nil {
		return nil, err
	}

	return &pb.ImportNoticeReply{CreateTotal: ct, UpdateTotal: ut}, nil
}

// ExportNotice 导出通知
func (s *NoticeService) ExportNotice(c context.Context, req *pb.ExportNoticeRequest) (*pb.ExportNoticeReply, error) {
	var (
		in  = notice.ExportNoticeRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.Transform()
	}
	src, err := s.uc.ExportNotice(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.ExportNoticeReply{Src: src}, nil
}

// UpdateNotice 更新通知
func (s *NoticeService) UpdateNotice(c context.Context, req *pb.UpdateNoticeRequest) (*pb.UpdateNoticeReply, error) {
	var (
		in  = notice.Notice{}
		ctx = kratosx.MustContext(c)
	)

	if err := value.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.Transform()
	}
	if err := s.uc.UpdateNotice(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateNoticeReply{}, nil
}

// DeleteNotice 删除通知
func (s *NoticeService) DeleteNotice(ctx context.Context, req *pb.DeleteNoticeRequest) (*pb.DeleteNoticeReply, error) {
	return &pb.DeleteNoticeReply{}, s.uc.DeleteNotice(kratosx.MustContext(ctx), req.Id)
}

// BatchDeleteNotice 批量删除通知
func (s *NoticeService) BatchDeleteNotice(ctx context.Context, req *pb.BatchDeleteNoticeRequest) (*pb.BatchDeleteNoticeReply, error) {
	total, err := s.uc.BatchDeleteNotice(kratosx.MustContext(ctx), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.BatchDeleteNoticeReply{Total: total}, nil
}
