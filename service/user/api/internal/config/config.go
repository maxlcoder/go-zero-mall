package config

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	logx.LogConf
	DataSourceName string
}
