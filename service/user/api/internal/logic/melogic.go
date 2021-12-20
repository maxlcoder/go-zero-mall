package logic

import (
	"context"
	"go-zero-mall/common/errorx"
	"go-zero-mall/service/user/model"

	"go-zero-mall/service/user/api/internal/svc"
	"go-zero-mall/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) MeLogic {
	return MeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeLogic) Me(req types.MeReq) (resp *types.MeRsp, err error) {
	logx.Infof("user_id: %v", l.ctx.Value("user_id"))
	userId := l.ctx.Value("user_id")
	var user model.User
	result := l.svcCtx.DbEngin.Where("id", userId).First(&user)
	if result.Error != nil {
		return nil, errorx.NewDefaultError(result.Error.Error())
	}
	return &types.MeRsp{
		Id: int64(user.ID),
		Email: user.Email,
	}, nil
}
