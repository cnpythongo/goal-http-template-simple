package router

import (
	"github.com/cnpythongo/goal/handler/admin"
	"github.com/gin-gonic/gin"

	"github.com/cnpythongo/goal/handler/liveness"
	"github.com/cnpythongo/goal/pkg/config"
)

func InitAdminRouters(cfg *config.Configuration) *gin.Engine {
	route := initDefaultRouter(cfg)

	// common test api
	apiGroup := route.Group("/api")
	apiGroup.GET("/ping", liveness.Ping)

	admin.RegisterAdminRoutes(route)

	return route
}
