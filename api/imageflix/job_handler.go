package imageflix

import (
	"github.com/gin-gonic/gin"
	"goal-app/api/user"
	"goal-app/pkg/render"
)

type IImageFlixJobHandler interface {
	Start(c *gin.Context)
}

type imageFlixJobHandler struct {
	svc IImageFlixJobService
}

func NewImageFlixJobHandler(svc IImageFlixJobService) IImageFlixJobHandler {
	return &imageFlixJobHandler{
		svc: svc,
	}
}

// Create 创建任务
// @Tags ImageFlix-任务
// @Summary 当前登录用户上传图片任务
// @Description 当前登录用户上传图片任务
// @Produce json
// @Param file formData file true "文件流"
// @Success 200 {object} render.RespJsonData "code不为0时表示有错误"
// @Failure 500
// @Security APIAuth
// @Router /create [post]
func (h *imageFlixJobHandler) Create(c *gin.Context) {
	var req JobCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		render.Json(c, render.ParamsError, nil)
		return
	}

	_, ferr := c.FormFile("file")
	if ferr != nil {
		render.Json(c, render.Error, "请选择要上传文件")
		return
	}

	_, errCode := user.GetLoginCtxUser(c)
	if errCode != render.OK {
		render.Json(c, errCode, nil)
		return
	}

}

// Start 开始任务
// @Tags ImageFlix-任务
// @Summary 当前登录用户出发开始任务
// @Description 当前登录用户出发开始任务，成功后扣除点数
// @Produce json
// @Param data body JobStartReq true "请求体"
// @Success 200 {object} render.RespJsonData "code不为0时表示有错误"
// @Failure 500
// @Security APIAuth
// @Router /start [post]
func (h *imageFlixJobHandler) Start(c *gin.Context) {
	var req JobStartReq
	if err := c.ShouldBindJSON(&req); err != nil {
		render.Json(c, render.ParamsError, nil)
		return
	}

	ctxUser, errCode := user.GetLoginCtxUser(c)
	if errCode != render.OK {
		render.Json(c, errCode, nil)
		return
	}

	err := h.svc.Start(req.JobId, ctxUser.ID)
	if err != nil {
		render.Json(c, render.UpdateError, err.Error())
		return
	}

	render.Json(c, render.OK, "ok")
}
