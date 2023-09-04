package admin

import "github.com/gin-gonic/gin"

type IAuthHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type authHandler struct {
}

func NewAuthHandler() IAuthHandler {
	return &authHandler{}
}

func (handler authHandler) Login(c *gin.Context) {
	panic("implement me")
}

func (handler authHandler) Logout(c *gin.Context) {
	panic("implement me")
}
