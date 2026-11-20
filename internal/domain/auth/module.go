package auth

import (
	"github.com/danielgtaylor/huma/v2"
	"gonvy.internal/ent"
	"gonvy.internal/internal/domain/users"
)

type Module interface {
	Register(api huma.API)
}

type module struct {
	handlers Handler
}

func BuildModule(db *ent.Client) Module {
	if db == nil {
		panic("db is nil")
	}
	userRepo := users.NewRepository(db)
	userSvc := users.NewService(userRepo)
	authRepo := newRepository(db)
	authSvc := newService(authRepo, userSvc)
	return &module{handlers: newhandler(authSvc)}
}

func (m *module) Register(api huma.API) {
	m.handlers.Mount(api)
}
