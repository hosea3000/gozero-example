package image

import (
	"context"

	"gozero_example/server/internal/svc"
	"gozero_example/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlagLogic {
	return &FlagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlagLogic) Flag(req *types.ImageFlagReq) (resp []types.ImageFlagResp, err error) {
	// todo: add your logic here and delete this line
	resp = []types.ImageFlagResp{
		{
			Name:       "dog",
			Confidence: 0.88,
		},
		{
			Name:       "cat",
			Confidence: 0.22,
		},
	}
	return
}
