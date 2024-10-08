package generator

import (
	"github.com/gin-gonic/gin"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"strings"
)

type IGeneratorHandler interface {
	List(c *gin.Context)
	GetDbTableList(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Preview(c *gin.Context)
	GenCode(c *gin.Context)

	GetGenColumnList(c *gin.Context)
	UpdateGenColumn(c *gin.Context)
	DeleteGenTableColumns(c *gin.Context)
}

type generatorHandler struct {
	svc IGeneratorService
}

func NewGeneratorHandler(svc IGeneratorService) IGeneratorHandler {
	return &generatorHandler{svc: svc}
}

// List 获取已导入的数据表列表
// @Tags 系统管理--代码生成器
// @Summary 获取已导入的数据表列表
// @Description 获取已导入的数据表列表
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqGenTableList false "请求体"
// @Success 200 {object} render.JsonDataResp{data=render.RespPageJson{result=[]RespGenTableItem}} "code不为0时表示有错误"
// @Failure 500
// // @Security AdminAuth
// @Router /system/generator/list [get]
func (h *generatorHandler) List(c *gin.Context) {
	var req ReqGenTableList
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

// GetDbTableList 获取数据库表名称列表
// @Tags 系统管理--代码生成器
// @Summary 获取数据库表名称列表
// @Description 获取数据库表名称列表
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqDbTableList false "请求体"
// @Success 200 {object} render.JsonDataResp{data=render.RespPageJson{result=[]RespDbTable}} "code不为0时表示有错误"
// @Failure 500
// // @Security AdminAuth
// @Router /system/generator/tables [get]
func (h *generatorHandler) GetDbTableList(c *gin.Context) {
	var req ReqDbTableList
	if err := c.ShouldBindQuery(&req); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.PayloadError, nil)
		return
	}
	result, total, code, err := h.svc.GetDbTableList(&req)
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

// Create 向 gen_tables 表中导入数据表信息
// @Tags 系统管理--代码生成器
// @Summary 向 gen_tables 表中导入数据表信息
// @Description 向 gen_tables 表中导入数据表信息
// @Accept json
// @Produce json
// @Param data body ReqGenTableCreate false "请求体"
// @Success 200 {object} render.JsonDataResp{data=string} "code不为0时表示有错误"
// @Failure 500
// // @Security AdminAuth
// @Router /system/generator/create [post]
func (h *generatorHandler) Create(c *gin.Context) {
	var req ReqGenTableCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.ParamsError, nil)
		return
	}
	code, err := h.svc.Create(&req)
	if err != nil {
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, "OK")
}

// Update 更新代码生成配置表数据
// @Tags 系统管理--代码生成器
// @Summary 更新代码生成配置表数据
// @Description 更新代码生成配置表数据
// @Accept json
// @Produce json
// @Param data body ReqPreview false "请求体"
// @Success 200 {object} render.JsonDataResp{data=string} "code不为0时表示有错误"
// @Failure 500
// // @Security AdminAuth
// @Router /system/generator/update [post]
func (h *generatorHandler) Update(c *gin.Context) {
	var req ReqUpdateGenTable
	if err := c.ShouldBindJSON(&req); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.ParamsError, nil)
		return
	}
	code, err := h.svc.Update(&req)
	if err != nil {
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, "")
}

// Delete 删除代码生成配置表数据
// @Tags 系统管理--代码生成器
// @Summary 删除代码生成配置表数据
// @Description 删除代码生成配置表数据
// @Accept json
// @Produce json
// @Param data body ReqDelTable false "请求体"
// @Success 200 {object} render.JsonDataResp{data=string} "code不为0时表示有错误"
// @Failure 500
// // @Security AdminAuth
// @Router /system/generator/delete [post]
func (h *generatorHandler) Delete(c *gin.Context) {
	var req ReqDelTable
	if err := c.ShouldBindJSON(&req); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.ParamsError, nil)
		return
	}
	code, err := h.svc.Delete(&req)
	if err != nil {
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, "")
}

// Preview 预览代码生成内容
// @Tags 系统管理--代码生成器
// @Summary 预览代码生成内容
// @Description 预览代码生成内容
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqPreview false "请求体"
// @Success 200 {object} render.JsonDataResp{data=[]RespPreviewItem} "code不为0时表示有错误"
// @Failure 500
// // @Security AdminAuth
// @Router /system/generator/preview [get]
func (h *generatorHandler) Preview(c *gin.Context) {
	var req ReqPreview
	if err := c.ShouldBindQuery(&req); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.ParamsError, nil)
		return
	}
	result, code, err := h.svc.Preview(&req)
	if err != nil {
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, result)
}

// GenCode 生成代码文件
// @Tags 系统管理--代码生成器
// @Summary 生成代码文件
// @Description 生成代码文件
// @Accept json
// @Produce json
// @Param data query ReqGenCode false "请求体"
// @Success 200 {object} render.JsonDataResp{data=string} "code不为0时表示有错误"
// @Failure 500
// // @Security AdminAuth
// @Router /system/generator/gencode [post]
func (h *generatorHandler) GenCode(c *gin.Context) {
	var payload ReqGenCode
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.ParamsError, nil)
		return
	}
	for _, table := range strings.Split(payload.Tables, ",") {
		code, err := h.svc.GenCode(table)
		if err != nil {
			render.Json(c, code, err)
			return
		}
	}
	render.Json(c, render.OK, "")
}

// GetGenColumnList 获取数据库表每个列的信息
// @Tags 系统管理--代码生成器
// @Summary 获取数据库表每个列的信息
// @Description 获取数据库表每个列的信息
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path int64 true "表格ID"
// @Success 200 {object} render.JsonDataResp{data=[]RespGenColumn} "code不为0时表示有错误"
// @Failure 500
// // @Security AdminAuth
// @Router /system/generator/tables/{id}/columns [get]
func (h *generatorHandler) GetGenColumnList(c *gin.Context) {
	var req ReqPreview
	if err := c.ShouldBindUri(&req); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.ParamsError, nil)
		return
	}
	result, code, err := h.svc.GetGenColumnList(req.ID)
	if err != nil {
		render.Json(c, code, err)
	}
	render.Json(c, render.OK, result)
}

// UpdateGenColumn 更新生成代码表列信息
// @Tags 系统管理--代码生成器
// @Summary 更新生成代码表列信息
// @Description 更新生成代码表列信息
// @Accept json
// @Produce json
// @Param id path int64 true "表格ID"
// @Success 200 {object} render.JsonDataResp "code不为0时表示有错误"
// @Failure 500
// // @Security AdminAuth
// @Router /system/generator/tables/{id}/columns/update [post]
func (h *generatorHandler) UpdateGenColumn(c *gin.Context) {
	var req ReqUpdateGenColumn
	if err := c.ShouldBindJSON(&req); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.ParamsError, nil)
		return
	}
	code, err := h.svc.UpdateGenColumn(&req)
	if err != nil {
		render.Json(c, code, err)
	}
	render.Json(c, render.OK, "")
}

// DeleteGenTableColumns 删除代码生成配置表的行属性数据
// @Tags 系统管理--代码生成器
// @Summary 删除代码生成配置表的行属性数据
// @Description 删除代码生成配置表的行属性数据
// @Accept json
// @Produce json
// @Param data body ReqDelGenTableColumn false "请求体"
// @Success 200 {object} render.JsonDataResp{data=string} "code不为0时表示有错误"
// @Failure 500
// // @Security AdminAuth
// @Router /system/generator/tables/{id}/columns/delete [post]
func (h *generatorHandler) DeleteGenTableColumns(c *gin.Context) {
	var req ReqDelGenTableColumn
	if err := c.ShouldBindJSON(&req); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.ParamsError, nil)
		return
	}
	code, err := h.svc.DeleteGenTableColumns(&req)
	if err != nil {
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, "")
}
