package handler

import (
	"fmt"
	"github.com/cnpythongo/goal/admin/service"
	"github.com/cnpythongo/goal/admin/types"
	"github.com/cnpythongo/goal/pkg/log"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/gin-gonic/gin"
)

// GetUserList 获取用户列表
// @Tags 用户管理
// @Summary 获取用户列表
// @Description 获取用户列表
// @Accept json
// @Produce json
// @Param data query types.ReqGetUserList false "请求体"
// @Success 200 {object} types.RespGetUserList
// @Failure 400 {object} types.RespFailJson
// @Security ApiKeyAuth
// @Router /account/users [get]
func GetUserList(c *gin.Context) {
	var req types.ReqGetUserList
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.GetLogger().Error(err)
		response.FailJson(c, response.ParamsError, err)
		return
	}
	result, code, err := service.NewUserService(c).GetUserList(&req)
	if err != nil {
		response.FailJson(c, code, err)
		return
	}
	response.SuccessJson(c, result, nil)
}

// UserCreate 创建用户
// @Tags 用户管理
// @Summary 创建用户
// @Description 创建新用户
// @Accept json
// @Produce json
// @Param data body types.ReqCreateUser true "请求体"
// @Success 200 {object} types.RespUserDetail
// @Failure 400 {object} types.RespFailJson
// @Security ApiKeyAuth
// @Router /account/users [post]
func UserCreate(c *gin.Context) {
	var payload types.ReqCreateUser
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		response.FailJson(c, response.PayloadError, nil)
		return
	}
	user, code, err := service.NewUserService(c).CreateUser(&payload)
	if err != nil {
		response.FailJson(c, code, err)
		return
	}
	response.SuccessJson(c, user, nil)
}

// UserBatchDelete 批量删除用户
func UserBatchDelete(c *gin.Context) {
	fmt.Println("BatchDeleteUserByUUID...")
	panic("implement me")
}

// UserDelete 删除用户
// @Tags 用户管理
// @Summary 删除用户
// @Description 删除单个用户
// @Accept json
// @Produce json
// @Param uuid path string true "用户UUID"
// @Success 200 {object} types.RespEmptyJson
// @Failure 400 {object} types.RespFailJson
// @Security ApiKeyAuth
// @Router /account/users/{uuid} [delete]
func UserDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	code, err := service.NewUserService(c).DeleteUserByUUID(uuid)
	if err != nil {
		response.FailJson(c, code, err)
		return
	}
	response.EmptyJsonResp(c, response.SuccessCode)
}

// UserUpdate 更新用户数据
// @Tags 用户管理
// @Summary 更新用户
// @Description 更新用户数据
// @Accept json
// @Produce json
// @Param uuid path string true "用户UUID"
// @Param data body types.ReqUpdateUser true "请求体"
// @Success 200 {object} types.RespEmptyJson
// @Failure 400 {object} types.RespFailJson
// @Security ApiKeyAuth
// @Router /account/users/{uuid} [patch]
func UserUpdate(c *gin.Context) {
	uuid := c.Param("uuid")
	var payload types.ReqUpdateUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.FailJson(c, response.ParamsError, err)
		return
	}
	code, err := service.NewUserService(c).UpdateUserByUUID(uuid, &payload)
	if err != nil {
		response.FailJson(c, code, err)
		return
	}
	response.EmptyJsonResp(c, code)
}

// UserDetail 根据用户UUID获取用户详情
// @Tags 用户管理
// @Summary 通过用户UUID获取用户详情
// @Description 获取用户详情
// @Accept json
// @Produce json
// @Param uuid path string true "用户UUID"
// @Success 200 {object} types.RespUserDetail
// @Failure 400 {object} types.RespFailJson
// @Security ApiKeyAuth
// @Router /account/users/{uuid} [get]
func UserDetail(c *gin.Context) {
	uuid := c.Param("uuid")
	result, code, err := service.NewUserService(c).GetUserDetail(uuid)
	if err != nil {
		response.FailJson(c, code, nil)
		return
	}
	response.SuccessJson(c, result, nil)
}
