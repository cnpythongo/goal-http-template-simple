package systemconfig

import (
	"github.com/gin-gonic/gin"
	"goal-app/pkg/render"
)

type ISystemConfigHandler interface {
	ConfigList(c *gin.Context)
}

type systemConfigHandler struct {
	svc ISystemConfigService
}

func NewSystemConfigHandler(svc ISystemConfigService) ISystemConfigHandler {
	return &systemConfigHandler{svc: svc}
}

func (h *systemConfigHandler) ConfigList(c *gin.Context) {
	render.Json(c, render.OK, "")
}
