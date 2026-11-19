package middleware

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func Recovery(api huma.API, log *slog.Logger) func(ctx huma.Context, next func(ctx huma.Context)) {
	return func(ctx huma.Context, next func(ctx huma.Context)) {
		paniced := false
		defer func() {
			rec := recover()
			if rec == nil {
				return
			}
			paniced = true

			err, ok := rec.(error)
			if !ok {
				err = fmt.Errorf("%v", rec)
			}
			log.Debug("panic",
				"error", err,
				"method", ctx.Method(),
				"path", ctx.URL().Path,
			)
			_ = huma.WriteErr(api, ctx, http.StatusInternalServerError, "", err)
		}()

		next(ctx)

		if paniced {
			return
		}
		if status := ctx.Status(); status >= http.StatusInternalServerError {
			log.Error("server_error",
				"method", ctx.Method(),
				"path", ctx.URL().Path,
				"status", status,
			)
		}
	}
}
