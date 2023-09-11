package auth

import (
	"github.com/cnpythongo/goal/admin/service/account"
	"github.com/cnpythongo/goal/admin/types"
	"github.com/cnpythongo/goal/pkg/log"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/gin-gonic/gin"
)

// Login 登录
// @Tags 登录退出
// @Summary 登录
// @Description 后台管理系统登录接口
// @Accept json
// @Produce json
// @Param data body types.ReqAdminAuth true "请求体"
// @Success 200 {object} types.RespAdminAuth
// @Failure 400 {object} types.RespFailJson
// @Router /account/login [post]
func Login(c *gin.Context) {
	var payload *types.ReqAdminAuth
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.FailJsonResp(c, response.PayloadError, err)
		return
	}

	data, code, err := account.NewAdminAuthService(c).Login(payload)
	if code != response.SuccessCode {
		response.FailJsonResp(c, code, err)
		return
	}
	response.SuccessJsonResp(c, data, nil)
}

// Logout 退出
// @Tags 登录退出
// @Summary 退出
// @Description 退出后台管理系统
// @Description 前端调用该接口，无需关注结果，自行清理掉请求头的 Authorization，页面跳转至首页
// @Description 后端可以执行清理redis缓存, 设置token黑名单等操作
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} types.RespEmptyJson
// @Router /account/logout [post]
func Logout(c *gin.Context) {
	go func() {
		err := account.NewAdminAuthService(c).Logout()
		if err != nil {
			log.GetLogger().Error(err)
		}
	}()
	response.EmptyJsonResp(c, response.SuccessCode)
}
