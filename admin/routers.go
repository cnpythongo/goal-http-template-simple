package admin

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"goal-app/admin/accountuser"
	"goal-app/admin/auth"
	"goal-app/admin/generator"
	"goal-app/admin/systemconfig"
	"goal-app/admin/systemlog"
	"goal-app/admin/systemmenu"
	"goal-app/admin/systemorg"
	"goal-app/admin/systemrole"
	"goal-app/admin/systemrolemenu"
	"goal-app/admin/systemroleuser"
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

	// account api
	_ = accountuser.RegisterRoute(route)

	// system api
	_ = generator.RegisterRoute(route)
	_ = systemconfig.RegisterRoute(route)
	_ = systemorg.RegisterRoute(route)
	_ = systemmenu.RegisterRoute(route)
	_ = systemlog.RegisterRoute(route)
	_ = systemrole.RegisterRoute(route)
	_ = systemrolemenu.RegisterRoute(route)
	_ = systemroleuser.RegisterRoute(route)

	return route
}
