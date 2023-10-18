package main

import (
	"context"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"go-middleware-http-server/server"
	"go-middleware-http-server/server/logger"
	"go-middleware-http-server/server/middleware"
)

func main() {
	// slogの初期化
	log := logger.New()
	slog.SetDefault(log)

	port := 3000
	s := server.New(port)
	s.Use(middleware.NewLoggingMiddleware())
	s.Use(middleware.NewRequestIDMiddleware())
	s.Use(middleware.NewLoggerMiddleware(log))

	s.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(200)
		_, _ = w.Write([]byte("OK"))
	}))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		<-ctx.Done()
		slog.Info("close by signal")
		if err := s.Shutdown(5 * time.Second); err != nil {
			slog.With("err", err).Error("shutdown error")
		}
		slog.Info("server closed")
	}()

	slog.With("port", port).Info("start server")
	if err := s.ListenAndServe(); err != nil {
		slog.With("err", err).Error("listen and serve error")
	}
}
