package account

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/log"
	resp "github.com/cnpythongo/goal/pkg/response"
)

type IUserHandler interface {
	CreateUser(c *gin.Context)
	GetUserByUUID(c *gin.Context)
	GetUserList(c *gin.Context)
	UpdateUserByUUID(c *gin.Context)
	DeleteUserByUUID(c *gin.Context)
	BatchDeleteUserByUUID(c *gin.Context)
}

type userHandler struct {
}

func NewUserHandler() IUserHandler {
	return &userHandler{}
}

// 创建用户
func (handler userHandler) CreateUser(c *gin.Context) {
	payload := model.NewAccountUser()
	err := c.ShouldBindJSON(payload)
	if err != nil {
		resp.FailJsonResp(c, resp.PayloadError, nil)
		return
	}
	eu, _ := model.GetUserByPhone(payload.Phone)
	if eu != nil {
		resp.FailJsonResp(c, resp.AccountUserExistError, nil)
		return
	}
	ue, _ := model.GetUserByEmail(payload.Email)
	if ue != nil {
		resp.FailJsonResp(c, resp.AccountEmailExistsError, nil)
		return
	}
	user, err := model.CreateUser(payload)
	if err != nil {
		resp.FailJsonResp(c, resp.AccountCreateError, nil)
		return
	}
	resp.SuccessJsonResp(c, user, nil)
}

// 根据用户UUID获取用户详情
func (handler userHandler) GetUserByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	result, err := model.GetUserByUUID(uuid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.FailJsonResp(c, resp.AccountUserNotExistError, nil)
		} else {
			resp.FailJsonResp(c, resp.AccountQueryUserError, nil)
		}
		return
	}
	resp.SuccessJsonResp(c, result, nil)
}

// 获取用户列表
func (handler userHandler) GetUserList(c *gin.Context) {
	var payload ReqGetUserListPayload
	err := c.ShouldBindQuery(&payload)
	if err != nil {
		log.GetLogger().Error(err)
		resp.FailJsonResp(c, resp.AccountQueryUserParamError, nil)
		return
	}
	page := payload.Page
	size := payload.Size
	// conditions := map[string]interface{}{}
	result, total, err := model.GetUserQueryset(page, size, nil)
	if err != nil {
		resp.FailJsonResp(c, resp.AccountQueryUserListError, nil)
		return
	}
	resp.SuccessJsonResp(c, result, map[string]interface{}{
		"total": total, "pages": utils.TotalPage(size, total),
	})
}

// 更新用户数据
func (handler userHandler) UpdateUserByUUID(c *gin.Context) {

}

// 删除单个用户
func (handler userHandler) DeleteUserByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	result, err := model.GetUserByUUID(uuid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.FailJsonResp(c, resp.AccountUserNotExistError, nil)
		} else {
			resp.FailJsonResp(c, resp.AccountQueryUserError, nil)
		}
		return
	}
	resp.SuccessJsonResp(c, result, nil)
}

// 批量删除用户
func (handler userHandler) BatchDeleteUserByUUID(c *gin.Context) {

}
