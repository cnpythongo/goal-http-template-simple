package admin

import (
	"github.com/cnpythongo/goal/admin/handler/account"
	"github.com/cnpythongo/goal/admin/handler/auth"
	"github.com/cnpythongo/goal/pkg/config"
	"github.com/cnpythongo/goal/router"
	"github.com/cnpythongo/goal/router/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitAdminRouters(cfg *config.Configuration) *gin.Engine {
	route := router.InitDefaultRouter(cfg)
	if cfg.App.Debug {
		route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	g := route.Group("/api/v1/account")
	// account login
	g.POST("/login", auth.Login)

	g.Use(middleware.JWTAuthenticationMiddleware())

	g.POST("/logout", auth.Logout)
	// account user api
	g.POST("/users", account.UserCreate)
	g.GET("/users", account.GetUserList)
	g.PUT("/users", account.UserBatchDelete)

	g.GET("/users/:uuid", account.UserDetail)
	g.PATCH("/users/:uuid", account.UserUpdate)
	g.DELETE("/users/:uuid", account.UserDelete)

	return route
}
