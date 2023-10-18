package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go-middleware-http-server/server/middleware"
)

type Server struct {
	mux *http.ServeMux
	srv *http.Server
}

// Use adds a middleware to the server.
func (p *Server) Use(m middleware.Middleware) {
	h := p.srv.Handler
	p.srv.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m(w, r, h)
	})
}

// Handle registers a handler for the given pattern.
func (p *Server) Handle(pattern string, handler http.Handler) {
	p.mux.Handle(pattern, handler)
}

// ListenAndServe listens on the TCP network address and then
func (p *Server) ListenAndServe() error {
	if err := p.srv.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
		return err
	}
	return nil
}

func (p *Server) Shutdown(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return p.srv.Shutdown(ctx)
}

func New(port int) *Server {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	return &Server{mux: mux, srv: srv}
}
