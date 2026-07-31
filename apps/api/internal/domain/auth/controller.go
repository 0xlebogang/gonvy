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
	authSvc Service
}

func NewController(s user.Service, a Service) Controller {
	return &controller{userSvc: s, authSvc: a}
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

		user, err := c.userSvc.CreateUser(ctx, &json)
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
	return func(ctx *gin.Context) {
		var json UserLogin
		if err := ctx.ShouldBind(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid input",
			})
			return
		}

		tokens, err := c.authSvc.Authenticate(ctx, &json)
		if err != nil {
			log.Printf("User registration failed: %v\n", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "An unexpected error occured",
			})
			return
		}

		ctx.JSON(http.StatusOK, tokens)
	}
}
