package middleware

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/rkunihiro/go-middleware-http-server/server/context"
)

func NewRequestIDMiddleware() Middleware {
	return func(w http.ResponseWriter, r *http.Request, next http.Handler) {
		id := uuid.New().String()
		r.Header.Set("X-Request-ID", id)
		if logger := context.GetLogger(r.Context()); logger != nil {
			logger = logger.With("reqId", id)
			r = r.WithContext(context.WithLogger(r.Context(), logger))
		}
		next.ServeHTTP(w, r)
	}
}
