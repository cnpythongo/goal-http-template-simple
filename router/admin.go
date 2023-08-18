package router

import (
	"github.com/cnpythongo/goal/handler/admin/account"
	"github.com/gin-gonic/gin"

	"github.com/cnpythongo/goal/handler/liveness"
	"github.com/cnpythongo/goal/pkg/config"
)

func InitAdminRouters(cfg *config.Configuration) *gin.Engine {
	route := initDefaultRouter(cfg)

	// common test api
	apiGroup := route.Group("/api")
	apiGroup.GET("/ping", liveness.Ping)
	// admin api
	authHandler := account.NewAuthHandler()
	adminGroup := route.Group("/api/account")
	adminGroup.GET("/login", authHandler.Login)
	// users
	userHandler := account.NewUserHandler()
	adminGroup.GET("/users", userHandler.GetUserList)
	adminGroup.POST("/users", userHandler.CreateUser)
	adminGroup.GET("/users/:uuid", userHandler.GetUserByUUID)
	adminGroup.PATCH("/users/:uuid", userHandler.UpdateUserByUUID)
	adminGroup.DELETE("/users/:uuid", userHandler.DeleteUserByUUID)
	adminGroup.POST("/users/delete", userHandler.BatchDeleteUserByUUID)
	return route
}
