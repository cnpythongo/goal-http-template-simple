package accountuser

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"gorm.io/gorm"
)

type UserHandler struct {
	svc IUserService
}

func NewHandler(svc IUserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// List 获取用户列表
// @Tags 用户管理
// @Summary 获取用户列表
// @Description 获取用户列表
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqGetUserList false "请求体"
// @Success 200 {object} render.RespJsonData{data=RespGetUserList{result=[]RespUserDetail}} "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router /account/user/list [get]
func (h *UserHandler) List(c *gin.Context) {
	var req ReqGetUserList
	if err := c.ShouldBindQuery(&req); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.ParamsError, err)
		return
	}

	result, code, err := h.svc.GetUserList(&req)
	if err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, result)
}

// Detail 根据用户UUID获取用户详情
// @Tags 用户管理
// @Summary 通过用户UUID获取用户详情
// @Description 获取用户详情
// @Accept x-www-form-urlencoded
// @Produce json
// @Param uuid query string true "用户UUID"
// @Success 200 {object} RespUserDetail
// @Failure 400 {object} render.RespJsonData
// @Security AdminAuth
// @Router /account/user/detail [get]
func (h *UserHandler) Detail(c *gin.Context) {
	uuid := c.Query("uuid")
	if uuid == "" {
		render.Json(c, render.ParamsError, nil)
		return
	}

	result, code, err := h.svc.GetUserDetail(uuid)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Errorln(err)
		}
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, result)
}

// Create 创建用户
// @Tags 用户管理
// @Summary 创建用户
// @Description 创建新用户
// @Accept json
// @Produce json
// @Param data body ReqCreateUser true "请求体"
// @Success 200 {object} render.RespJsonData{data=RespUserDetail} "code不为0时表示错误"
// @Failure 500
// @Security AdminAuth
// @Router /account/user/create [post]
func (h *UserHandler) Create(c *gin.Context) {
	var payload ReqCreateUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.PayloadError, nil)
		return
	}
	if payload.Password != payload.PasswordConfirm {
		render.Json(c, render.ParamsError, nil)
		return
	}
	user, code, err := h.svc.CreateUser(&payload)
	if err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, user)
}

// Update 更新用户数据
// @Tags 用户管理
// @Summary 更新用户
// @Description 更新用户数据
// @Accept json
// @Produce json
// @Param data body ReqUpdateUser true "请求体"
// @Success 200 {object} render.RespJsonData
// @Failure 400 {object} render.RespJsonData
// @Security AdminAuth
// @Router /account/user/update [post]
func (h *UserHandler) Update(c *gin.Context) {
	uuid := c.Param("uuid")
	var payload ReqUpdateUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		render.Json(c, render.ParamsError, err)
		return
	}
	code, err := h.svc.UpdateUserByUUID(uuid, &payload)
	if err != nil {
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, "ok")
}

// Delete 删除用户
// @Tags 用户管理
// @Summary 删除用户
// @Description 删除单个用户
// @Accept json
// @Produce json
// @Success 200 {object} render.RespJsonData
// @Failure 400 {object} render.RespJsonData
// @Security AdminAuth
// @Router /account/user/delete [post]
func (h *UserHandler) Delete(c *gin.Context) {
	uuid := c.Param("uuid")
	code, err := h.svc.DeleteUserByUUID(uuid)
	if err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, "ok")
}
