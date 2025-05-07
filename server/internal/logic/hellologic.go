package logic

import (
	"context"
	"fmt"
	"gozero_example/server/internal/middleware"
	"gozero_example/server/internal/pkg/third/http_any"
	"gozero_example/server/internal/svc"
	"gozero_example/server/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

const ServiceName = "http_any"

type HelloLogic struct {
	logx.Logger
	ctx         context.Context
	svcCtx      *svc.ServiceContext
	serviceName string
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		Logger:      logx.WithContext(ctx),
		ctx:         ctx,
		svcCtx:      svcCtx,
		serviceName: ServiceName,
	}
}

func (l *HelloLogic) Hello(req *types.HelloReq) (resp *types.HelloResp, err error) {
	apiKey := l.ctx.Value("ApiKey")
	fmt.Println(apiKey)
	res, err := http_any.Hello()
	if err != nil {
		return nil, err
	}
	resp = &types.HelloResp{
		Res: res,
	}

	// 将值写入 Header（需修改 Logic 层返回的 Writer）
	// 注意：需在 HTTP 路由层传递 http.ResponseWriter
	w := l.ctx.Value("writer").(http.ResponseWriter)
	w.Header().Set(string(middleware.ServiceName), l.serviceName)

	return
}
