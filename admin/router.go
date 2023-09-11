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
	g.POST("/logout", auth.Logout)
	g.Use(middleware.JWTAuthenticationMiddleware())
	// account user api
	u := route.Group("/api/v1/account/users")
	u.POST("", account.UserCreate)
	u.GET("", account.GetUserList)
	u.PUT("", account.UserBatchDelete)

	u.GET("/:uuid", account.UserDetail)
	u.PATCH("/:uuid", account.UserUpdate)
	u.DELETE("/:uuid", account.UserDelete)

	return route
}
