package attachment

import (
	"github.com/gin-gonic/gin"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"goal-app/router/middleware"
)

type IAttachmentHandler interface {
	Create(c *gin.Context)
}

type attachmentHandler struct {
	svc IAttachmentService
}

func NewAttachmentHandler(svc IAttachmentService) IAttachmentHandler {
	return &attachmentHandler{svc: svc}
}

// Create 上传附件
// @Tags 通用
// @Summary 上传附件
// @Description 用户上传新附件
// @Accept mpfd
// @Produce json
// @Param biz formData string true "上传文件所属的业务名"
// @Param file formData file true "文件流"
// @Success 200 {object} RespAttachment "http状态码是200，并且code是200, 表示正常返回；code不是200时表示有业务错误"
// @Security AdminAuth
// @Router /attachments/create [post]
func (h *attachmentHandler) Create(c *gin.Context) {
	biz := c.PostForm("biz")
	if biz == "" {
		render.Json(c, render.ParamsError, "biz参数不能为空")
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		render.Json(c, render.PayloadError, err.Error())
		return
	}

	ctxUser, errCode := middleware.GetLoginCtxUser(c)
	if errCode != render.OK {
		render.Json(c, errCode, nil)
		return
	}

	payload := &ReqAttachmentCreate{
		UserId: ctxUser.ID,
		IP:     c.ClientIP(),
		Biz:    biz,
	}

	res, code, err := h.svc.Create(c, payload, file)
	if err != nil {
		log.GetLogger().Errorf("attachmentHandler service upload err = [%+v]", err)
		render.Json(c, code, err)
		return
	}

	render.Json(c, render.OK, res)
}
