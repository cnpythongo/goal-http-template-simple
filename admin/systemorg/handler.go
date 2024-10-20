package systemorg

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type ISystemOrgHandler interface {
	Tree(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type systemOrgHandler struct {
	svc ISystemOrgService
}

func NewSystemOrgHandler(svc ISystemOrgService) ISystemOrgHandler {
	return &systemOrgHandler{svc: svc}
}

// Tree 获取组织机构树结构数据
// @Tags 组织管理
// @Summary 获取组织机构树结构数据
// @Description 获取组织机构树结构数据
// @Accept json
// @Produce json
// @Success 200 {object} render.JsonDataResp{data=[]RespSystemOrgTree} "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/orgs/tree [get]
func (h *systemOrgHandler) Tree(c *gin.Context) {
	tree, code, err := h.svc.Tree()
	if err != nil {
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, tree)
}

// Create 创建组织机构
// @Tags 组织管理
// @Summary 创建组织机构
// @Description 创建组织机构
// @Accept json
// @Produce json
// @Param data body ReqSystemOrgCreate true "请求体"
// @Success 200 {object} render.JsonDataResp{data=RespSystemOrgDetail} "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/orgs/create [post]
func (h *systemOrgHandler) Create(c *gin.Context) {
	var payload ReqSystemOrgCreate
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.PayloadError, nil)
		return
	}

	instance, code, err := h.svc.Create(&payload)
	if err != nil {
		render.Json(c, code, err)
		return
	}
	result := &RespSystemOrgDetail{}
	err = copier.Copy(result, &instance)
	if err != nil {
		log.GetLogger().Error(err)
		render.Json(c, render.DBAttributesCopyError, err)
		return
	}
	render.Json(c, render.OK, result)
}

// Update 更新组织机构
// @Tags 组织管理
// @Summary 更新组织机构
// @Description 更新组织机构
// @Accept json
// @Produce json
// @Param data body ReqSystemOrgUpdate true "请求体"
// @Success 200 {object} render.JsonDataResp "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/orgs/update [post]
func (h *systemOrgHandler) Update(c *gin.Context) {
	var payload ReqSystemOrgUpdate
	if err := c.ShouldBindJSON(&payload); err != nil {
		render.Json(c, render.PayloadError, err)
		return
	}
	code, err := h.svc.Update(&payload)
	if err != nil {
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, "ok")
}

// Delete 删除组织机构
// @Tags 组织管理
// @Summary 删除组织机构
// @Description 删除组织机构
// @Accept json
// @Produce json
// @Param data body ReqSystemOrgId true "请求体"
// @Success 200 {object} render.JsonDataResp "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router /system/orgs/delete [post]
func (h *systemOrgHandler) Delete(c *gin.Context) {
	var payload ReqSystemOrgIds
	if err := c.ShouldBindJSON(&payload); err != nil {
		render.Json(c, render.PayloadError, err)
		return
	}
	if err := h.svc.Delete(payload.IDs); err != nil {
		render.Json(c, render.DeleteError, err)
		return
	}
	render.Json(c, render.OK, "ok")
}
