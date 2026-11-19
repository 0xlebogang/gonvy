package main

import (
	"log/slog"
	"os"

	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/gin-gonic/gin"
	"gonvy.internal/internal/config"
	"gonvy.internal/internal/di"
)

type Options struct {
	Debug bool `doc:"Enable debug logging"`
	Port  int  `default:"8080" doc:"Port to listen on" short:"p"`
}

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		cfg := config.Load()
		log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: map[bool]slog.Level{true: slog.LevelInfo, false: slog.LevelDebug}[gin.Mode() == gin.ReleaseMode],
		}))

		var app di.App
		var err error

		hooks.OnStart(func() {
			app, err = di.NewApp(cfg, log)
			if err != nil {
				log.Error("application initialization failed", "error", err)
				app.Close()
				os.Exit(1)
			}
			if err := app.Run(); err != nil {
				log.Error("application startup failed", "error", err)
				app.Close()
				os.Exit(1)
			}
		})

		hooks.OnStop(func() {
			app.Close()
		})
	})
	cli.Run()
}
