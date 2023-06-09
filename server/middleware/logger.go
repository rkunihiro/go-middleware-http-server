package middleware

import (
	"net/http"
	"time"

	"golang.org/x/exp/slog"
)

func NewLoggerMiddleware(l *slog.Logger) Middleware {
	return func(w http.ResponseWriter, r *http.Request, next http.Handler) {
		s := time.Now()
		next.ServeHTTP(w, r)
		l.With("method", r.Method, "path", r.URL.Path, "ms", time.Now().Sub(s).Milliseconds()).Info("request")
	}
}
