package svc

import (
	"go-zero-mall/service/user/rpc/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 数据库 gorm 处理
	db, err := gorm.Open(mysql.Open(c.DataSourceName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "", // 表前缀
			SingularTable: false, // 是否单数表名
		},
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		DbEngin: db,
	}
}
