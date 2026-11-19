package middleware

import (
	"log/slog"
	"net/http"
	"time"

	huma "github.com/danielgtaylor/huma/v2"
)

func RequestLogger(log *slog.Logger) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		start := time.Now()
		next(ctx)
		path := ctx.URL().Path
		if path == "/healthz" || path == "/readyz" {
			return
		}
		status := ctx.Status()
		attrs := []any{
			"method", ctx.Method(),
			"path", path,
			"status", status,
			"latency", time.Since(start),
			"ip_address", ctx.RemoteAddr(),
		}
		switch {
		case status >= http.StatusInternalServerError:
			log.Error("request", attrs...)
		case status >= http.StatusBadRequest:
			log.Error("request", attrs...)
		default:
			log.Info("request", attrs...)
		}
	}
}
