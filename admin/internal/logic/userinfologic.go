package logic

import (
	"context"

	"gozero_example/admin/internal/svc"
	"gozero_example/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line

	// 获取 jwt 载体信息
	value := l.ctx.Value("payload")
	resp = &types.UserInfoResp{}
	resp.Name = value.(string)

	return
}
