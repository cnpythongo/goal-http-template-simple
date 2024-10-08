package systemlog

import (
	"github.com/gin-gonic/gin"
    "goal-app/pkg/log"
    "goal-app/pkg/render"
)

type ISystemLogHandler interface {
	list(c *gin.Context)
	detail(c *gin.Context)
	create(c *gin.Context)
	update(c *gin.Context)
	delete(c *gin.Context)
}

type SystemLogsHandler struct {
	svc ISystemLogService
}

func NewSystemLogHandler(svc ISystemLogService) ISystemLogHandler {
	return &SystemLogsHandler{svc: svc}
}

// list 系统日志列表
// @Tags 系统日志
// @Summary 系统日志列表
// @Description 系统日志列表
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqSystemLogList false "请求体"
// @Success 200 {object} render.JsonDataResp{data=render.RespPageJson{result=[]RespSystemLogItem}} "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/logs/list [get]
func (h *SystemLogsHandler) list(c *gin.Context) {
	var req ReqSystemLogList
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


// detail 系统日志详情
// @Tags 系统日志
// @Summary 系统日志详情
// @Description 系统日志详情
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqSystemLogDetail true "请求体"
// @Success 200 {object} RespSystemLogItem
// @Failure 400 {object} render.JsonDataResp
// @Security AdminAuth
// @Router /system/logs/detail [get]
func (h *SystemLogsHandler) detail(c *gin.Context) {
	var req ReqSystemLogDetail
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

// create 创建系统日志
// @Tags 系统日志
// @Summary 创建系统日志
// @Description 创建系统日志
// @Accept json
// @Produce json
// @Param data body ReqSystemLogCreate true "请求体"
// @Success 200 {object} render.JsonDataResp{data=RespSystemLogItem} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/logs/create [post]
func (h *SystemLogsHandler) create(c *gin.Context) {
	var payload ReqSystemLogCreate
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

// update 更新系统日志
// @Tags 系统日志
// @Summary 更新系统日志
// @Description 更新系统日志
// @Accept json
// @Produce json
// @Param data body ReqSystemLogUpdate true "请求体"
// @Success 200 {object} render.JsonDataResp{data=RespSystemLogItem} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/logs/update [post]
func (h *SystemLogsHandler) update(c *gin.Context) {
	var payload ReqSystemLogUpdate
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

// delete 删除系统日志
// @Tags 系统日志
// @Summary 删除系统日志
// @Description 删除系统日志
// @Accept json
// @Produce json
// @Success 200 {object} render.JsonDataResp
// @Failure 400 {object} render.JsonDataResp
// @Security AdminAuth
// @Router /system/logs/delete [post]
func (h *SystemLogsHandler) delete(c *gin.Context) {
	var payload ReqSystemLogDelete
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
