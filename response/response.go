package response

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error)  {
	var body Body
	if err != nil {
		body.Code = 400
		body.Message = err.Error()
	} else {
		body.Code = 200
		body.Message = "success"
		body.Data = resp
	}

	httpx.OkJson(w, body)
}
