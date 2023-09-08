package auth

import (
	"fmt"
	"github.com/cnpythongo/goal/pkg/jwt"
	resp "github.com/cnpythongo/goal/pkg/response"
	"github.com/cnpythongo/goal/service/account"
	"github.com/cnpythongo/goal/types"
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

// Login 登录接口
// @Tags 登录退出
// @Summary 登录
// @Description 后台管理系统登录接口
// @Accept json
// @Produce json
// @Param request body types.ReqAdminAuth true "请求体"
// @Success 200 {object} types.RespAdminAuth
// @failure 400 {object} types.RespFailJson
// @Router /account/login [post]
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

// Logout 退出接口
// @Tags 登录退出
// @Summary 退出
// @Description 退出后台管理系统，前端调用该接口，无需关注结果，自行清理掉请求头的 Authorization，页面跳转至首页
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} types.RespEmptyJson
// @Router /account/logout [post]
func (h *authHandler) Logout(c *gin.Context) {
	value, ok := c.Get(jwt.ContextUserKey)
	if ok {
		claims := value.(*jwt.Claims)
		userId := claims.ID
		token, _ := c.Get(jwt.ContextUserTokenKey)
		go func() {
			// 清理redis缓存, 设置token黑名单等操作
			fmt.Println(userId)
			fmt.Println(token)
		}()
	}
	resp.EmptyJsonResp(c, resp.SuccessCode)
}
