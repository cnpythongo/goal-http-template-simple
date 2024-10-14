package systemrolemenu

import (
	"github.com/gin-gonic/gin"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type ISystemRoleMenuHandler interface {
	list(c *gin.Context)
	detail(c *gin.Context)
	create(c *gin.Context)
	update(c *gin.Context)
	delete(c *gin.Context)
}

type SystemRoleMenuHandler struct {
	svc ISystemRoleMenuService
}

func NewSystemRoleMenuHandler(svc ISystemRoleMenuService) ISystemRoleMenuHandler {
	return &SystemRoleMenuHandler{svc: svc}
}

// list 角色菜单关联列表
// @Tags 角色菜单关联
// @Summary 角色菜单关联列表
// @Description 角色菜单关联列表
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqSystemRoleMenuList false "请求体"
// @Success 200 {object} render.JsonDataResp{data=render.RespPageJson{result=[]RespSystemRoleMenuItem}} "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/roles/menus/list [get]
func (h *SystemRoleMenuHandler) list(c *gin.Context) {
	var req ReqSystemRoleMenuList
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

// detail 角色菜单关联详情
// @Tags 角色菜单关联
// @Summary 角色菜单关联详情
// @Description 角色菜单关联详情
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqSystemRoleMenuDetail true "请求体"
// @Success 200 {object} RespSystemRoleMenuItem
// @Failure 400 {object} render.JsonDataResp
// @Security AdminAuth
// @Router /system/roles/menus/detail [get]
func (h *SystemRoleMenuHandler) detail(c *gin.Context) {
	var req ReqSystemRoleMenuDetail
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

// create 创建角色菜单关联
// @Tags 角色菜单关联
// @Summary 创建角色菜单关联
// @Description 创建角色菜单关联
// @Accept json
// @Produce json
// @Param data body ReqSystemRoleMenuCreate true "请求体"
// @Success 200 {object} render.JsonDataResp{data=string} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/roles/menus/create [post]
func (h *SystemRoleMenuHandler) create(c *gin.Context) {
	var payload ReqSystemRoleMenuCreate
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.PayloadError, nil)
		return
	}

	code, err := h.svc.Create(&payload)
	if err != nil {
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, "ok")
}

// update 更新角色菜单关联
// @Tags 角色菜单关联
// @Summary 更新角色菜单关联
// @Description 更新角色菜单关联
// @Accept json
// @Produce json
// @Param data body ReqSystemRoleMenuUpdate true "请求体"
// @Success 200 {object} render.JsonDataResp{data=RespSystemRoleMenuItem} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/roles/menus/update [post]
func (h *SystemRoleMenuHandler) update(c *gin.Context) {
	var payload ReqSystemRoleMenuUpdate
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

// delete 删除角色菜单关联
// @Tags 角色菜单关联
// @Summary 删除角色菜单关联
// @Description 删除角色菜单关联
// @Accept json
// @Produce json
// @Success 200 {object} render.JsonDataResp
// @Failure 400 {object} render.JsonDataResp
// @Security AdminAuth
// @Router /system/roles/menus/delete [post]
func (h *SystemRoleMenuHandler) delete(c *gin.Context) {
	var payload ReqSystemRoleMenuDelete
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
