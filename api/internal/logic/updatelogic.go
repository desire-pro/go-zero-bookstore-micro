package logic

import (
	"context"

	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"
	"bookstore/rpc/change/changer"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateReq) (*types.UpdateResp, error) {
	resp, err := l.svcCtx.Changer.Update(l.ctx, &changer.UpdateReq{
		Book:  req.Book,
		Price: req.Price,
	})
	if err != nil {
		return &types.UpdateResp{
			Ok: false,
		}, nil
	}

	return &types.UpdateResp{
		Ok: resp.Ok,
	}, nil
}
