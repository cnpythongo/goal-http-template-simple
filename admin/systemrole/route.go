package systemrole

import (
	"github.com/gin-gonic/gin"
	"goal-app/router/middleware"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewSystemRoleService()
	h := NewSystemRoleHandler(svc)

	r := route.Group("/api/v1/system/roles")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/list", h.list)
	r.GET("/detail", h.detail)
	r.POST("/create", h.create)
	r.POST("/update", h.update)
	r.POST("/delete", h.delete)
	return r
}
