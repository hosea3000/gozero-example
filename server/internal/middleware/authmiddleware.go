package middleware

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm.io/gorm"
	"gozero_example/server/internal/config"
	"gozero_example/server/internal/logic/credential"
	"net/http"
)

type AuthMiddleware struct {
	DB     *gorm.DB
	Config config.Config
}

func NewAuthMiddleware(DB *gorm.DB, c config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		DB:     DB,
		Config: c,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取 Header 参数
		apiKey := r.Header.Get("X-API-Key")
		timestamp := r.Header.Get("X-API-Timestamp")
		signature := r.Header.Get("X-API-Signature")

		// 校验签名
		l := credential.NewCredentialLogic(r.Context(), m.DB, m.Config)
		payload := credential.SignaturePayload{
			ApiKey:    apiKey,
			Timestamp: timestamp,
			Path:      r.URL.Path,
		}
		err := l.CheckCredential(payload, signature)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), AppKey, apiKey)
		newReq := r.WithContext(ctx)

		next(w, newReq)
	}
}
