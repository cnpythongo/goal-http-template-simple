package imageflix

import (
	"github.com/gin-gonic/gin"
	"goal-app/router/middleware"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	creditSvc := NewImageFlixCreditService()
	creditHandler := NewImageFlixCreditHandler(creditSvc)

	r := route.Group("/api/v1/flix/credits")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/usable", creditHandler.UserCredit)
	return r
}
