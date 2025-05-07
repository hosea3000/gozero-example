package handler

import (
	"gozero_example/server/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozero_example/server/internal/logic"
	"gozero_example/server/internal/svc"
	"gozero_example/server/internal/types"
)

func helloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewHelloLogic(r.Context(), svcCtx)
		resp, err := l.Hello(&req)

		response.Response(w, resp, err)
	}
}
