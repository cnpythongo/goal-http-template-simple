package {{{ .PackageName }}}

import (
	"github.com/gin-gonic/gin"
    "goal-app/pkg/log"
    "goal-app/pkg/render"
)

type I{{{ .EntityName }}}Handler interface {
	list(c *gin.Context)
	detail(c *gin.Context)
	create(c *gin.Context)
	update(c *gin.Context)
	delete(c *gin.Context)
	{{{- if eq .GenTpl "tree" }}}
	tree(c *gin.Context)
	{{{- end }}}
}

type {{{ toCamelCaseWithoutFirst .Name }}}Handler struct {
	svc I{{{ .EntityName }}}Service
}

func New{{{ .EntityName }}}Handler(svc I{{{ .EntityName }}}Service) I{{{ .EntityName }}}Handler {
	return &{{{ toCamelCaseWithoutFirst .Name }}}Handler{svc: svc}
}

// list {{{ .FunctionName }}}列表
// @Tags {{{ .FunctionName }}}
// @Summary {{{ .FunctionName }}}列表
// @Description {{{ .FunctionName }}}列表
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query Req{{{ .EntityName }}}List false "请求体"
// @Success 200 {object} render.JsonDataResp{data=render.RespPageJson{result=[]Resp{{{ .EntityName }}}Item}} "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router {{{ .GenPath }}}/list [get]
func (h *{{{ toCamelCaseWithoutFirst .Name }}}Handler) list(c *gin.Context) {
	var req Req{{{ .EntityName }}}List
	if err := c.ShouldBindQuery(&req); err != nil {
        log.GetLogger().Errorln(err)
        render.Json(c, render.PayloadError, nil)
        return
    }

    result, total, code, err := h.svc.List(&req)
    if err != nil {
        render.Json(c, code, err)
        return
    }

    var resp = &render.RespPageJson{
        Page:   req.Page,
        Limit:  req.Limit,
        Total:  total,
        Result: result,
    }
    render.Json(c, render.OK, resp)
}


// detail {{{ .FunctionName }}}详情
// @Tags {{{ .FunctionName }}}
// @Summary {{{ .FunctionName }}}详情
// @Description {{{ .FunctionName }}}详情
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query Req{{{ .EntityName }}}Detail true "请求体"
// @Success 200 {object} Resp{{{ .EntityName }}}Item
// @Failure 400 {object} render.JsonDataResp
// @Security AdminAuth
// @Router {{{ .GenPath }}}/detail [get]
func (h *{{{ toCamelCaseWithoutFirst .Name }}}Handler) detail(c *gin.Context) {
	var req Req{{{ .EntityName }}}Detail
	if err := c.ShouldBindQuery(&req); err != nil {
        log.GetLogger().Errorln(err)
        render.Json(c, render.PayloadError, nil)
        return
    }

    result, code, err := h.svc.Detail(&req)
	if err != nil {
        render.Json(c, code, err)
        return
    }
	render.Json(c, render.OK, result)
}

// create 创建{{{ .FunctionName }}}
// @Tags {{{ .FunctionName }}}
// @Summary 创建{{{ .FunctionName }}}
// @Description 创建{{{ .FunctionName }}}
// @Accept json
// @Produce json
// @Param data body Req{{{ .EntityName }}}Create true "请求体"
// @Success 200 {object} render.JsonDataResp{data=Resp{{{ .EntityName }}}Item} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router {{{ .GenPath }}}/create [post]
func (h *{{{ toCamelCaseWithoutFirst .Name }}}Handler) create(c *gin.Context) {
	var payload Req{{{ .EntityName }}}Create
	if err := c.ShouldBindJSON(&payload); err != nil {
        log.GetLogger().Errorln(err)
        render.Json(c, render.PayloadError, nil)
        return
    }

	result, code, err := h.svc.Create(&payload)
    if err != nil {
        render.Json(c, code, err)
        return
    }
    render.Json(c, render.OK, result)
}

// update 更新{{{ .FunctionName }}}
// @Tags {{{ .FunctionName }}}
// @Summary 更新{{{ .FunctionName }}}
// @Description 更新{{{ .FunctionName }}}
// @Accept json
// @Produce json
// @Param data body Req{{{ .EntityName }}}Update true "请求体"
// @Success 200 {object} render.JsonDataResp{data=Resp{{{ .EntityName }}}Item} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router {{{ .GenPath }}}/update [post]
func (h *{{{ toCamelCaseWithoutFirst .Name }}}Handler) update(c *gin.Context) {
	var payload Req{{{ .EntityName }}}Update
    if err := c.ShouldBindJSON(&payload); err != nil {
        log.GetLogger().Errorln(err)
        render.Json(c, render.PayloadError, nil)
        return
    }

    result, code, err := h.svc.Update(&payload)
    if err != nil {
        render.Json(c, code, err)
        return
    }
    render.Json(c, render.OK, result)
}

// delete 删除{{{ .FunctionName }}}
// @Tags {{{ .FunctionName }}}
// @Summary 删除{{{ .FunctionName }}}
// @Description 删除{{{ .FunctionName }}}
// @Accept json
// @Produce json
// @Success 200 {object} render.JsonDataResp
// @Failure 400 {object} render.JsonDataResp
// @Security AdminAuth
// @Router {{{ .GenPath }}}/delete [post]
func (h *{{{ toCamelCaseWithoutFirst .Name }}}Handler) delete(c *gin.Context) {
	var payload Req{{{ .EntityName }}}Delete
    if err := c.ShouldBindJSON(&payload); err != nil {
        log.GetLogger().Errorln(err)
        render.Json(c, render.PayloadError, nil)
        return
    }

    code, err := h.svc.Delete(&payload)
    if err != nil {
        render.Json(c, code, err)
        return
    }
    render.Json(c, render.OK, "OK")
}

{{{- if eq .GenTpl "tree" }}}
// tree {{{ .FunctionName }}}树结构数据
// @Tags {{{ .FunctionName }}}
// @Summary {{{ .FunctionName }}}树结构数据
// @Description {{{ .FunctionName }}}树结构数据
// @Accept json
// @Produce json
// @Success 200 {object} render.JsonDataResp{data=Resp{{{ .EntityName }}}Tree} "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router {{{ .GenPath }}}/tree [get]
func (h *{{{ toCamelCaseWithoutFirst .Name }}}Handler) tree(c *gin.Context) {
	var req Req{{{ .EntityName }}}Tree
	if err := c.ShouldBindQuery(&req); err != nil {
        log.GetLogger().Errorln(err)
        render.Json(c, render.PayloadError, nil)
        return
    }

    tree, code, err := h.svc.Tree(&req)
    if err != nil {
        render.Json(c, code, err)
        return
    }
    render.Json(c, render.OK, tree)
}
{{{- end }}}
