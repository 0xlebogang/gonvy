package auth

import (
	"github.com/0xlebogang/gonvy/api/internal/common"
	"github.com/0xlebogang/gonvy/api/internal/config"
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

func BuildModule(c *config.EnvConfig, db *gorm.DB) *Module {
	hasher := common.NewHasher()
	token := NewToken(c)

	repo := user.NewRepository(db)

	userSvc := user.NewService(repo, hasher)
	authSvc := NewSvc(repo, token)

	controller := NewController(userSvc, authSvc)
	return newModule(controller)
}

func (m *Module) Register(r *gin.RouterGroup) {
	auth := r.Group("/auth")

	auth.POST("/register", m.controller.CreateUser())
	auth.POST("/authenticate", m.controller.Login())
}
