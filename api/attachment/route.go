package attachment

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewAttachmentService()
	handler := NewAttachmentHandler(svc)

	r := route.Group("/api/v1/attachments")
	// r.Use(middleware.JWTAuthenticationMiddleware())
	r.POST("", handler.Add)
	return r
}
