package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewAuthService()
	handler := NewAuthHandler(svc)

	r := route.Group("/api/v1/account")
	r.POST("/login", handler.Login)
	r.POST("/logout", handler.Logout)
	return r
}
