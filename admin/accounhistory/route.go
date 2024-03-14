package accounhistory

import (
	"github.com/gin-gonic/gin"
	"goal-app/router/middleware"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewHistoryService()
	handler := NewHistoryHandler(svc)

	r := route.Group("/api/v1/account/history")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/list", handler.HistoryList)
	return r
}
