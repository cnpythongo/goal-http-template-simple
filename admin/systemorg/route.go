package systemorg

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewService()
	h := NewHandler(svc)

	r := route.Group("/api/v1/system/orgs")
	r.GET("/tree", h.GetTreeData)
	r.POST("/create", h.Create)
	r.POST("/update", h.Update)
	r.POST("/delete", h.Delete)
	return r
}
