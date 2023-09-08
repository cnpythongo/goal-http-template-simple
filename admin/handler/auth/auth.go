package auth

import (
	resp "github.com/cnpythongo/goal/pkg/response"
	"github.com/cnpythongo/goal/service/account"
	"github.com/cnpythongo/goal/service/types"
	"github.com/gin-gonic/gin"
)

type (
	IAuthHandler interface {
		Login(c *gin.Context)
		Logout(c *gin.Context)
	}

	authHandler struct {
		svc account.IAdminAuthService
	}
)

func NewAuthHandler() IAuthHandler {
	return &authHandler{
		svc: account.NewAdminAuthService(),
	}
}

func (h *authHandler) Login(c *gin.Context) {
	var payload *types.ReqAdminAuth
	if err := c.ShouldBindJSON(&payload); err != nil {
		resp.FailJsonResp(c, resp.PayloadError, err)
		return
	}

	data, code := h.svc.AdminLogin(payload)
	if code != resp.SuccessCode {
		resp.FailJsonResp(c, code, nil)
		return
	}
	resp.SuccessJsonResp(c, data, nil)
}

func (h *authHandler) Logout(c *gin.Context) {
	panic("implement me")
}
