package server

import (
	"fmt"
	"net/http"

	"github.com/0xlebogang/gonvy/api/internal/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server interface {
	Start() error
}

type server struct {
	conf   *config.EnvConfig
	db     *gorm.DB
	router *gin.Engine
}

func New(c *config.EnvConfig, db *gorm.DB) Server {
	return &server{
		conf:   c,
		db:     db,
		router: gin.Default(),
	}
}

func (s *server) createHttpServer() *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", s.conf.Port),
		Handler: s.router.Handler(),
	}
}

func (s *server) Start() error {
	s.SetupRoutes()
	svr := s.createHttpServer()
	return svr.ListenAndServe()
}
