package httpapi

import (
	"log/slog"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
	"gonvy.internal/internal/persistance/postgres"
)

const (
	appName = "Gonv"
	version = "0.0.0"
)

func NewRouter(log *slog.Logger, db postgres.DB) (http.Handler, huma.API) {
	if log == nil {
		panic("logger is nil")
	}
	if db == nil {
		panic("db is nil")
	}
	setupErrors()
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	api := humagin.New(engine, huma.DefaultConfig(appName, version))
	registerHealthz(api)
	registerReadyz(api, db)
	return engine, api
}
