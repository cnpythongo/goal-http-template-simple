package systemconfig

import (
	"github.com/gin-gonic/gin"
	"goal-app/router/middleware"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewSystemConfigService()
	handler := NewSystemConfigHandler(svc)

	r := route.Group("/api/v1/system/config")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/list", handler.ConfigList)
	return r
}
