package generator

import (
	"github.com/gin-gonic/gin"
	"goal-app/router/middleware"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := NewGeneratorService()
	h := NewGeneratorHandler(svc)

	r := route.Group("/api/v1/system/generator")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/list", h.List)
	r.GET("/tables", h.GetDbTableList)
	r.POST("/create", h.Create)
	r.POST("/update", h.Update)
	r.POST("/delete", h.Delete)
	r.GET("/preview", h.Preview)
	r.POST("/gencode", h.GenCode)

	r.GET("/tables/:id/columns", h.GetGenColumnList)
	r.POST("/tables/:id/columns/update", h.UpdateGenColumn)
	r.POST("/tables/:id/columns/delete", h.DeleteGenTableColumns)
	return r
}
