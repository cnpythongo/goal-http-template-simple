package account

import (
	"github.com/cnpythongo/goal/model"
	resp "github.com/cnpythongo/goal/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IUserHandler interface {
	GetUserByUuid(c *gin.Context)
}

type userHandler struct {
}

func NewUserHandler() IUserHandler {
	return &userHandler{}
}

func (handler userHandler) GetUserByUuid(c *gin.Context) {
	uuid := c.Param("uuid")
	result, err := model.GetUserByUUID(uuid)
	if err != nil {
		code := resp.AccountQueryUserError
		if err == gorm.ErrRecordNotFound {
			code = resp.AccountUserNotExistError
		}
		resp.FailJsonResp(c, code, nil)
		return
	}
	resp.SuccessJsonResp(c, result, nil)
}
