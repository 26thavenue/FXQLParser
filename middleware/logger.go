package middlewares

import (
	"log/slog"
	"net/http"
	"time"
)

type ResponseWriter struct {
	http.ResponseWriter
	status int 
}


func (r ResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.status = statusCode
}

func Logger(logger *slog.Logger, next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// ctx := r.Context()
        // ctx = context.WithValue(ctx, "startTime", start)
        // r = r.WithContext(ctx)

		writer := &ResponseWriter{
			ResponseWriter: w,
			status:http.StatusOK,
		}

		next.ServeHTTP(writer, r)

		logger.Info(
			"handled request",
			slog.Int("statusCode", writer.status),
			slog.String("remoteAddr", r.RemoteAddr),
			slog.String("xffHeader", r.Header.Get("X-Forwarded-For")),
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Any("duration", time.Since(start)),
		)
	})
}