package attachment

import (
	"github.com/gin-gonic/gin"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type IAttachmentHandler interface {
	Add(c *gin.Context)
}

type attachmentHandler struct {
	svc IAttachmentService
}

func NewAttachmentHandler(svc IAttachmentService) IAttachmentHandler {
	return &attachmentHandler{svc: svc}
}

// Add 新增附件
// @Tags 通用
// @Summary 新增附件
// @Description 用户上传附件
// @Accept mpfd
// @Produce json
// @Param file formData file true "文件流"
// @Success 200 {object} AttachmentResp "http状态码是200，并且code是200, 表示正常返回；code不是200时表示有业务错误"
// // @Security ApiKeyAuth
// @Router /attachments [post]
func (h *attachmentHandler) Add(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		render.Json(c, render.PayloadError, err.Error())
		return
	}

	//ctxUser, errCode := user.GetLoginCtxUser(c)
	//if errCode != render.OK {
	//	render.Json(c, errCode, nil)
	//	return
	//}

	var payload AttachmentAddReq
	//payload.UserUuid = ctxUser.UUID
	//payload.UserId = ctxUser.ID
	payload.IP = c.ClientIP()

	res, _, err := h.svc.Add(payload, file)
	if err != nil {
		log.GetLogger().Errorf("attachmentHandler service upload err=[%+v]", err)
		render.Json(c, render.UploadFileError, err.Error())
		return
	}

	render.Json(c, render.OK, res)
	return
}
