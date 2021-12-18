package logic

import (
	"context"

	"go-zero-mall/service/user/rpc/greet/greet"
	"go-zero-mall/service/user/rpc/greet/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *greet.Request) (*greet.Response, error) {
	// todo: add your logic here and delete this line

	return &greet.Response{}, nil
}
