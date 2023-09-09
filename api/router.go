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
	authHandler := auth.NewAuthHandler()
	g.POST("/login", authHandler.Login)
	// user api
	userHandler := account.NewUserHandler()
	g.GET("/me", userHandler.GetUserByUuid)
	g.GET("/users/:uuid", userHandler.GetUserByUuid)

	return route
}
