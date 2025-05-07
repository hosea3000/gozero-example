package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"
)

type CallLogMiddleware struct {
}

func NewCallLogMiddleware() *CallLogMiddleware {
	return &CallLogMiddleware{}
}

func (m *CallLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		// 记录开始时间
		startTime := time.Now()

		apiKey := r.Context().Value(AppKey).(string)
		// 包装 ResponseWriter 以获取状态码
		lrw := &loggingResponseWriter{w: w}

		// 处理请求
		defer func() {
			// 计算耗时
			latency := time.Since(startTime)
			// 从 Header 中获取值
			serviceName := w.Header().Get(ServiceName)
			// 记录日志
			logx.WithContext(r.Context()).Infof(
				"===[%s] %s - Status: %d - Latency: %v apiKey: %s， service: %s ==",
				r.Method,
				r.RequestURI,
				lrw.status,
				latency,
				apiKey,
				serviceName,
			)
		}()

		next(w, r)
	}
}

// 自定义 ResponseWriter 用于获取状态码
type loggingResponseWriter struct {
	w      http.ResponseWriter
	status int
}

func (l *loggingResponseWriter) Header() http.Header {
	return l.w.Header()
}

func (l *loggingResponseWriter) Write(b []byte) (int, error) {
	return l.w.Write(b)
}

func (l *loggingResponseWriter) WriteHeader(statusCode int) {
	l.status = statusCode
	l.w.WriteHeader(statusCode)
}
