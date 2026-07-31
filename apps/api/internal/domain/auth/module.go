package auth

import (
	"github.com/0xlebogang/gonvy/api/internal/common"
	"github.com/0xlebogang/gonvy/api/internal/domain/user"
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

	repo := user.NewRepository(db)

	userSvc := user.NewService(repo, hasher)
	controller := NewController(userSvc)
	return newModule(controller)
}

func (m *Module) Register(r *gin.RouterGroup) {
	auth := r.Group("/auth")

	auth.POST("/register", m.controller.CreateUser())
	auth.POST("/authenticate", m.controller.Login())
}
