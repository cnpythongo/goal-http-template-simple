package systemmenu

import (
	"github.com/gin-gonic/gin"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type ISystemMenuHandler interface {
	list(c *gin.Context)
	tree(c *gin.Context)
	detail(c *gin.Context)
	create(c *gin.Context)
	update(c *gin.Context)
	delete(c *gin.Context)
}

type SystemMenusHandler struct {
	svc ISystemMenuService
}

func NewSystemMenuHandler(svc ISystemMenuService) ISystemMenuHandler {
	return &SystemMenusHandler{svc: svc}
}

// list 菜单管理列表
// @Tags 菜单管理
// @Summary 菜单管理列表
// @Description 菜单管理列表
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqSystemMenuList false "请求体"
// @Success 200 {object} render.JsonDataResp{data=render.RespPageJson{result=[]RespSystemMenuItem}} "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/menus/list [get]
func (h *SystemMenusHandler) list(c *gin.Context) {
	var req ReqSystemMenuList
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

// tree 菜单管理树结构数据
// @Tags 菜单管理
// @Summary 菜单管理树结构数据
// @Description 菜单管理树结构数据
// @Accept json
// @Produce json
// @Success 200 {object} render.JsonDataResp{data=[]RespSystemMenuTree} "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/menus/tree [get]
func (h *SystemMenusHandler) tree(c *gin.Context) {
	var req ReqSystemMenuTree
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

// detail 菜单管理详情
// @Tags 菜单管理
// @Summary 菜单管理详情
// @Description 菜单管理详情
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqSystemMenuDetail true "请求体"
// @Success 200 {object} RespSystemMenuItem
// @Failure 400 {object} render.JsonDataResp
// @Security AdminAuth
// @Router /system/menus/detail [get]
func (h *SystemMenusHandler) detail(c *gin.Context) {
	var req ReqSystemMenuDetail
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

// create 创建菜单管理
// @Tags 菜单管理
// @Summary 创建菜单管理
// @Description 创建菜单管理
// @Accept json
// @Produce json
// @Param data body ReqSystemMenuCreate true "请求体"
// @Success 200 {object} render.JsonDataResp{data=RespSystemMenuItem} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/menus/create [post]
func (h *SystemMenusHandler) create(c *gin.Context) {
	var payload ReqSystemMenuCreate
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

// update 更新菜单管理
// @Tags 菜单管理
// @Summary 更新菜单管理
// @Description 更新菜单管理
// @Accept json
// @Produce json
// @Param data body ReqSystemMenuUpdate true "请求体"
// @Success 200 {object} render.JsonDataResp{data=RespSystemMenuItem} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/menus/update [post]
func (h *SystemMenusHandler) update(c *gin.Context) {
	var payload ReqSystemMenuUpdate
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

// delete 删除菜单管理
// @Tags 菜单管理
// @Summary 删除菜单管理
// @Description 删除菜单管理
// @Accept json
// @Produce json
// @Success 200 {object} render.JsonDataResp
// @Failure 400 {object} render.JsonDataResp
// @Security AdminAuth
// @Router /system/menus/delete [post]
func (h *SystemMenusHandler) delete(c *gin.Context) {
	var payload ReqSystemMenuDelete
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
