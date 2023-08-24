package context

import (
	context2 "context"
	"log/slog"
)

type loggerKey struct{}

func WithLogger(ctx context2.Context, logger *slog.Logger) context2.Context {
	return context2.WithValue(ctx, loggerKey{}, logger)
}

func GetLogger(ctx context2.Context) *slog.Logger {
	return ctx.Value(loggerKey{}).(*slog.Logger)
}
