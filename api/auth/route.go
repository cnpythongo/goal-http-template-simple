package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewAuthService()
	handler := NewAuthHandler(svc)

	r := route.Group("/api/v1")
	r.POST("/signup", handler.Signup)
	r.POST("/signin", handler.Signin)
	// r.POST("/login", handler.Signin)
	r.POST("/logout", handler.Logout)
	return r
}
