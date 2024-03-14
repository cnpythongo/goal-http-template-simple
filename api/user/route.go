package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewUserService()
	handler := NewUserHandler(svc)

	r := route.Group("/api/v1/users")
	r.GET("me", handler.Me)
	r.GET("/:uuid", handler.GetUserByUUID)
	return r
}
