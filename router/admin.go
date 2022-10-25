package router

import (
	"github.com/cnpythongo/goal/handler/admin/account"
	"github.com/gin-gonic/gin"

	"github.com/cnpythongo/goal/handler/liveness"
	"github.com/cnpythongo/goal/pkg/common/config"
)

func InitAdminRouters(cfg *config.Configuration) *gin.Engine {
	route := initDefaultRouter(cfg)

	// common test api
	apiGroup := route.Group("/api")
	apiGroup.GET("/ping", liveness.Ping)

	// admin api
	adminGroup := route.Group("/api/account")
	adminGroup.GET("/users", account.GetUserList)
	adminGroup.GET("/users/:uid", account.GetUserByUuid)
	adminGroup.POST("/users", account.CreateUser)
	return route
}
