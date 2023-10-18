package logger

import (
	"log/slog"
	"os"
	"strings"
	"time"
)

func New() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				a.Value = slog.StringValue(time.Now().Format("2006-01-02T15:04:05.000Z07:00"))
				break
			case slog.LevelKey:
				a.Value = slog.StringValue(strings.ToLower(a.Value.String()))
				break
			case slog.MessageKey:
				a.Key = "message"
				break
			}
			return a
		},
	}))
}
