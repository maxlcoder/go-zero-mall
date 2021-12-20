package svc

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-mall/service/user/api/internal/config"
	"go-zero-mall/service/user/api/internal/middleware"
	"go-zero-mall/service/user/model"
	"go-zero-mall/service/user/rpc/userclient"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config config.Config
	DbEngin *gorm.DB
	Example rest.Middleware
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 初始化日志配置
	logx.MustSetup(c.LogConf)

	db, err := gorm.Open(mysql.Open(c.DataSourceName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "", // 表前缀
			SingularTable: false, // 是否单数表名
		},
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})

	return &ServiceContext{
		Config: c,
		DbEngin: db,
		Example: middleware.NewExampleMiddleware().Handle,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
