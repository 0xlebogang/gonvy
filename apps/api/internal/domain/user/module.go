package user

import (
	"github.com/0xlebogang/gonvy/api/internal/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	controller Controller
}

func newModule(c Controller) *Module {
	return &Module{controller: c}
}

func BuildModule(db *gorm.DB) *Module {
	hasher := common.NewHasher()

	repo := NewRepository(db)
	service := NewService(repo, hasher)
	controller := NewController(service)
	return newModule(controller)
}

func (m *Module) Register(r *gin.RouterGroup) {
}
