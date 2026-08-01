package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	PostUser() gin.HandlerFunc
}

type controller struct {
	service Service
}

func NewController(s Service) Controller {
	return &controller{service: s}
}

func (c *controller) PostUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var json User
		if err := ctx.ShouldBind(&json); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "Bad request",
			})
			return
		}

		user, err := c.service.CreateUser(ctx.Request.Context(), &json)
		if err != nil {
			log.Printf("User registration failed: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "An unexpected error occured",
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"user": user,
		})
	}
}
