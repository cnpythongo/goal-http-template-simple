package systemorg

import (
	"github.com/gin-gonic/gin"
	"goal-app/router/middleware"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewService()
	h := NewSystemOrgHandler(svc)

	r := route.Group("/api/v1/system/orgs")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/tree", h.Tree)
	r.POST("/create", h.Create)
	r.POST("/update", h.Update)
	r.POST("/delete", h.Delete)
	return r
}
