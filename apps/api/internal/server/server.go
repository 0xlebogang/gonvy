package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Config struct {
	Port string
}

type Server interface {
	createHttpServer() *http.Server
	Start() error
}

type server struct {
	conf   *Config
	db     *gorm.DB
	router *gin.Engine
}

func New(c *Config, db *gorm.DB) Server {
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
	svr := s.createHttpServer()
	return svr.ListenAndServe()
}
