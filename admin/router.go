package admin

import (
	"github.com/cnpythongo/goal/admin/handler/account"
	"github.com/cnpythongo/goal/admin/handler/auth"
	"github.com/cnpythongo/goal/pkg/config"
	"github.com/cnpythongo/goal/router"
	"github.com/gin-gonic/gin"
)

func InitAdminRouters(cfg *config.Configuration) *gin.Engine {
	route := router.InitDefaultRouter(cfg)

	g := route.Group("/api/v1/account")
	// account login
	authHandler := auth.NewAuthHandler()
	g.POST("/login", authHandler.Login)
	// account user api
	userHandler := account.NewUserHandler()
	g.GET("/users", userHandler.GetList)
	g.POST("/users", userHandler.Create)
	g.PUT("/users", userHandler.BatchDelete)
	g.GET("/users/:uuid", userHandler.Detail)
	g.PATCH("/users/:uuid", userHandler.Update)
	g.DELETE("/users/:uuid", userHandler.Delete)

	return route
}
