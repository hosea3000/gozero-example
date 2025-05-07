package logic

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"gozero_example/model"
	"strconv"
	"time"

	"gozero_example/admin/internal/svc"
	"gozero_example/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	user := &model.User{}
	err = l.svcCtx.DB.Where("username = ? and password = ?", req.Username, req.Password).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

		}
		return nil, err
	}

	token, err := getJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.Auth.AccessExpire, "1")
	if err != nil {
		return nil, err
	}

	resp = &types.LoginResp{
		ID:    strconv.Itoa(int(user.ID)),
		Name:  user.Username,
		Token: token,
	}
	return
}

// @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @payload: 数据载体
func getJwtToken(secretKey string, iat, seconds int64, payload string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
