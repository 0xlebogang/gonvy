package user

import "github.com/gin-gonic/gin"

type Controller interface {
	PostUser() gin.HandlerFunc
	GetAllUsers() gin.HandlerFunc
	GetUserByID() gin.HandlerFunc
	GetUserByEmail() gin.HandlerFunc
	PatchUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
}

type controller struct {
	service Service
}

func NewController(s Service) Controller {
	return &controller{service: s}
}

func (c *controller) PostUser() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (c *controller) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (c *controller) GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (c *controller) GetUserByEmail() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (c *controller) PatchUser() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (c *controller) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
