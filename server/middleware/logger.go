package middleware

import (
	"log/slog"
	"net/http"

	"github.com/rkunihiro/go-middleware-http-server/server/context"
)

func NewLoggerMiddleware(l *slog.Logger) Middleware {
	return func(w http.ResponseWriter, r *http.Request, next http.Handler) {
		ctx := context.WithLogger(r.Context(), l)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
