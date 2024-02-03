package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"goal-app/api/handler"
	"goal-app/pkg/config"
	"goal-app/router"
)

func InitAPIRouters(cfg *config.Configuration) *gin.Engine {
	route := router.InitDefaultRouter(cfg)
	if cfg.App.Debug {
		route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	g := route.Group("/api/v1/account")
	// account login
	g.POST("/login", handler.Login)
	// user api
	g.GET("/me", handler.GetUserByUuid)
	g.GET("/users/:uuid", handler.GetUserByUuid)

	return route
}
