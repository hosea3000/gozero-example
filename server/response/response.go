package response

import (
	"gozero_example/server/response/http_error"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		body.Code = http_error.ErrInternalServiceError.Code
		body.Msg = http_error.ErrInternalServiceError.Message
		body.Details = err.Error()
	} else {
		body.Msg = "OK"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}

func ErrorFormat(baseError *http_error.APIError) any {
	body := Body{
		Code: baseError.Code,
		Msg:  baseError.Message,
		Data: nil,
	}
	return body
}
