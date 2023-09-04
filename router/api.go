package router

import (
	"github.com/cnpythongo/goal/handler/api"
	"github.com/gin-gonic/gin"

	"github.com/cnpythongo/goal/handler/liveness"
	"github.com/cnpythongo/goal/pkg/config"
)

func InitAPIRouters(cfg *config.Configuration) *gin.Engine {
	route := initDefaultRouter(cfg)

	// common test api
	apiGroup := route.Group("/api")
	apiGroup.GET("/ping", liveness.Ping)

	api.RegisterApiRoutes(route)

	return route
}
