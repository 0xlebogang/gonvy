package auth

import (
	"errors"
	"net/http"

	"github.com/0xlebogang/gonvy/api/internal/domain/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
			_ = ctx.Error(ErrInvalidInput)
			return
		}

		user, err := c.userSvc.CreateUser(ctx.Request.Context(), &json)
		if err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				_ = ctx.Error(ErrEmailExists)
				return
			}
			_ = ctx.Error(err)
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
			_ = ctx.Error(ErrInvalidInput)
			return
		}

		tokens, err := c.authSvc.Authenticate(ctx.Request.Context(), &json)
		if err != nil {
			if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) || errors.Is(err, gorm.ErrRecordNotFound) {
				_ = ctx.Error(ErrInvalidCredentials)
				return
			}
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, tokens)
	}
}
