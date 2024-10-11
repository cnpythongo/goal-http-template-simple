package systemroleuser

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewSystemRoleUserService()
	h := NewSystemRoleUserHandler(svc)

	r := route.Group("/api/v1/system/roles/users")
	r.GET("/list", h.list)
    r.GET("/detail", h.detail)
	r.POST("/create", h.create)
	r.POST("/update", h.update)
	r.POST("/delete", h.delete)
	return r
}
