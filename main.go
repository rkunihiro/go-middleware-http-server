package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/rkunihiro/go-middleware-http-server/server"
	"github.com/rkunihiro/go-middleware-http-server/server/middleware"
)

func main() {
	// slogの初期化
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(time.Now().Format("2006-01-02T15:04:05.000Z07:00"))
			}
			if a.Key == slog.LevelKey {
				a.Value = slog.StringValue(strings.ToLower(a.Value.String()))
			}
			if a.Key == slog.MessageKey {
				a.Key = "message"
			}
			return a
		},
	}))

	port := 3000
	s := server.New(port)
	s.Use(middleware.NewLoggingMiddleware())
	s.Use(middleware.NewRequestIDMiddleware())
	s.Use(middleware.NewLoggerMiddleware(log))
	s.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("OK"))
	}))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		<-ctx.Done()
		log.Info("close by signal")
		if err := s.Shutdown(5 * time.Second); err != nil {
			log.With("err", err).Error("shutdown error")
		}
		log.Info("server closed")
	}()

	log.With("port", port).Info("start server")
	if err := s.ListenAndServe(); err != nil {
		log.With("err", err).Error("listen and serve error")
	}
}
