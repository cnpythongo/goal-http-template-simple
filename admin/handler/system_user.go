package handler

import (
	"github.com/gin-gonic/gin"
	"goal-app/admin/service"
	_ "goal-app/admin/types"
	_ "goal-app/pkg/render"
	"goal-app/router/middleware"
)

func SystemUserRouteRegister(route *gin.Engine) *gin.RouterGroup {
	svc := service.NewAccountUserService()
	handler := NewSystemUserHandler(svc)

	r := route.Group("/api/v1/system/user")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/list", handler.List)
	return r
}

type SystemUserHandler struct {
	accountHandler *AccountUserHandler
	svc            service.IAccountUserService
}

func NewSystemUserHandler(svc service.IAccountUserService) *SystemUserHandler {
	return &SystemUserHandler{
		accountHandler: NewAccountUserHandler(svc),
		svc:            svc,
	}
}

// List 获取系统用户列表
// @Tags 系统管理
// @Summary 获取系统用户列表
// @Description 获取系统用户列表
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query types.ReqGetUserList false "请求体"
// @Success 200 {object} render.RespJsonData{data=types.RespGetUserList{result=[]types.RespUserDetail}} "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/user/list [get]
func (h *SystemUserHandler) List(c *gin.Context) {
	h.accountHandler.List(c)
}
