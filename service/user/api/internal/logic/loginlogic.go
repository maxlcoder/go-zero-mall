package logic

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-mall/common/errorx"
	"go-zero-mall/service/user/api/internal/svc"
	"go-zero-mall/service/user/api/internal/types"
	"go-zero-mall/service/user/model"
	"golang.org/x/crypto/bcrypt"
	"time"
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
	email := req.Email
	// 查询是否存在用户
	var user model.User
	result := l.svcCtx.DbEngin.Where("email", email).First(&user)
	if result.Error != nil {
		return nil, errorx.NewDefaultError("当前 email 未注册")
	}

	fmt.Println(user)

	// 检查密码是否匹配
	password := req.Password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errorx.NewDefaultError("密码错误")
	}

	// jwt
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, int64(user.ID))
	if err != nil {
		return nil, errorx.NewDefaultError("登录失败")
	}
	fmt.Println(user.Email)
	return &types.LoginRsp{
		Id:    int64(user.ID),
		Name:  user.Name,
		Email: user.Email,
		AccessToken: jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["user_id"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}