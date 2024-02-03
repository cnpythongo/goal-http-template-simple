package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goal-app/admin/service"
	"goal-app/admin/types"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"goal-app/router/middleware"
	"gorm.io/gorm"
)

func AccountUserRouteRegister(route *gin.Engine) *gin.RouterGroup {
	svc := service.NewAccountUserService()
	handler := NewAccountUserHandler(svc)

	r := route.Group("/api/v1/account/user")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/list", handler.List)
	r.GET("/detail", handler.Detail)
	r.POST("/create", handler.Create)
	r.POST("/update", handler.Update)
	r.POST("/delete", handler.Delete)
	return r
}

type AccountUserHandler struct {
	svc service.IAccountUserService
}

func NewAccountUserHandler(svc service.IAccountUserService) *AccountUserHandler {
	return &AccountUserHandler{svc: svc}
}

// List 获取用户列表
// @Tags 用户管理
// @Summary 获取用户列表
// @Description 获取用户列表
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query types.ReqGetUserList false "请求体"
// @Success 200 {object} types.ReqGetUserList
// @Failure 400 {object} types.RespFailJson
// @Security ApiKeyAuth
// @Router /account/user/list [get]
func (h *AccountUserHandler) List(c *gin.Context) {
	var req types.ReqGetUserList
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
// @Success 200 {object} types.RespUserDetail
// @Failure 400 {object} types.RespFailJson
// @Security ApiKeyAuth
// @Router /account/user/detail [get]
func (h *AccountUserHandler) Detail(c *gin.Context) {
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
// @Param data body types.ReqCreateUser true "请求体"
// @Success 200 {object} types.RespUserDetail
// @Failure 400 {object} types.RespFailJson
// @Security ApiKeyAuth
// @Router /account/user/create [post]
func (h *AccountUserHandler) Create(c *gin.Context) {
	var payload types.ReqCreateUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, render.PayloadError, nil)
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
// @Param data body types.ReqUpdateUser true "请求体"
// @Success 200 {object} types.RespEmptyJson
// @Failure 400 {object} types.RespFailJson
// @Security ApiKeyAuth
// @Router /account/user/update [post]
func (h *AccountUserHandler) Update(c *gin.Context) {
	uuid := c.Param("uuid")
	var payload types.ReqUpdateUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		render.Json(c, render.ParamsError, err)
		return
	}
	code, err := service.NewAccountUserService().UpdateUserByUUID(uuid, &payload)
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
// @Success 200 {object} types.RespEmptyJson
// @Failure 400 {object} types.RespFailJson
// @Security ApiKeyAuth
// @Router /account/user/delete [post]
func (h *AccountUserHandler) Delete(c *gin.Context) {
	uuid := c.Param("uuid")
	code, err := service.NewAccountUserService().DeleteUserByUUID(uuid)
	if err != nil {
		log.GetLogger().Errorln(err)
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, "ok")
}
