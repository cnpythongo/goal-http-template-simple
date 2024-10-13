package systemmenu

import (
	"github.com/gin-gonic/gin"
	"goal-app/router/middleware"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewSystemMenuService()
	h := NewSystemMenuHandler(svc)

	r := route.Group("/api/v1/system/menus")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/list", h.list)
	r.GET("/tree", h.tree)
	r.GET("/detail", h.detail)
	r.POST("/create", h.create)
	r.POST("/update", h.update)
	r.POST("/delete", h.delete)
	return r
}
