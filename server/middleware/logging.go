package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"go-middleware-http-server/server/context"
)

func NewLoggingMiddleware() Middleware {
	return func(w http.ResponseWriter, r *http.Request, next http.Handler) {
		s := time.Now()
		l := context.GetLogger(r.Context())
		l.With(
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
		).Info("request")

		next.ServeHTTP(w, r)

		l.With(
			slog.Int64("ms", time.Now().Sub(s).Milliseconds()),
		).Info("response")
	}
}
