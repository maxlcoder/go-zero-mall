package logic

import (
	"context"
	"fmt"

	"go-zero-mall/service/user/api/internal/svc"
	"go-zero-mall/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginRsp, err error) {
	fmt.Println(req.Email)
	fmt.Println(req.Phone)
	fmt.Println(req.Password)
	return
}
