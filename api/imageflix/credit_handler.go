package imageflix

import (
	"github.com/gin-gonic/gin"
	"goal-app/api/user"
	"goal-app/pkg/render"
)

type IImageFlixCreditHandler interface {
	UserCreditUsable(c *gin.Context)
	UserCreditReduce(c *gin.Context)
}

type imageFlixCreditHandler struct {
	svc IImageFlixCreditService
}

func NewImageFlixCreditHandler(svc IImageFlixCreditService) IImageFlixCreditHandler {
	return &imageFlixCreditHandler{
		svc: svc,
	}
}

// UserCreditUsable 用户可用资源包余额
// @Tags ImageFlix-资源包
// @Summary 获取当前登录用户可用资源包余额
// @Description 获取当前登录用户可用资源包余额
// @Produce json
// @Success 200 {object} render.JsonDataResp{data=UserCreditUsableResp} "code不为0时表示有错误"
// @Failure 500
// @Security APIAuth
// @Router /usable [get]
func (h *imageFlixCreditHandler) UserCreditUsable(c *gin.Context) {
	ctxUser, errCode := user.GetLoginCtxUser(c)
	if errCode != render.OK {
		render.Json(c, errCode, nil)
		return
	}

	userCredit, err := h.svc.GetUserImageFlixCredit(ctxUser.ID)
	if err != nil {
		render.Json(c, render.QueryError, err.Error())
		return
	}

	result := UserCreditUsableResp{Usable: userCredit.Usable}
	render.Json(c, render.OK, result)
}

// UserCreditReduce 用户使用资源包点数
// @Tags ImageFlix-资源包
// @Summary 获取当前登录用户可用资源包余额
// @Description 获取当前登录用户可用资源包余额
// @Produce json
// @Param data body UserCreditReduceReq true "请求体"
// @Success 200 {object} render.JsonDataResp "code不为0时表示有错误"
// @Failure 500
// @Security APIAuth
// @Router /reduce [post]
func (h *imageFlixCreditHandler) UserCreditReduce(c *gin.Context) {
	var req UserCreditReduceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		render.Json(c, render.ParamsError, nil)
		return
	}

	if req.Point == 0 {
		render.Json(c, render.OK, "ok")
		return
	}

	ctxUser, errCode := user.GetLoginCtxUser(c)
	if errCode != render.OK {
		render.Json(c, errCode, nil)
		return
	}

	userCredit, err := h.svc.GetUserImageFlixCredit(ctxUser.ID)
	if err != nil {
		render.Json(c, render.QueryError, err.Error())
		return
	}

	if userCredit.Usable < req.Point {
		render.Json(c, render.FlixCreditBalanceError, nil)
		return
	}

	err = h.svc.UpdateUserImageFlixCredit(ctxUser.ID, -req.Point)
	if err != nil {
		render.Json(c, render.FlixCreditBalanceReduceError, nil)
		return
	}

	render.Json(c, render.OK, "ok")
}
