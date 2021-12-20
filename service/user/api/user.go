package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-mall/common/errorx"
	"net/http"

	"go-zero-mall/service/user/api/internal/config"
	"go-zero-mall/service/user/api/internal/handler"
	"go-zero-mall/service/user/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		fmt.Println(err)
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusBadRequest, &errorx.CodeErrorResponse{
				Code: 400,
				Message: e.Error(),
			}
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
