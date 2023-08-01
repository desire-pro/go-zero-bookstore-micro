package logic

import (
	"context"

	"bookstore/rpc/change/internal/svc"
	__ "bookstore/rpc/change/pb"
	"bookstore/rpc/check/checker"
	"bookstore/rpc/model/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *__.UpdateReq) (*__.UpdateResp, error) {
	// 4. 在 updateLogic 調用 check 服務
	resp, err := l.svcCtx.Checker.Check(l.ctx, &checker.CheckReq{
		Book: in.Book,
	})

	if resp.Found {
		err = l.svcCtx.Model.Update(model.Book{
			Book:  in.Book,
			Price: in.Price,
		})

		if err != nil {
			return nil, err
		}
		return &__.UpdateResp{
			Ok: true,
		}, nil
	} else if err != nil {
		return &__.UpdateResp{
			Ok: false,
		}, nil
	} else {
		return &__.UpdateResp{
			Ok: false,
		}, nil
	}
}
