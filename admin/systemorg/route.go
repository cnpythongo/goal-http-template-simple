package systemorg

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewSystemOrgService()
	handler := NewSystemOrgHandler(svc)

	r := route.Group("/api/v1/system/orgs")
	r.GET("/tree", handler.GetTreeData)
	r.POST("/create", handler.Create)
	r.POST("/update", handler.Update)
	r.POST("/delete", handler.Delete)
	return r
}
