package account

import (
	"fmt"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/log"
	resp "github.com/cnpythongo/goal/pkg/response"
	"github.com/cnpythongo/goal/service/account"
	"github.com/cnpythongo/goal/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IUserHandler interface {
	GetList(c *gin.Context)
	Create(c *gin.Context)
	BatchDelete(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Detail(c *gin.Context)
}

type userHandler struct {
	svc account.IUserService
}

func NewUserHandler() IUserHandler {
	return &userHandler{
		svc: account.NewUserService(),
	}
}

// GetList 获取用户列表
func (h *userHandler) GetList(c *gin.Context) {
	var req types.ReqGetUserList
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.GetLogger().Error(err)
		resp.FailJsonResp(c, resp.AccountQueryUserParamError, nil)
		return
	}
	page := req.Page
	size := req.Size
	conditions := map[string]interface{}{}
	result, count, err := h.svc.GetUserList(page, size, conditions)
	if err != nil {
		resp.FailJsonResp(c, resp.AccountQueryUserListError, nil)
		return
	}
	resp.SuccessJsonResp(c, result, map[string]interface{}{
		"count": count,
	})
}

// Create 创建用户
func (h *userHandler) Create(c *gin.Context) {
	payload := model.NewUser()
	err := c.ShouldBindJSON(payload)
	if err != nil {
		resp.FailJsonResp(c, resp.PayloadError, nil)
		return
	}
	eu, _ := h.svc.GetUserByPhone(payload.Phone)
	if eu != nil {
		resp.FailJsonResp(c, resp.AccountUserExistError, nil)
		return
	}
	ue, _ := h.svc.GetUserByEmail(payload.Email)
	if ue != nil {
		resp.FailJsonResp(c, resp.AccountEmailExistsError, nil)
		return
	}
	user, err := h.svc.CreateUser(payload)
	if err != nil {
		resp.FailJsonResp(c, resp.AccountCreateError, nil)
		return
	}
	resp.SuccessJsonResp(c, user, nil)
}

// BatchDelete 批量删除用户
func (h *userHandler) BatchDelete(c *gin.Context) {
	fmt.Println("BatchDeleteUserByUUID...")
	panic("implement me")
}

// Delete 删除单个用户
func (h *userHandler) Delete(c *gin.Context) {
	uuid := c.Param("uuid")
	err := h.svc.DeleteUserByUUID(uuid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.FailJsonResp(c, resp.AccountUserNotExistError, nil)
		} else {
			resp.FailJsonResp(c, resp.AccountQueryUserError, nil)
		}
		return
	}
	resp.EmptyJsonResp(c, resp.SuccessCode)
}

// 更新用户数据
func (h *userHandler) Update(c *gin.Context) {

}

// Detail 根据用户UUID获取用户详情
func (h *userHandler) Detail(c *gin.Context) {
	uuid := c.Param("uuid")
	result, err := h.svc.GetUserByUUID(uuid)
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
