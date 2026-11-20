package httpapi

import (
	"log/slog"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
	"gonvy.internal/internal/domain/users"
	"gonvy.internal/internal/httpapi/middleware"
	"gonvy.internal/internal/persistance/postgres"
)

const (
	appName = "Gonv"
	version = "0.0.0"
)

func NewRouter(log *slog.Logger, db postgres.DB) http.Handler {
	if log == nil {
		panic("logger is nil")
	}
	if db == nil {
		panic("db is nil")
	}
	setupErrors()
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	if gin.Mode() != gin.ReleaseMode {
		engine.Use(gin.Logger(), gin.Recovery())
	}
	api := humagin.New(engine, huma.DefaultConfig(appName, version))
	if gin.Mode() == gin.ReleaseMode {
		api.UseMiddleware(middleware.RequestLogger(log), middleware.Recovery(api, log))
	}
	registerHealthz(api)
	registerReadyz(api, db)
	users.BuildModule(db.Client()).
		Register(api)
	return engine
}
