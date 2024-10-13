package {{{ .PackageName }}}

import (
	"github.com/gin-gonic/gin"
	"goal-app/router/middleware"
)

func RegisterRoute(route *gin.Engine) *gin.RouterGroup {
	svc := New{{{ .EntityName }}}Service()
	h := New{{{ .EntityName }}}Handler(svc)

	r := route.Group("/api/v1{{{ .GenPath }}}")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/list", h.list)
    r.GET("/detail", h.detail)
	r.POST("/create", h.create)
	r.POST("/update", h.update)
	r.POST("/delete", h.delete)
	{{{- if eq .GenTpl "tree" }}}
    r.GET("/tree", h.tree)
    {{{- end }}}
	return r
}
