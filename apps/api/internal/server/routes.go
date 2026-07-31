package server

import (
	"github.com/0xlebogang/gonvy/api/internal/domain/auth"
	"github.com/0xlebogang/gonvy/api/internal/domain/user"
)

func (s *server) SetupRoutes() {
	api := s.router.Group("/api")
	v1 := api.Group("/v1")

	authModule := auth.BuildModule(s.db)
	authModule.Register(v1)

	userModule := user.BuildModule(s.db)
	userModule.Register(v1)
}
