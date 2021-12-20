package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"go-zero-mall/service/user/rpc/userclient"

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
	logx.Infof(l.svcCtx.Config.LogConf.Mode)
	userId := l.ctx.Value("user_id")
	// api 直接数据库取值
	//var user model.User
	//result := l.svcCtx.DbEngin.Where("id", userId).First(&user)
	//if result.Error != nil {
	//	return nil, errorx.NewDefaultError(result.Error.Error())
	//}

	// 调用 rpc 服务
	userIdNumber := json.Number(fmt.Sprintf("%v", userId))
	userIdInt, err := userIdNumber.Int64()
	if err != nil {
		return nil, err
	}
	userRsp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.IdReq{
		Id: userIdInt,
	})
	if err != nil {
		return nil, err
	}

	return &types.MeRsp{
		Id: userRsp.Id,
		Email: userRsp.Email,
	}, nil
}
