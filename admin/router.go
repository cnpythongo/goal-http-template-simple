package admin

import (
	"github.com/cnpythongo/goal/admin/handler"
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
	g.POST("/login", handler.Login)
	g.POST("/logout", handler.Logout)
	g.Use(middleware.JWTAuthenticationMiddleware())
	// account user api
	u := route.Group("/api/v1/account/users")
	u.POST("", handler.UserCreate)
	u.GET("", handler.GetUserList)
	u.PUT("", handler.UserBatchDelete)

	u.GET("/:uuid", handler.UserDetail)
	u.PUT("/:uuid", handler.UserUpdate)
	u.DELETE("/:uuid", handler.UserDelete)

	h := route.Group("/api/v1/account/history")
	h.GET("", handler.GetHistoryList)
	h.DELETE("/:user_id", handler.UserDelete)

	return route
}
