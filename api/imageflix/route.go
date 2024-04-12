package imageflix

import (
	"github.com/gin-gonic/gin"
	"goal-app/router/middleware"
)

func RegisterCreditRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewImageFlixCreditService()
	handler := NewImageFlixCreditHandler(svc)

	r := route.Group("/api/v1/flix/credits")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/usable", handler.UserCreditUsable)
	r.POST("/reduce", handler.UserCreditReduce)
	return r
}

func RegisterJobRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewImageFlixJobService()
	handler := NewImageFlixJobHandler(svc)

	r := route.Group("/api/v1/flix/job")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/start", handler.Start)
	return r
}
