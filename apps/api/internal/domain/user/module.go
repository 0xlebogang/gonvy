package user

import "gorm.io/gorm"

type Module struct {
	controller Controller
}

func newModule(c Controller) *Module {
	return &Module{controller: c}
}

func BuildModule(db *gorm.DB) *Module {
	repo := NewRepository(db)
	service := NewService(repo)
	controller := NewController(service)
	return newModule(controller)
}
