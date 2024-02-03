package handler

import (
	"github.com/cnpythongo/goal/admin/service"
	"github.com/cnpythongo/goal/admin/types"
	"github.com/cnpythongo/goal/pkg/log"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/cnpythongo/goal/router/middleware"
	"github.com/gin-gonic/gin"
)

func AccountHistoryRouteRegister(route *gin.Engine) *gin.RouterGroup {
	svc := service.NewAccountUserService()
	handler := NewAccountUserHandler(svc)

	r := route.Group("/api/v1/account/history")
	r.Use(middleware.JWTAuthenticationMiddleware())
	r.GET("/list", handler.List)
	return r
}

type AccountHistoryHandler struct {
	svc service.IAccountUserService
}

func NewAccountHistoryHandler(svc service.IAccountUserService) *AccountHistoryHandler {
	return &AccountHistoryHandler{svc: svc}
}

// List 获取用户登录历史记录列表
// @Tags 用户管理
// @Summary 登录历史记录列表
// @Description 获取用户登录历史记录列表
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query types.ReqGetHistoryList false "请求体"
// @Success 200 {object} types.ReqGetHistoryList
// @Failure 400 {object} types.RespFailJson
// @Security ApiKeyAuth
// @Router /account/history/list [get]
func (h *AccountHistoryHandler) List(c *gin.Context) {
	var req types.ReqGetUserList
	if err := c.ShouldBindQuery(&req); err != nil {
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
