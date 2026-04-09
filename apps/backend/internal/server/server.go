package server

import (
	"fmt"
	"net/http"

	"github.com/0xlebogang/gonvy/backend/internal/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type svr struct {
	cfg    *config.Config
	db     *gorm.DB
	router *gin.Engine
}

func New(cfg *config.Config, db *gorm.DB) Server {
	return &svr{cfg: cfg, db: db, router: gin.Default()}
}

func (s *svr) Start() error {
	svr := s.createHttpServer()
	return svr.ListenAndServe()
}

func (s *svr) createHttpServer() *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", s.cfg.Port),
		Handler: s.router.Handler(),
	}
}
