package admin

import (
	"github.com/cnpythongo/goal/handler/admin/account"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(r *gin.Engine) {
	g := r.Group("/api/account")
	// account login
	authHandler := account.NewAuthHandler()
	g.GET("/login", authHandler.Login)
	// account user api
	userHandler := account.NewUserHandler()
	g.GET("/users", userHandler.GetUserList)
	g.GET("/users/:uuid", userHandler.GetUserByUUID)
	g.POST("/users", userHandler.CreateUser)
	g.PATCH("/users/:uuid", userHandler.UpdateUserByUUID)
	g.DELETE("/users/:uuid", userHandler.DeleteUserByUUID)
}
