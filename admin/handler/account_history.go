package handler

import (
	"github.com/cnpythongo/goal/admin/service"
	"github.com/cnpythongo/goal/admin/types"
	"github.com/cnpythongo/goal/pkg/log"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/gin-gonic/gin"
)

// GetHistoryList 获取用户列表
// @Tags 用户管理
// @Summary 获取用户列表
// @Description 获取用户列表
// @Accept json
// @Produce json
// @Param data query types.ReqGetHistoryList false "请求体"
// @Success 200 {object} types.RespGetUserList
// @Failure 400 {object} types.RespFailJson
// @Security ApiKeyAuth
// @Router /account/history [get]
func GetHistoryList(c *gin.Context) {
	var req types.ReqGetUserList
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.GetLogger().Error(err)
		response.FailJson(c, response.ParamsError, err)
		return
	}
	result, code, err := service.NewAccountUserService().GetUserList(&req)
	if err != nil {
		response.FailJson(c, code, err)
		return
	}
	response.SuccessJson(c, result, nil)
}
