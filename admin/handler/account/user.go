package account

import (
	"errors"
	"fmt"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/log"
	resp "github.com/cnpythongo/goal/pkg/response"
	"github.com/cnpythongo/goal/service/account"
	"github.com/cnpythongo/goal/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUserList 获取用户列表
func GetUserList(c *gin.Context) {
	var req types.ReqGetUserList
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.GetLogger().Error(err)
		resp.FailJsonResp(c, resp.AccountQueryUserParamError, nil)
		return
	}
	result, code := account.NewUserService(c).GetUserList(&req)
	if err != nil {
		resp.FailJsonResp(c, code, nil)
		return
	}
	resp.SuccessJsonResp(c, result, nil)
}

// UserCreate 创建用户
func UserCreate(c *gin.Context) {
	payload := model.NewUser()
	err := c.ShouldBindJSON(payload)
	if err != nil {
		resp.FailJsonResp(c, resp.PayloadError, nil)
		return
	}
	eu, _ := account.NewUserService(c).GetUserByPhone(payload.Phone)
	if eu != nil {
		resp.FailJsonResp(c, resp.AccountUserExistError, nil)
		return
	}
	ue, _ := account.NewUserService(c).GetUserByEmail(payload.Email)
	if ue != nil {
		resp.FailJsonResp(c, resp.AccountEmailExistsError, nil)
		return
	}
	user, err := account.NewUserService(c).CreateUser(payload)
	if err != nil {
		resp.FailJsonResp(c, resp.AccountCreateError, nil)
		return
	}
	resp.SuccessJsonResp(c, user, nil)
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
			resp.FailJsonResp(c, resp.AccountUserNotExistError, nil)
		} else {
			resp.FailJsonResp(c, resp.AccountQueryUserError, nil)
		}
		return
	}
	resp.EmptyJsonResp(c, resp.SuccessCode)
}

// UserUpdate 更新用户数据
func UserUpdate(c *gin.Context) {

}

// UserDetail 根据用户UUID获取用户详情
func UserDetail(c *gin.Context) {
	uuid := c.Param("uuid")
	result, err := account.NewUserService(c).GetUserByUUID(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.FailJsonResp(c, resp.AccountUserNotExistError, nil)
		} else {
			resp.FailJsonResp(c, resp.AccountQueryUserError, nil)
		}
		return
	}
	resp.SuccessJsonResp(c, result, nil)
}
