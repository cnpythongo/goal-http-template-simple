package accountuser

import (
	"github.com/gin-gonic/gin"
	"goal-app/router/middleware"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewUserService()
	handler := NewHandler(svc)

	r := route.Group("/api/v1/account/user")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/list", handler.List)
	r.GET("/detail", handler.Detail)
	r.POST("/create", handler.Create)
	r.POST("/update", handler.Update)
	r.POST("/delete", handler.Delete)
	return r
}
