package api

import (
	account2 "github.com/cnpythongo/goal/api/handler/account"
	auth2 "github.com/cnpythongo/goal/api/handler/auth"
	"github.com/cnpythongo/goal/pkg/config"
	"github.com/cnpythongo/goal/router"
	"github.com/gin-gonic/gin"
)

func InitAPIRouters(cfg *config.Configuration) *gin.Engine {
	route := router.InitDefaultRouter(cfg)

	g := route.Group("/api/account")
	// account login
	auth := auth2.NewAuthHandler()
	g.POST("/login", auth.Login)
	// user api
	userHandler := account2.NewUserHandler()
	g.GET("/me", userHandler.GetUserByUuid)
	g.GET("/users/:uuid", userHandler.GetUserByUuid)

	return route
}
