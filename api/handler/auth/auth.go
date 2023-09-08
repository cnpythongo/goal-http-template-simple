package auth

import (
	"github.com/cnpythongo/goal/service/account"
	"github.com/gin-gonic/gin"
)

type IAuthHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type authHandler struct {
	svc account.IApiAuthService
}

func NewAuthHandler() IAuthHandler {
	return &authHandler{
		svc: account.NewApiAuthService(),
	}
}

// @Summary 用户登录
// @Schemes
// @Description 前端应用的用户登录接口
// @Tags 登录
// @Accept json
// @Produce json
// @Success 200 {json} Helloworld
// @Router /account/login [post]
func (handler *authHandler) Login(c *gin.Context) {
	panic("implement me")
}

func (handler *authHandler) Logout(c *gin.Context) {
	panic("implement me")
}
