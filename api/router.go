package api

import (
	"github.com/cnpythongo/goal/api/handler/account"
	"github.com/cnpythongo/goal/api/handler/auth"
	"github.com/cnpythongo/goal/pkg/config"
	"github.com/cnpythongo/goal/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitAPIRouters(cfg *config.Configuration) *gin.Engine {
	route := router.InitDefaultRouter(cfg)
	if cfg.App.Debug {
		route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	g := route.Group("/api/v1/account")
	// account login
	g.POST("/login", auth.Login)
	// user api
	g.GET("/me", account.GetUserByUuid)
	g.GET("/users/:uuid", account.GetUserByUuid)

	return route
}
