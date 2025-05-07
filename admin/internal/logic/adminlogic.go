package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"

	"gozero_example/admin/internal/svc"
	"gozero_example/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLogic {
	return &AdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLogic) Admin(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	resp.Message = req.Name
	logc.Info(context.Background(), "hello world")
	return
}
