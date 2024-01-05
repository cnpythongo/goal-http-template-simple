package account

import (
	"github.com/cnpythongo/goal/admin/service/account"
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

//// UserDelete 删除用户
//// @Tags 用户管理
//// @Summary 删除用户
//// @Description 删除单个用户
//// @Accept json
//// @Produce json
//// @Param uuid path string true "用户UUID"
//// @Success 200 {object} types.RespEmptyJson
//// @Failure 400 {object} types.RespFailJson
//// @Security ApiKeyAuth
//// @Router /account/users/{uuid} [delete]
//func UserDelete(c *gin.Context) {
//	uuid := c.Param("uuid")
//	code, err := account.NewUserService(c).DeleteUserByUUID(uuid)
//	if err != nil {
//		response.FailJsonResp(c, code, err)
//		return
//	}
//	response.EmptyJsonResp(c, response.SuccessCode)
//}
//
//// UserUpdate 更新用户数据
//// @Tags 用户管理
//// @Summary 更新用户
//// @Description 更新用户数据
//// @Accept json
//// @Produce json
//// @Param uuid path string true "用户UUID"
//// @Param data body types.ReqUpdateUser true "请求体"
//// @Success 200 {object} types.RespEmptyJson
//// @Failure 400 {object} types.RespFailJson
//// @Security ApiKeyAuth
//// @Router /account/users/{uuid} [patch]
//func UserUpdate(c *gin.Context) {
//	uuid := c.Param("uuid")
//	var payload types.ReqUpdateUser
//	if err := c.ShouldBindJSON(&payload); err != nil {
//		response.FailJsonResp(c, response.ParamsError, err)
//		return
//	}
//	code, err := account.NewUserService(c).UpdateUserByUUID(uuid, &payload)
//	if err != nil {
//		response.FailJsonResp(c, code, err)
//		return
//	}
//	response.EmptyJsonResp(c, code)
//}