package {{{ .PackageName }}}

import (
	"github.com/gin-gonic/gin"
    "goal-app/pkg/log"
    "goal-app/pkg/render"
)

type I{{{ .EntityName }}}Handler interface {
	list(c *gin.Context)
	tree(c *gin.Context)
	detail(c *gin.Context)
	create(c *gin.Context)
	update(c *gin.Context)
	delete(c *gin.Context)
}

type {{{ toCamelCaseWithoutFirst .Name }}}Handler struct {
	svc I{{{ .EntityName }}}Service
}

func New{{{ .EntityName }}}Handler(svc I{{{ .EntityName }}}Service) I{{{ .EntityName }}}Handler {
	return &{{{ toCamelCaseWithoutFirst .Name }}}Handler{svc: svc}
}

// list {{{ .EntityName }}}列表
// @Tags {{{ .EntityName }}}
// @Summary {{{ .EntityName }}}列表
// @Description {{{ .EntityName }}}列表
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query {{{ .EntityName }}}ListReq false "请求体"
// @Success 200 {object} render.JsonDataResp{data=render.RespPageJson{result=[]{{{ .EntityName }}}ItemResp}} "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router {{{ .GenPath }}}/list [get]
func (h *{{{ toCamelCaseWithoutFirst .Name }}}Handler) list(c *gin.Context) {
	var req {{{ .EntityName }}}ListReq
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

// tree {{{ .EntityName }}}树结构数据
// @Tags {{{ .EntityName }}}
// @Summary {{{ .EntityName }}}树结构数据
// @Description {{{ .EntityName }}}树结构数据
// @Accept json
// @Produce json
// @Success 200 {object} render.JsonDataResp{data={{{ .EntityName }}}TreeResp} "code不为0时表示有错误"
// @Failure 500
// @Router {{{ .GenPath }}}/tree [get]
func (h *{{{ toCamelCaseWithoutFirst .Name }}}Handler) tree(c *gin.Context) {
	var req {{{ .EntityName }}}TreeReq
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

// detail {{{ .EntityName }}}详情
// @Tags {{{ .EntityName }}}
// @Summary {{{ .EntityName }}}详情
// @Description {{{ .EntityName }}}详情
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query {{{ .EntityName }}}DetailReq true "请求体"
// @Success 200 {object} {{{ .EntityName }}}ItemResp
// @Failure 400 {object} render.JsonDataResp
// @Security AdminAuth
// @Router {{{ .GenPath }}}/detail [get]
func (h *{{{ toCamelCaseWithoutFirst .Name }}}Handler) detail(c *gin.Context) {
	var req {{{ .EntityName }}}DetailReq
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

// create 创建{{{ .EntityName }}}
// @Tags {{{ .EntityName }}}
// @Summary 创建{{{ .EntityName }}}
// @Description 创建{{{ .EntityName }}}
// @Accept json
// @Produce json
// @Param data body {{{ .EntityName }}}CreateReq true "请求体"
// @Success 200 {object} render.JsonDataResp{data={{{ .EntityName }}}ItemResp} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router {{{ .GenPath }}}/create [post]
func (h *{{{ toCamelCaseWithoutFirst .Name }}}Handler) create(c *gin.Context) {
	var payload {{{ .EntityName }}}CreateReq
	if err := c.ShouldBindJson(&req); err != nil {
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

// update 更新{{{ .EntityName }}}
// @Tags {{{ .EntityName }}}
// @Summary 更新{{{ .EntityName }}}
// @Description 更新{{{ .EntityName }}}
// @Accept json
// @Produce json
// @Param data body {{{ .EntityName }}}UpdateReq true "请求体"
// @Success 200 {object} render.JsonDataResp{data={{{ .EntityName }}}ItemResp} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router {{{ .GenPath }}}/update [post]
func (h *{{{ toCamelCaseWithoutFirst .Name }}}Handler) update(c *gin.Context) {
	var payload {{{ .EntityName }}}UpdateReq
    if err := c.ShouldBindJson(&req); err != nil {
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

// delete 删除{{{ .EntityName }}}
// @Tags {{{ .EntityName }}}
// @Summary 删除{{{ .EntityName }}}
// @Description 删除{{{ .EntityName }}}
// @Accept json
// @Produce json
// @Success 200 {object} render.JsonDataResp
// @Failure 400 {object} render.JsonDataResp
// @Security AdminAuth
// @Router {{{ .GenPath }}}/delete [post]
func (h *{{{ toCamelCaseWithoutFirst .Name }}}Handler) delete(c *gin.Context) {
	var payload {{{ .EntityName }}}DeleteReq
    if err := c.ShouldBindJson(&req); err != nil {
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
