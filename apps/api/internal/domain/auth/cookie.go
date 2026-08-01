package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ACCESS_TOKEN_COOKIE    = "access_token"
	REFRESH_TOKEN_COOKIE   = "refresh_token"
	ACCESS_TOKEN_DURATION  = 60 * 60
	REFRESH_TOKEN_DURATION = 7 * 24 * 60 * 60
)

func setTokenCookies(ctx *gin.Context, accessToken, refreshToken string) {
	isProduction := gin.Mode() == gin.ReleaseMode

	ctx.SetCookie(
		ACCESS_TOKEN_COOKIE,
		accessToken,
		ACCESS_TOKEN_DURATION,
		"/",
		"",
		isProduction,
		true,
	)

	ctx.SetCookie(
		REFRESH_TOKEN_COOKIE,
		refreshToken,
		REFRESH_TOKEN_DURATION,
		"/api/v1/auth/refresh",
		"",
		isProduction,
		true,
	)

	ctx.SetSameSite(http.SameSiteLaxMode)
}
