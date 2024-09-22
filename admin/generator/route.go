package generator

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewService()
	h := NewHandler(svc)

	r := route.Group("/api/v1/system/generator")
	// r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/db-tables", h.GetDbTableList)
	r.POST("/import-table", h.ImportTable)
	r.GET("/gen-code", h.GenCode)
	return r
}
