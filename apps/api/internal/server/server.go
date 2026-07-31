package server

import (
	"fmt"
	"net/http"

	"github.com/0xlebogang/gonvy/api/internal/config"
	"github.com/0xlebogang/gonvy/api/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server interface {
	Start() error
}

type server struct {
	conf       *config.EnvConfig
	db         *gorm.DB
	middleware middleware.Middleware
	router     *gin.Engine
}

func New(c *config.EnvConfig, db *gorm.DB, m middleware.Middleware) Server {
	return &server{
		conf:       c,
		db:         db,
		middleware: m,
		router:     gin.Default(),
	}
}

func (s *server) createHttpServer() *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", s.conf.Port),
		Handler: s.router.Handler(),
	}
}

func (s *server) Start() error {
	api := s.router.Group("/api")
	api.Use(s.middleware.ErrorHandling())

	s.SetupV1Routes(api)
	svr := s.createHttpServer()
	return svr.ListenAndServe()
}
