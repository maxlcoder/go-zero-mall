package logic

import (
	"context"
	"go-zero-mall/common/errorx"
	"go-zero-mall/service/user/model"
	"golang.org/x/crypto/bcrypt"

	"go-zero-mall/service/user/api/internal/svc"
	"go-zero-mall/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (resp *types.RegisterRsp, err error) {
	email := req.Email
	phone := req.Phone

	// 检查是否已经被注册了
	user := model.User{}
	result := l.svcCtx.DbEngin.Model(&user).Where("email", email).Or("phone", phone).First(&user)
	if result.RowsAffected != 0 {
		return nil, errorx.NewDefaultError("邮箱或手机号已经被使用")
	}

	// 注册用户
	password := req.Password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user = model.User{
		Email: email,
		Password: string(hashedPassword),
	}
	result = l.svcCtx.DbEngin.Model(&user).Create(&user)
	if result.Error != nil {
		return nil, errorx.NewDefaultError("注册失败")
	}

	return &types.RegisterRsp{
		Id: int64(user.ID),
		Email: user.Email,
	}, nil


	return
}
