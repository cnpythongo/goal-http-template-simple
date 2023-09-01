package router

import (
	"github.com/gin-gonic/gin"
	
	"github.com/cnpythongo/goal/handler/api/account"
	"github.com/cnpythongo/goal/handler/liveness"
	"github.com/cnpythongo/goal/pkg/config"
)

func InitAPIRouters(cfg *config.Configuration) *gin.Engine {
	route := initDefaultRouter(cfg)

	// common test api
	apiGroup := route.Group("/api")
	apiGroup.GET("/ping", liveness.Ping)
	// api
	userHandler := account.NewUserHandler()
	userGroup := route.Group("/api/account")
	userGroup.GET("/me", userHandler.GetUserByUuid)
	userGroup.GET("/users/:uuid", userHandler.GetUserByUuid)
	return route
}
