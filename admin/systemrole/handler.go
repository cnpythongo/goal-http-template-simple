package systemrole

import (
	"github.com/gin-gonic/gin"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type ISystemRoleHandler interface {
	list(c *gin.Context)
	detail(c *gin.Context)
	create(c *gin.Context)
	update(c *gin.Context)
	delete(c *gin.Context)
}

type SystemRolesHandler struct {
	svc ISystemRoleService
}

func NewSystemRoleHandler(svc ISystemRoleService) ISystemRoleHandler {
	return &SystemRolesHandler{svc: svc}
}

// list 角色管理列表
// @Tags 角色管理
// @Summary 角色管理列表
// @Description 角色管理列表
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqSystemRoleList false "请求体"
// @Success 200 {object} render.JsonDataResp{data=render.RespPageJson{result=[]RespSystemRoleItem}} "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/roles/list [get]
func (h *SystemRolesHandler) list(c *gin.Context) {
	var req ReqSystemRoleList
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

// detail 角色管理详情
// @Tags 角色管理
// @Summary 角色管理详情
// @Description 角色管理详情
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqSystemRoleDetail true "请求体"
// @Success 200 {object} RespSystemRoleItem
// @Failure 400 {object} render.JsonDataResp
// @Security AdminAuth
// @Router /system/roles/detail [get]
func (h *SystemRolesHandler) detail(c *gin.Context) {
	var req ReqSystemRoleDetail
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

// create 创建角色管理
// @Tags 角色管理
// @Summary 创建角色管理
// @Description 创建角色管理
// @Accept json
// @Produce json
// @Param data body ReqSystemRoleCreate true "请求体"
// @Success 200 {object} render.JsonDataResp{data=RespSystemRoleItem} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/roles/create [post]
func (h *SystemRolesHandler) create(c *gin.Context) {
	var payload ReqSystemRoleCreate
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

// update 更新角色管理
// @Tags 角色管理
// @Summary 更新角色管理
// @Description 更新角色管理
// @Accept json
// @Produce json
// @Param data body ReqSystemRoleUpdate true "请求体"
// @Success 200 {object} render.JsonDataResp{data=RespSystemRoleItem} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/roles/update [post]
func (h *SystemRolesHandler) update(c *gin.Context) {
	var payload ReqSystemRoleUpdate
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

// delete 删除角色管理
// @Tags 角色管理
// @Summary 删除角色管理
// @Description 删除角色管理
// @Accept json
// @Produce json
// @Success 200 {object} render.JsonDataResp
// @Failure 400 {object} render.JsonDataResp
// @Security AdminAuth
// @Router /system/roles/delete [post]
func (h *SystemRolesHandler) delete(c *gin.Context) {
	var payload ReqSystemRoleDelete
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
