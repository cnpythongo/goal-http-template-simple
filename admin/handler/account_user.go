package handler

import (
	"errors"
	"github.com/cnpythongo/goal/admin/service"
	"github.com/cnpythongo/goal/admin/types"
	"github.com/cnpythongo/goal/pkg/log"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/cnpythongo/goal/router/middleware"
	"github.com/gin-gonic/gin"
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
// @Success 200 {object} types.RespGetUserList
// @Failure 400 {object} types.RespFailJson
// @Security ApiKeyAuth
// @Router /account/user/list [get]
func (h *AccountUserHandler) List(c *gin.Context) {
	var req types.ReqGetUserList
	if err := c.ShouldBindQuery(&req); err != nil {
		log.GetLogger().Errorln(err)
		response.FailJson(c, response.ParamsError, err)
		return
	}

	result, code, err := h.svc.GetUserList(&req)
	if err != nil {
		log.GetLogger().Errorln(err)
		response.FailJson(c, code, err)
		return
	}
	response.SuccessJson(c, result, nil)
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
		response.FailJson(c, response.ParamsError, nil)
		return
	}

	result, code, err := h.svc.GetUserDetail(uuid)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Errorln(err)
		}
		response.FailJson(c, code, err)
		return
	}
	response.SuccessJson(c, result, nil)
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
		response.FailJson(c, response.PayloadError, nil)
		return
	}
	user, code, err := h.svc.CreateUser(&payload)
	if err != nil {
		log.GetLogger().Errorln(err)
		response.FailJson(c, code, err)
		return
	}
	response.SuccessJson(c, user, nil)
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
		response.FailJson(c, response.ParamsError, err)
		return
	}
	code, err := service.NewAccountUserService().UpdateUserByUUID(uuid, &payload)
	if err != nil {
		response.FailJson(c, code, err)
		return
	}
	response.SuccessJson(c, "ok", nil)
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
		response.FailJson(c, code, err)
		return
	}
	response.SuccessJson(c, "ok", nil)
}
