package systemmenu

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type IHandler interface {
	GetTreeData(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type handler struct {
	svc IService
}

func NewHandler(svc IService) IHandler {
	return &handler{svc: svc}
}

// GetTreeData 获取菜单树结构数据
// @Tags 组织管理
// @Summary 获取菜单树结构数据
// @Description 获取菜单树结构数据
// @Accept json
// @Produce json
// @Success 200 {object} render.JsonDataResp{data=RespSystemMenuTree} "code不为0时表示有错误"
// @Failure 500
// @Router /system/menus/tree [get]
func (h *handler) GetTreeData(c *gin.Context) {
	tree, err := h.svc.BuildTree()
	if err != nil {
		render.Json(c, render.Error, err)
		return
	}
	render.Json(c, render.OK, tree)
}

// Create 创建菜单
// @Tags 组织管理
// @Summary 创建菜单
// @Description 创建菜单
// @Accept json
// @Produce json
// @Param data body ReqSystemMenuCreate true "请求体"
// @Success 200 {object} render.JsonDataResp{data=RespSystemMenuDetail} "code不为0时表示有错误"
// @Failure 500
// @Router /system/menus/create [post]
func (h *handler) Create(c *gin.Context) {
	var payload ReqSystemMenuCreate
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.PayloadError, nil)
		return
	}

	instance, code, err := h.svc.Create(&payload)
	if code != render.OK {
		render.Json(c, code, err)
		return
	}
	result := &RespSystemMenuDetail{}
	err = copier.Copy(&result, &instance)
	if err != nil {
		render.Json(c, render.DBAttributesCopyError, nil)
		return
	}
	render.Json(c, render.OK, result)
}

// Update 更新菜单
// @Tags 组织管理
// @Summary 更新菜单
// @Description 更新菜单
// @Accept json
// @Produce json
// @Param data body ReqSystemMenuUpdate true "请求体"
// @Success 200 {object} render.JsonDataResp "code不为0时表示有错误"
// @Failure 500
// @Router /system/menus/update [post]
func (h *handler) Update(c *gin.Context) {
	var payload ReqSystemMenuUpdate
	if err := c.ShouldBindJSON(&payload); err != nil {
		render.Json(c, render.PayloadError, err)
		return
	}
	err := h.svc.Update(&payload)
	if err != nil {
		render.Json(c, render.UpdateError, err)
		return
	}
	render.Json(c, render.OK, "ok")
}

// Delete 删除菜单
// @Tags 组织管理
// @Summary 删除菜单
// @Description 删除菜单
// @Accept json
// @Produce json
// @Param data body ReqSystemMenuIds true "请求体"
// @Success 200 {object} render.JsonDataResp "code不为0时表示有错误"
// @Failure 500
// @Router /system/menus/delete [post]
func (h *handler) Delete(c *gin.Context) {
	var payload ReqSystemMenuIds
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
