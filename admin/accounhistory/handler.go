package accounhistory

import (
	"github.com/gin-gonic/gin"
)

type IHistoryHandler interface {
	HistoryList(c *gin.Context)
}

type historyHandler struct {
	svc IHistoryService
}

func NewHistoryHandler(svc IHistoryService) IHistoryHandler {
	return &historyHandler{svc: svc}
}

// HistoryList 获取用户登录历史记录列表
// @Tags 用户管理
// @Summary 登录历史记录列表
// @Description 获取用户登录历史记录列表
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data query ReqGetHistoryList false "请求体"
// @Success 200 {object} ReqGetHistoryList
// @Failure 500
// @Security AdminAuth
// @Router /account/history/list [get]
func (h *historyHandler) HistoryList(c *gin.Context) {
	//var req ReqGetUserList
	//if err := c.ShouldBindQuery(&req); err != nil {
	//	log.GetLogger().Error(err)
	//	render.Json(c, render.ParamsError, err)
	//	return
	//}
	//result, code, err := service.NewAccountUserService().GetUserList(&req)
	//if err != nil {
	//	render.Json(c, code, err)
	//	return
	//}
	//render.Json(c, render.OK, result)
}
