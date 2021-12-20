package svc

import (
	"github.com/tal-tech/go-zero/rest"
	"go-zero-mall/service/user/api/internal/config"
	"go-zero-mall/service/user/api/internal/middleware"
	"go-zero-mall/service/user/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config config.Config
	DbEngin *gorm.DB
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

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
	}
}
