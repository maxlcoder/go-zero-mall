// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zero-mall/service/user/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Example},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/register",
					Handler: registerHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/login",
					Handler: loginHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/me",
				Handler: meHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
