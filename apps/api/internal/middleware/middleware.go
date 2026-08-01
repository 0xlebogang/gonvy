package middleware

import (
	"net/http"

	"github.com/0xlebogang/gonvy/api/internal/httperr"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	ErrorHandling() gin.HandlerFunc
}

type middleware struct{}

func New() Middleware {
	return &middleware{}
}

func (m *middleware) ErrorHandling() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last()
			apiErr := httperr.FromError(err)
			if apiErr.Status == http.StatusInternalServerError {
				apiErr.Log()
			}
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
		}
	}
}
