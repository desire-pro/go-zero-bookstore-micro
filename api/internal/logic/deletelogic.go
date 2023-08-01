package logic

import (
	"context"

	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"
	"bookstore/rpc/change/changer"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DeleteReq) (*types.DeleteResp, error) {
	resp, err := l.svcCtx.Changer.Delete(l.ctx, &changer.DeleteReq{
		Book: req.Book,
	})

	if err != nil {
		return &types.DeleteResp{
			Ok: false,
		}, nil
	}

	return &types.DeleteResp{
		Ok: resp.Ok,
	}, nil
}
