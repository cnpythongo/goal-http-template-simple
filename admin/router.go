package admin

import (
	"github.com/cnpythongo/goal/admin/handler"
	"github.com/cnpythongo/goal/pkg/config"
	"github.com/cnpythongo/goal/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitAdminRouters(cfg *config.Configuration) *gin.Engine {
	route := router.InitDefaultRouter(cfg)
	if cfg.App.Debug {
		route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	g := route.Group("/api/v1/account")
	// account login
	g.POST("/login", handler.Login)
	g.POST("/logout", handler.Logout)

	// account user api
	_ = handler.AccountUserRouteRegister(route)
	_ = handler.AccountHistoryRouteRegister(route)

	return route
}
