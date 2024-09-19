package systemconfig

import (
	"github.com/gin-gonic/gin"
	"goal-app/router/middleware"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewService()
	h := NewHandler(svc)

	r := route.Group("/api/v1/system/config")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/list", h.GetList)
	return r
}
