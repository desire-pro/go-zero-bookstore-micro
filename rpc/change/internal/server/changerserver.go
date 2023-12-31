// Code generated by goctl. DO NOT EDIT.
// Source: change.proto

package server

import (
	"context"

	"bookstore/rpc/change/internal/logic"
	"bookstore/rpc/change/internal/svc"
	"bookstore/rpc/change/pb"
)

type ChangerServer struct {
	svcCtx *svc.ServiceContext
	__.UnimplementedChangerServer
}

func NewChangerServer(svcCtx *svc.ServiceContext) *ChangerServer {
	return &ChangerServer{
		svcCtx: svcCtx,
	}
}

func (s *ChangerServer) Update(ctx context.Context, in *__.UpdateReq) (*__.UpdateResp, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *ChangerServer) Delete(ctx context.Context, in *__.DeleteReq) (*__.DeleteResp, error) {
	l := logic.NewDeleteLogic(ctx, s.svcCtx)
	return l.Delete(in)
}
