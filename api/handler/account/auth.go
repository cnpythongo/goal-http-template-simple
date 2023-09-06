package account

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

func (handler authHandler) Login(c *gin.Context) {
	panic("implement me")
}

func (handler authHandler) Logout(c *gin.Context) {
	panic("implement me")
}
