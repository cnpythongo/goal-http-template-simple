package {{{ .PackageName }}}

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := New{{{ .EntityName }}}Service()
	h := New{{{ .EntityName }}}Handler(svc)

	r := route.Group("/api/v1/{{{ .GenPath }}}")
	r.GET("/list", h.list)
	r.GET("/tree", h.tree)
    r.GET("/detail", h.detail)
	r.POST("/create", h.create)
	r.POST("/update", h.update)
	r.POST("/delete", h.delete)
	return r
}
