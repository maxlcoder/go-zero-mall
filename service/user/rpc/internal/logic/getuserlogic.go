package logic

import (
	"context"
	"go-zero-mall/service/user/model"

	"go-zero-mall/service/user/rpc/internal/svc"
	"go-zero-mall/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdReq) (*user.UserRsp, error) {
	var usermodel model.User
	result := l.svcCtx.DbEngin.Where("id", in.Id).First(&usermodel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user.UserRsp{
		Id: int64(usermodel.ID),
		Name: usermodel.Name,
		Email: usermodel.Email,
	}, nil
}
