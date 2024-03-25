package user

import (
	"github.com/gin-gonic/gin"
	"goal-app/router/middleware"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewUserService()
	handler := NewUserHandler(svc)

	r := route.Group("/api/v1/users")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/me", handler.Me)
	r.GET("/me/profile", handler.Profile)
	r.POST("/me/profile/update", handler.UpdateProfile)
	r.GET("/:uuid", handler.GetUserInfoByUUID)
	return r
}
