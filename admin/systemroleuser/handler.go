package systemroleuser

import (
	"github.com/gin-gonic/gin"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type ISystemRoleUserHandler interface {
	list(c *gin.Context)
	detail(c *gin.Context)
	create(c *gin.Context)
	update(c *gin.Context)
	delete(c *gin.Context)
}

type SystemRoleUsersHandler struct {
	svc ISystemRoleUserService
}

func NewSystemRoleUserHandler(svc ISystemRoleUserService) ISystemRoleUserHandler {
	return &SystemRoleUsersHandler{svc: svc}
}

// list 角色用户关联列表
// @Tags 角色用户关联
// @Summary 角色用户关联列表
// @Description 角色用户关联列表
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqSystemRoleUserList false "请求体"
// @Success 200 {object} render.JsonDataResp{data=render.RespPageJson{result=[]RespSystemRoleUserItem}} "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/roles/users/list [get]
func (h *SystemRoleUsersHandler) list(c *gin.Context) {
	var req ReqSystemRoleUserList
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

// detail 角色用户关联详情
// @Tags 角色用户关联
// @Summary 角色用户关联详情
// @Description 角色用户关联详情
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqSystemRoleUserDetail true "请求体"
// @Success 200 {object} RespSystemRoleUserItem
// @Failure 400 {object} render.JsonDataResp
// @Security AdminAuth
// @Router /system/roles/users/detail [get]
func (h *SystemRoleUsersHandler) detail(c *gin.Context) {
	var req ReqSystemRoleUserDetail
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

// create 创建角色用户关联
// @Tags 角色用户关联
// @Summary 创建角色用户关联
// @Description 创建角色用户关联
// @Accept json
// @Produce json
// @Param data body ReqSystemRoleUserCreate true "请求体"
// @Success 200 {object} render.JsonDataResp{data=RespSystemRoleUserItem} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/roles/users/create [post]
func (h *SystemRoleUsersHandler) create(c *gin.Context) {
	var payload ReqSystemRoleUserCreate
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

// update 更新角色用户关联
// @Tags 角色用户关联
// @Summary 更新角色用户关联
// @Description 更新角色用户关联
// @Accept json
// @Produce json
// @Param data body ReqSystemRoleUserUpdate true "请求体"
// @Success 200 {object} render.JsonDataResp{data=RespSystemRoleUserItem} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/roles/users/update [post]
func (h *SystemRoleUsersHandler) update(c *gin.Context) {
	var payload ReqSystemRoleUserUpdate
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

// delete 删除角色用户关联
// @Tags 角色用户关联
// @Summary 删除角色用户关联
// @Description 删除角色用户关联
// @Accept json
// @Produce json
// @Success 200 {object} render.JsonDataResp
// @Failure 400 {object} render.JsonDataResp
// @Security AdminAuth
// @Router /system/roles/users/delete [post]
func (h *SystemRoleUsersHandler) delete(c *gin.Context) {
	var payload ReqSystemRoleUserDelete
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
