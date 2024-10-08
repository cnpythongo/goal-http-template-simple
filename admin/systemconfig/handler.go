package systemconfig

import (
	"github.com/gin-gonic/gin"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type ISystemConfigHandler interface {
	list(c *gin.Context)
	detail(c *gin.Context)
	create(c *gin.Context)
	update(c *gin.Context)
	delete(c *gin.Context)
}

type SystemConfigHandler struct {
	svc ISystemConfigService
}

func NewSystemConfigHandler(svc ISystemConfigService) ISystemConfigHandler {
	return &SystemConfigHandler{svc: svc}
}

// list 系统配置项列表
// @Tags 系统配置项
// @Summary 系统配置项列表
// @Description 系统配置项列表
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqSystemConfigList false "请求体"
// @Success 200 {object} render.JsonDataResp{data=render.RespPageJson{result=[]RespSystemConfigItem}} "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/config/list [get]
func (h *SystemConfigHandler) list(c *gin.Context) {
	var req ReqSystemConfigList
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

// detail 系统配置项详情
// @Tags 系统配置项
// @Summary 系统配置项详情
// @Description 系统配置项详情
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqSystemConfigDetail true "请求体"
// @Success 200 {object} RespSystemConfigItem
// @Failure 400 {object} render.JsonDataResp
// @Security AdminAuth
// @Router /system/config/detail [get]
func (h *SystemConfigHandler) detail(c *gin.Context) {
	var req ReqSystemConfigDetail
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

// create 创建系统配置项
// @Tags 系统配置项
// @Summary 创建系统配置项
// @Description 创建系统配置项
// @Accept json
// @Produce json
// @Param data body ReqSystemConfigCreate true "请求体"
// @Success 200 {object} render.JsonDataResp{data=RespSystemConfigItem} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/config/create [post]
func (h *SystemConfigHandler) create(c *gin.Context) {
	var payload ReqSystemConfigCreate
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

// update 更新系统配置项
// @Tags 系统配置项
// @Summary 更新系统配置项
// @Description 更新系统配置项
// @Accept json
// @Produce json
// @Param data body ReqSystemConfigUpdate true "请求体"
// @Success 200 {object} render.JsonDataResp{data=RespSystemConfigItem} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/config/update [post]
func (h *SystemConfigHandler) update(c *gin.Context) {
	var payload ReqSystemConfigUpdate
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

// delete 删除系统配置项
// @Tags 系统配置项
// @Summary 删除系统配置项
// @Description 删除系统配置项
// @Accept json
// @Produce json
// @Success 200 {object} render.JsonDataResp
// @Failure 400 {object} render.JsonDataResp
// @Security AdminAuth
// @Router /system/config/delete [post]
func (h *SystemConfigHandler) delete(c *gin.Context) {
	var payload ReqSystemConfigDelete
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
