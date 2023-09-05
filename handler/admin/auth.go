package admin

import (
	"github.com/cnpythongo/goal/service/account"
	"github.com/gin-gonic/gin"
)

type IAuthHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type authHandler struct {
	svc account.IUserService
}

func NewAuthHandler() IAuthHandler {
	return &authHandler{
		svc: account.NewUserService(),
	}
}

func (h authHandler) Login(c *gin.Context) {
	panic("implement me")
}

func (h authHandler) Logout(c *gin.Context) {
	panic("implement me")
}
