package admin

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"goal-app/admin/accounhistory"
	"goal-app/admin/accountuser"
	"goal-app/admin/auth"
	"goal-app/admin/systemconfig"
	"goal-app/pkg/config"
	"goal-app/router"
)

func InitAdminRouters(cfg *config.Configuration) *gin.Engine {
	route := router.InitDefaultRouter(cfg)
	if cfg.App.Debug {
		route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// auth
	_ = auth.RegisterRoute(route)

	// account user api
	_ = accountuser.RegisterRoute(route)
	_ = accounhistory.RegisterRoute(route)

	// system api
	_ = systemconfig.RegisterRoute(route)

	return route
}
