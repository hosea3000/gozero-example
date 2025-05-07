package credential

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gozero_example/model"
	"gozero_example/server/internal/config"
	"gozero_example/server/response/http_error"
	"time"
)

type CredentialLogic struct {
	logx.Logger
	ctx    context.Context
	DB     *gorm.DB
	Config config.Config
}

func NewCredentialLogic(ctx context.Context, DB *gorm.DB, C config.Config) *CredentialLogic {
	return &CredentialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		DB:     DB,
		Config: C,
	}
}

type SignaturePayload struct {
	ApiKey    string
	Path      string
	Timestamp string
}

func (l *CredentialLogic) CheckCredential(signaturePayload SignaturePayload, signature string) error {
	// 基础校验
	if signaturePayload.ApiKey == "" || signaturePayload.Timestamp == "" || signature == "" {
		return http_error.ErrInvalidAPIKey
	}

	credential := &model.Credential{}
	err := l.DB.WithContext(l.ctx).First(&credential, "api_key = ? and is_active = true", signaturePayload.ApiKey).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http_error.ErrInvalidAPIKey
		}
		return err
	}

	if credential.ExpiresAt != nil && time.Now().After(*credential.ExpiresAt) {
		// 不设置过期时间，则永不过期
		return http_error.ErrTimestampExpired
	}

	// 验证时间戳有效性
	reqTime, err := time.Parse(time.RFC1123, signaturePayload.Timestamp)
	if err != nil {
		return http_error.ErrInvalidTimestamp
	}
	if IsTimeWithin(time.Now(), reqTime, time.Duration(l.Config.Auth.SignatureExpireMinutes)*time.Minute) == false {
		return http_error.ErrTimestampExpired
	}

	// 重新计算签名
	expectedSignature := CalSignature(signaturePayload, credential.Secret)
	// 对比签名
	if !hmac.Equal([]byte(signature), []byte(expectedSignature)) {
		logx.WithContext(l.ctx).Errorf("signature mismatch, expected: %s, got: %s",
			expectedSignature, signature)
		return http_error.ErrSignatureMismatch
	}

	// 记录最后使用时间
	now := time.Now()
	err = l.DB.WithContext(l.ctx).Model(&credential).Updates(&model.Credential{LastUsedAt: &now}).Error
	if err != nil {
		return err
	}
	return nil
}

func CalSignature(signaturePayload SignaturePayload, secret string) string {
	payload := fmt.Sprintf("%s\n%s\n%s", signaturePayload.ApiKey, signaturePayload.Timestamp, signaturePayload.Path)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	expectedSignature := hex.EncodeToString(mac.Sum(nil))
	return expectedSignature
}

// IsTimeWithin 判断目标时间是否在基准时间的前后 5 分钟内
func IsTimeWithin(target, base time.Time, duration time.Duration) bool {
	diff := target.Sub(base)
	ok := diff.Abs() <= duration
	return ok
}
