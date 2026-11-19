package main

import (
	"log/slog"
	"os"

	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gonvy.internal/internal/config"
	"gonvy.internal/internal/di"
	"gonvy.internal/internal/persistance/postgres"
)

type Options struct {
	Debug bool `doc:"Enable debug logging"`
	Port  int  `default:"8080" doc:"Port to listen on" short:"p"`
}

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		// Server wiring stays in OnStart so subcommands (e.g. migrate) do not
		// construct the HTTP app — only the default command starts the server.
		var app di.App

		hooks.OnStart(func() {
			cfg := config.Load()
			log := newLogger()

			var err error
			app, err = di.NewApp(cfg, log)
			if err != nil {
				log.Error("application initialization failed", "error", err)
				os.Exit(1)
			}
			if err := app.Run(); err != nil {
				log.Error("application startup failed", "error", err)
				app.Close()
				os.Exit(1)
			}
		})

		hooks.OnStop(func() {
			if app != nil {
				app.Close()
			}
		})
	})

	cli.Root().AddCommand(&cobra.Command{
		Use:   "migrate",
		Short: "Apply pending database migrations",
		Run: humacli.WithOptions(func(cmd *cobra.Command, args []string, opts *Options) {
			cfg := config.Load()
			log := newLogger()

			db, err := postgres.New(cfg.DatabaseURL).Connect()
			if err != nil {
				log.Error("database connection failed", "error", err)
				os.Exit(1)
			}
			defer db.Close()

			if err := db.Migrate(); err != nil {
				log.Error("migrations failed", "error", err)
				os.Exit(1)
			}
			log.Info("migrations applied successfully")
		}),
	})

	cli.Run()
}

func newLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: map[bool]slog.Level{true: slog.LevelInfo, false: slog.LevelDebug}[gin.Mode() == gin.ReleaseMode],
	}))
}
