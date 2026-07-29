package user

import (
	"github.com/0xlebogang/gonvy/api/internal/password"
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
	passwordHandler := password.New()

	repo := NewRepository(db)
	service := NewService(repo, passwordHandler)
	controller := NewController(service)
	return newModule(controller)
}

func (m *Module) Register(r *gin.RouterGroup) {
	user := r.Group("/user")

	user.POST("", m.controller.PostUser())
	user.POST("/authenticate", m.controller.Authenticate())
	user.GET("", m.controller.GetAllUsers())
	user.GET("/:id", m.controller.GetUserByID())
	user.PATCH("/:id", m.controller.PatchUser())
	user.DELETE("/:id", m.controller.DeleteUser())
}
