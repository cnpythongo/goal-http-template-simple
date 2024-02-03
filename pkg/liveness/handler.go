package liveness

import (
	"github.com/gin-gonic/gin"
	"goal-app/pkg/render"
)

func Ping(c *gin.Context) {
	render.Json(c, render.OK, "hello world")
}
