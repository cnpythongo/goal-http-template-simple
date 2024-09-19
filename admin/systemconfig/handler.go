package systemconfig

import (
	"github.com/gin-gonic/gin"
	"goal-app/pkg/render"
)

type IHandler interface {
	GetList(c *gin.Context)
}

type handler struct {
	svc IService
}

func NewHandler(svc IService) IHandler {
	return &handler{svc: svc}
}

func (h *handler) GetList(c *gin.Context) {
	render.Json(c, render.OK, "")
}
