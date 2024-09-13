package systemorg

import (
	"github.com/gin-gonic/gin"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type ISystemOrgHandler interface {
	GetTreeData(c *gin.Context)
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

// GetTreeData 获取组织机构树结构数据
// @Tags 组织管理
// @Summary 获取组织机构树结构数据
// @Description 获取组织机构树结构数据
// @Accept json
// @Produce json
// @Success 200 {object} render.JsonDataResp{data=RespSystemOrgTree} "code不为0时表示有错误"
// @Failure 500
// @Router /system/orgs/tree [get]
func (h *systemOrgHandler) GetTreeData(c *gin.Context) {
	tree, err := h.svc.BuildOrgTree()
	if err != nil {
		render.Json(c, render.Error, err)
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
// @Success 200 {object} render.JsonDataResp "code不为0时表示有错误"
// @Failure 500
// @Router /system/orgs/create [post]
func (h *systemOrgHandler) Create(c *gin.Context) {
	var payload ReqSystemOrgCreate
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.PayloadError, nil)
		return
	}

	err := h.svc.CreateOrg(payload)
	if err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.CreateError, err)
		return
	}
	render.Json(c, render.OK, "ok")
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
// @Router /system/orgs/update [post]
func (h *systemOrgHandler) Update(c *gin.Context) {
	var payload ReqSystemOrgUpdate
	if err := c.ShouldBindJSON(&payload); err != nil {
		render.Json(c, render.ParamsError, err)
		return
	}
	err := h.svc.UpdateOrg(&payload)
	if err != nil {
		render.Json(c, render.UpdateError, err)
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
// @Router /system/orgs/delete [post]
func (h *systemOrgHandler) Delete(c *gin.Context) {
	var payload ReqSystemOrgId
	if err := c.ShouldBindJSON(&payload); err != nil {
		render.Json(c, render.ParamsError, err)
		return
	}
	if err := h.svc.DeleteOrg(payload.ID); err != nil {
		render.Json(c, render.DeleteError, err)
		return
	}
	render.Json(c, render.OK, "ok")
}
