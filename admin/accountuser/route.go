package accountuser

import (
	"github.com/gin-gonic/gin"
	"goal-app/router/middleware"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewUserService()
	handler := NewHandler(svc)

	r := route.Group("/api/v1/system/account/users")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/list", handler.List)
	r.GET("/detail", handler.Detail)
	r.POST("/create", handler.Create)
	r.POST("/edit", handler.Edit)
	r.POST("/reset-password", handler.ResetPwd)
	// r.POST("/update", handler.Update)
	r.POST("/delete", handler.Delete)

	// profile
	r.GET("/:uuid/profile", handler.Profile)
	r.GET("/:uuid/profile/update", handler.UpdateProfile)
	return r
}
