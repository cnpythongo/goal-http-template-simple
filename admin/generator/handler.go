package generator

import (
	"github.com/gin-gonic/gin"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type IHandler interface {
	GetDbTableList(c *gin.Context)
	ImportTable(c *gin.Context)
	GenCode(c *gin.Context)
}

type handler struct {
	svc IService
}

func NewHandler(svc IService) IHandler {
	return &handler{svc: svc}
}

// GetDbTableList 获取数据库表名称列表
// @Tags 系统管理--代码生成器
// @Summary 获取数据库表名称列表
// @Description 获取数据库表名称列表
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqDbTables false "请求体"
// @Success 200 {object} render.JsonDataResp{data=RespDbTableList{result=[]RespDbTableItem}} "code不为0时表示有错误"
// @Failure 500
// @Router /system/generator/db-tables [get]
func (h *handler) GetDbTableList(c *gin.Context) {
	var req ReqDbTables
	if err := c.ShouldBindQuery(&req); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.PayloadError, nil)
		return
	}
	result, code, err := h.svc.GetDbTableList(&req)
	if err != nil {
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, result)
}

func (h *handler) ImportTable(c *gin.Context) {
	var payload ReqImportTable
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.PayloadError, nil)
		return
	}
	render.Json(c, render.OK, "")
}

func (h *handler) GenCode(c *gin.Context) {
	render.Json(c, render.OK, "")
}
