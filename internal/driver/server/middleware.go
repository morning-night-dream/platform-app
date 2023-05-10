package server

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/morning-night-dream/platform-app/pkg/log"
	"go.uber.org/zap"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		ctx := log.SetLogCtx(r.Context())

		logger := log.GetLogCtx(ctx)

		lrw := NewLoggingResponseWriter(w)

		next.ServeHTTP(lrw, r)

		logger.Info(
			"access-log",
			zap.String("method", r.Method),
			zap.String("path", r.RequestURI),
			zap.String("addr", r.RemoteAddr),
			zap.String("host", r.Host),
			zap.String("proto", r.Proto),
			zap.String("scheme", r.URL.Scheme),
			zap.String("user-agent", r.Header["User-Agent"][0]),
			zap.String("code", strconv.Itoa(lrw.statusCode)),
			zap.String("elapsed", time.Since(now).String()),
			zap.Int64("elapsed(ns)", time.Since(now).Nanoseconds()),
		)

		if IsClientError(lrw.statusCode) {
			body, _ := io.ReadAll(r.Body)
			defer r.Body.Close()
			logger.Warn(
				"request-body",
				zap.String("body", string(body)),
			)
		}
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func IsClientError(code int) bool {
	return 400 <= code && code < 500
}
