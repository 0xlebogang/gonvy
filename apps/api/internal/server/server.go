package server

import (
	"fmt"
	"net/http"

	"github.com/0xlebogang/gonvy/api/internal/config"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Start() error
}

type server struct {
	conf   *config.Config
	router *gin.Engine
}

func New(c *config.Config) Server {
	return &server{
		conf:   c,
		router: gin.Default(),
	}
}

func (s *server) Start() error {
	server := s.createHttpServer()
	return server.ListenAndServe()
}

func (s *server) createHttpServer() *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", s.conf.Port),
		Handler: s.router.Handler(),
	}
}
