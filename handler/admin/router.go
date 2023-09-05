package admin

import (
	"github.com/cnpythongo/goal/handler/admin/account"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(r *gin.Engine) {
	g := r.Group("/api/account")
	// account login
	auth := NewAuthHandler()
	g.GET("/login", auth.Login)
	// account user api
	userHandler := account.NewUserHandler()
	g.GET("/users", userHandler.GetList)
	g.POST("/users", userHandler.Create)
	g.PUT("/users", userHandler.BatchDelete)

	g.GET("/users/:uuid", userHandler.Detail)
	g.PATCH("/users/:uuid", userHandler.Update)
	g.DELETE("/users/:uuid", userHandler.Delete)
}
