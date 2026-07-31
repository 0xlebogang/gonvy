package auth

import (
	"log"
	"net/http"

	"github.com/0xlebogang/gonvy/api/internal/domain/user"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	CreateUser() gin.HandlerFunc
	Login() gin.HandlerFunc
}

type controller struct {
	userSvc user.Service
}

func NewController(s user.Service) Controller {
	return &controller{userSvc: s}
}

func (c *controller) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var json user.User
		if err := ctx.ShouldBind(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid input",
			})
			return
		}

		user, err := c.userSvc.Register(ctx, &json)
		if err != nil {
			log.Printf("User registration failed: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "An Unexpected error occured",
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"user": user,
		})
	}
}

func (c *controller) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
