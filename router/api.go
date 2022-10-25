package router

import (
	"github.com/cnpythongo/goal/handler/api/account"
	"github.com/cnpythongo/goal/handler/liveness"
	"github.com/cnpythongo/goal/pkg/common/config"
	"github.com/gin-gonic/gin"
)

func InitAPIRouters(cfg *config.Configuration) *gin.Engine {
	route := initDefaultRouter(cfg)

	// common test api
	apiGroup := route.Group("/api")
	apiGroup.GET("/ping", liveness.Ping)
	// api
	userGroup := route.Group("/api/account")
	userGroup.GET("/me", account.GetUserByUuid)
	userGroup.GET("/users/:uid", account.GetUserByUuid)
	return route
}
