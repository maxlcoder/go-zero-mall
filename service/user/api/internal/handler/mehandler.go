package handler

import (
	"go-zero-mall/response"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-mall/service/user/api/internal/logic"
	"go-zero-mall/service/user/api/internal/svc"
	"go-zero-mall/service/user/api/internal/types"
)

func meHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMeLogic(r.Context(), ctx)
		resp, err := l.Me(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
