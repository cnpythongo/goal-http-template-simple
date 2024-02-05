package handler

import (
	"github.com/gin-gonic/gin"
	"goal-app/admin/service"
	"goal-app/admin/types"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

// Login 登录
// @Tags 登录退出
// @Summary 登录
// @Description 后台管理系统登录接口
// @Accept json
// @Produce json
// @Param data body types.ReqAdminAuth true "请求体"
// @Success 200 {object} render.RespJsonData{data=types.RespAdminAuth} "code不为0时表示有错误"
// @Failure 500
// @Router /account/login [post]
func Login(c *gin.Context) {
	var payload *types.ReqAdminAuth
	if err := c.ShouldBindJSON(&payload); err != nil {
		render.Json(c, render.PayloadError, err)
		return
	}

	result, code, err := service.NewAdminAuthService(c).Login(payload)
	if code != render.OK {
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, result)
}

// Logout 退出
// @Tags 登录退出
// @Summary 退出
// @Description 退出后台管理系统
// @Description 前端调用该接口，无需关注结果，自行清理掉请求头的 Authorization，页面跳转至首页
// @Description 后端可以执行清理redis缓存, 设置token黑名单等操作
// @Produce json
// @Security AdminAuth
// @Success 200 {object} render.RespJsonData
// @Router /account/logout [post]
func Logout(c *gin.Context) {
	go func() {
		err := service.NewAdminAuthService(c).Logout()
		if err != nil {
			log.GetLogger().Error(err)
		}
	}()
	render.Json(c, render.OK, "ok")
}
