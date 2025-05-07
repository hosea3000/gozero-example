package image

import (
	"gozero_example/server/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozero_example/server/internal/logic/image"
	"gozero_example/server/internal/svc"
	"gozero_example/server/internal/types"
)

func FlagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageFlagReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := image.NewFlagLogic(r.Context(), svcCtx)
		resp, err := l.Flag(&req)
		response.Response(w, resp, err)
	}
}
