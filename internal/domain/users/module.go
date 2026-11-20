package users

import (
	"github.com/danielgtaylor/huma/v2"
	"gonvy.internal/ent"
)

type Module interface {
	Register(api huma.API)
}

type module struct {
	handler Handler
}

func BuildModule(db *ent.Client) Module {
	if db == nil {
		panic("db is nil")
	}
	repo := NewRepository(db)
	svc := NewService(repo)
	return &module{handler: Newhandler(svc)}
}

func (m module) Register(api huma.API) {
	m.handler.Mount(api)
}
