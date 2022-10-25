package account

import (
	"github.com/cnpythongo/goal/model"
	resp "github.com/cnpythongo/goal/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserByUuid(c *gin.Context) {
	uid := c.Param("uid")
	result, err := model.GetUserByUuid(uid)
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
