package middleware

import (
	"net/http"
	"time"

	"github.com/rkunihiro/go-middleware-http-server/server/context"
)

func NewLoggingMiddleware() Middleware {
	return func(w http.ResponseWriter, r *http.Request, next http.Handler) {
		s := time.Now()
		l := context.GetLogger(r.Context())
		l.With("method", r.Method, "path", r.URL.Path).Info("request")
		next.ServeHTTP(w, r)
		l.With("ms", time.Now().Sub(s).Milliseconds()).Info("response")
	}
}
