package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"goal-app/api/attachment"
	"goal-app/api/auth"
	"goal-app/api/imageflix"
	"goal-app/api/user"
	"goal-app/pkg/config"
	"goal-app/router"
)

func InitAPIRouters(cfg *config.Configuration) *gin.Engine {
	route := router.InitDefaultRouter(cfg)
	if cfg.App.Debug {
		route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// account login
	_ = auth.RegisterRoute(route)

	// user api
	_ = user.RegisterRoute(route)

	// attachment api
	_ = attachment.RegisterRoute(route)

	// imageflix api
	_ = imageflix.RegisterCreditRoute(route)
	_ = imageflix.RegisterJobRoute(route)
	return route
}
