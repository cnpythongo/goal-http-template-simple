package auth

import (
	"github.com/gin-gonic/gin"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type IAuthHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type authHandler struct {
	svc IAuthService
}

func NewAuthHandler(svc IAuthService) IAuthHandler {
	return &authHandler{svc: svc}
}

// Login 登录
// @Tags 登录退出
// @Summary 登录
// @Description 前台用户登录接口
// @Accept json
// @Produce json
// @Param data body ReqUserAuth true "请求体"
// @Success 200 {object} render.RespJsonData{data=RespUserAuth} "code不为0时表示有错误"
// @Failure 500
// @Router /login [post]
func (h *authHandler) Login(c *gin.Context) {
	var payload *ReqUserAuth
	if err := c.ShouldBindJSON(&payload); err != nil {
		render.Json(c, render.PayloadError, err)
		return
	}

	result, code, err := h.svc.Login(c, payload)
	if code != render.OK {
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, result)
}

// Logout 退出
// @Tags 登录退出
// @Summary 退出
// @Description 前台用户退出
// @Description 前端调用该接口，无需关注结果，自行清理掉请求头的 Authorization，页面跳转至首页
// @Description 后端可以执行清理redis缓存, 设置token黑名单等操作
// @Produce json
// @Security APIAuth
// @Success 200 {object} render.RespJsonData
// @Router /logout [post]
func (h *authHandler) Logout(c *gin.Context) {
	go func() {
		err := h.svc.Logout(c)
		if err != nil {
			log.GetLogger().Error(err)
		}
	}()
	render.Json(c, render.OK, "ok")
}
