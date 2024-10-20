package attachment

import (
	"github.com/gin-gonic/gin"
	"goal-app/router/middleware"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewAttachmentService()
	h := NewAttachmentHandler(svc)

	r := route.Group("/api/v1/attachments")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.POST("/create", h.Create)
	return r
}
