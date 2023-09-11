package account

import (
	"errors"
	"fmt"
	"github.com/cnpythongo/goal/admin/service/account"
	"github.com/cnpythongo/goal/admin/types"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/log"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		response.FailJsonResp(c, response.AccountQueryUserParamError, err)
		return
	}
	result, code, err := account.NewUserService(c).GetUserList(&req)
	if err != nil {
		response.FailJsonResp(c, code, err)
		return
	}
	response.SuccessJsonResp(c, result, nil)
}

// UserCreate 创建用户
func UserCreate(c *gin.Context) {
	payload := model.NewUser()
	err := c.ShouldBindJSON(payload)
	if err != nil {
		response.FailJsonResp(c, response.PayloadError, nil)
		return
	}
	eu, _ := account.NewUserService(c).GetUserByPhone(payload.Phone)
	if eu != nil {
		response.FailJsonResp(c, response.AccountUserExistError, nil)
		return
	}
	ue, _ := account.NewUserService(c).GetUserByEmail(payload.Email)
	if ue != nil {
		response.FailJsonResp(c, response.AccountEmailExistsError, nil)
		return
	}
	user, err := account.NewUserService(c).CreateUser(payload)
	if err != nil {
		response.FailJsonResp(c, response.AccountCreateError, nil)
		return
	}
	response.SuccessJsonResp(c, user, nil)
}

// UserBatchDelete 批量删除用户
func UserBatchDelete(c *gin.Context) {
	fmt.Println("BatchDeleteUserByUUID...")
	panic("implement me")
}

// UserDelete 删除单个用户
func UserDelete(c *gin.Context) {
	uuid := c.Param("uuid")
	err := account.NewUserService(c).DeleteUserByUUID(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.FailJsonResp(c, response.AccountUserNotExistError, nil)
		} else {
			log.GetLogger().Error(err)
			response.FailJsonResp(c, response.AccountQueryUserError, err)
		}
		return
	}
	response.EmptyJsonResp(c, response.SuccessCode)
}

// UserUpdate 更新用户数据
func UserUpdate(c *gin.Context) {

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
	result, code, err := account.NewUserService(c).GetUserByUUID(uuid)
	if err != nil {
		response.FailJsonResp(c, code, nil)
		return
	}
	response.SuccessJsonResp(c, result, nil)
}
