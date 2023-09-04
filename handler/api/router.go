package api

import (
	"github.com/cnpythongo/goal/handler/api/account"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(r *gin.Engine) {
	g := r.Group("/api/account")
	// account login
	authHandler := account.NewAuthHandler()
	g.GET("/login", authHandler.Login)
	// user api
	userHandler := account.NewUserHandler()
	g.GET("/me", userHandler.GetUserByUuid)
	g.GET("/users/:uuid", userHandler.GetUserByUuid)
}
