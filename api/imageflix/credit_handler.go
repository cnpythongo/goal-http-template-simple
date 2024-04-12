package imageflix

import (
	"github.com/gin-gonic/gin"
	"goal-app/api/user"
	"goal-app/pkg/render"
)

type IImageFlixCreditHandler interface {
	UserCreditUsable(c *gin.Context)
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
