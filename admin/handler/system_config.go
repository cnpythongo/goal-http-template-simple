package handler

import (
	"github.com/gin-gonic/gin"
	"goal-app/admin/service"
	"goal-app/pkg/render"
	"goal-app/router/middleware"
)

func SystemConfigRouteRegister(route *gin.Engine) *gin.RouterGroup {
	svc := service.NewSystemConfigService()
	handler := NewSystemConfigHandler(svc)

	r := route.Group("/api/v1/system/config")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/list", handler.List)
	return r
}

type SystemConfigHandler struct {
	svc service.ISystemConfigService
}

func NewSystemConfigHandler(svc service.ISystemConfigService) *SystemConfigHandler {
	return &SystemConfigHandler{svc: svc}
}

func (h *SystemConfigHandler) List(c *gin.Context) {
	render.Json(c, render.OK, "")
}
