package admin

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"goal-app/admin/handler"
	"goal-app/pkg/config"
	"goal-app/router"
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

	// system api

	return route
}
