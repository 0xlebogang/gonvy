package di

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"gonvy.internal/internal/config"
	"gonvy.internal/internal/httpapi"
	"gonvy.internal/internal/persistance/postgres"
)

type App interface {
	Run() error
	Close()
}
type app struct {
	router http.Handler
	db     postgres.DB
	cfg    *config.Config
}

func NewApp(cfg *config.Config, log *slog.Logger) (App, error) {
	if cfg == nil {
		return nil, fmt.Errorf("app: missing config")
	}
	if log == nil {
		return nil, fmt.Errorf("app: missing logger")
	}
	pgMod := postgres.New(cfg.DatabaseURL)
	db, err := pgMod.Connect()
	if err != nil {
		return nil, fmt.Errorf("app: failed to connect to database: %w", err)
	}
	r := httpapi.NewRouter(log, db)
	return &app{
		router: r,
		db:     db,
		cfg:    cfg,
	}, nil
}

func (a *app) Run() error {
	svr := http.Server{
		Addr:        ":" + a.cfg.Port,
		Handler:     a.router,
		ReadTimeout: 5 * time.Second,
	}
	return svr.ListenAndServe()
}

func (a *app) Close() {
	a.db.Close()
}
