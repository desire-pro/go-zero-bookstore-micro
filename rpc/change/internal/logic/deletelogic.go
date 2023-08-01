package logic

import (
	"context"

	"bookstore/rpc/change/internal/svc"
	__ "bookstore/rpc/change/pb"
	"bookstore/rpc/check/checker"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *__.DeleteReq) (*__.DeleteResp, error) {
	// 在 deleteLogic 調用 check 服務
	resp, err := l.svcCtx.Checker.Check(l.ctx, &checker.CheckReq{
		Book: in.Book,
	})

	if resp == nil {
		err = l.svcCtx.Model.Delete(in.Book)
		if err != nil {
			return nil, err
		}
		return &__.DeleteResp{
			Ok: true,
		}, nil
	} else if err != nil {
		return &__.DeleteResp{
			Ok: false,
		}, nil
	}
	if resp.Found {
		return &__.DeleteResp{
			Ok: false,
		}, nil
	}
	return &__.DeleteResp{
		Ok: false,
	}, nil
}
