package http_error

import (
	"fmt"
	"net/http"
)

// 基础错误结构
type APIError struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
	Details    string `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("code: %d, message: %s, details: %s", e.Code, e.Message, e.Details)
}

// 通用错误 (10xxxx)
var (
	ErrMissingParameter = &APIError{
		Code:       100001,
		Message:    "请求参数缺失",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrInvalidParameterFormat = &APIError{
		Code:       100002,
		Message:    "参数格式错误",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrRateLimitExceeded = &APIError{
		Code:       100003,
		Message:    "请求频率超限",
		HTTPStatus: http.StatusTooManyRequests,
	}
)

// 鉴权错误 (40xxxx)
var (
	ErrInvalidAPIKey = &APIError{
		Code:       401001,
		Message:    "API Key 无效或过期",
		HTTPStatus: http.StatusUnauthorized,
	}

	ErrSignatureMismatch = &APIError{
		Code:       401002,
		Message:    "签名验证失败",
		HTTPStatus: http.StatusForbidden,
	}

	ErrTimestampExpired = &APIError{
		Code:       401003,
		Message:    "时间戳过期",
		HTTPStatus: http.StatusForbidden,
	}

	ErrInvalidTimestamp = &APIError{
		Code:       401004,
		Message:    "时间戳格式错误",
		HTTPStatus: http.StatusForbidden,
	}
)

// 业务逻辑错误 (42xxxx)
var (
	ErrImageProcessingFailed = &APIError{
		Code:       402001,
		Message:    "图片解析失败",
		HTTPStatus: http.StatusUnprocessableEntity,
	}

	ErrInvalidConfidenceThreshold = &APIError{
		Code:       402002,
		Message:    "置信度阈值不合法",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrUnsupportedLanguage = &APIError{
		Code:       402003,
		Message:    "文本语言不支持",
		HTTPStatus: http.StatusBadRequest,
	}
)

// 第三方服务错误 (50xxxx)
var (
	ErrAWSServiceFailure = &APIError{
		Code:       503001,
		Message:    "AWS服务调用失败",
		HTTPStatus: http.StatusBadGateway,
	}

	ErrAliyunRateLimit = &APIError{
		Code:       503002,
		Message:    "阿里云服务限流",
		HTTPStatus: http.StatusTooManyRequests,
	}

	ErrGoogleServiceTimeout = &APIError{
		Code:       503003,
		Message:    "Google服务响应超时",
		HTTPStatus: http.StatusGatewayTimeout,
	}
)

// 内部服务错误 (50xxxx)
var (
	ErrInternalServiceError = &APIError{
		Code:       500000,
		Message:    "内部服务错误",
		HTTPStatus: http.StatusBadGateway,
	}

	ErrDatabaseConnectionFailed = &APIError{
		Code:       500001,
		Message:    "数据库连接失败",
		HTTPStatus: http.StatusInternalServerError,
	}

	ErrRedisCacheError = &APIError{
		Code:       500002,
		Message:    "Redis缓存异常",
		HTTPStatus: http.StatusInternalServerError,
	}

	ErrAsyncTaskQueueOverflow = &APIError{
		Code:       500003,
		Message:    "异步任务队列堆积",
		HTTPStatus: http.StatusInternalServerError,
	}
)

// 包装错误细节的构造函数
func NewAPIError(baseError *APIError, details string) *APIError {
	return &APIError{
		Code:       baseError.Code,
		Message:    baseError.Message,
		HTTPStatus: baseError.HTTPStatus,
		Details:    details,
	}
}
