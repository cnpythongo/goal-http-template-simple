package router

import (
	"github.com/gin-gonic/gin"

	"github.com/cnpythongo/goal/handler/api"
	"github.com/cnpythongo/goal/pkg/config"
)

func InitAPIRouters(cfg *config.Configuration) *gin.Engine {
	route := initDefaultRouter(cfg)
	api.RegisterApiRoutes(route)
	return route
}
