package server

import (
	"github.com/0xlebogang/gonvy/api/internal/domain/auth"
	"github.com/gin-gonic/gin"
)

func (s *server) SetupV1Routes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	authModule := auth.BuildModule(s.conf, s.db)
	authModule.Register(v1)
}
