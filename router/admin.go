package router

import (
	"github.com/gin-gonic/gin"

	"github.com/cnpythongo/goal/handler/admin"
	"github.com/cnpythongo/goal/pkg/config"
)

func InitAdminRouters(cfg *config.Configuration) *gin.Engine {
	route := initDefaultRouter(cfg)
	admin.RegisterAdminRoutes(route)
	return route
}
