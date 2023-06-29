package account

import (
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/common/log"
	resp "github.com/cnpythongo/goal/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func CreateUser(c *gin.Context) {
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

func GetUserById(c *gin.Context) {
	pk := c.Param("id")
	id, e := strconv.Atoi(pk)
	if e != nil {
		resp.FailJsonResp(c, resp.AccountUserIdError, nil)
		return
	}
	result, err := model.GetUserById(id)
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

func GetUserByUuid(c *gin.Context) {
	uid := c.Param("uid")
	result, err := model.GetUserByUuid(uid)
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
func GetUserList(c *gin.Context) {
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
