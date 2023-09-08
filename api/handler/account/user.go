package account

import (
	"github.com/gin-gonic/gin"
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
	//uuid := c.Param("uuid")
	//result, err := account.GetUserByUUID(uuid)
	//if err != nil {
	//	code := resp.AccountQueryUserError
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		code = resp.AccountUserNotExistError
	//	}
	//	resp.FailJsonResp(c, code, nil)
	//	return
	//}
	//resp.SuccessJsonResp(c, result, nil)
}