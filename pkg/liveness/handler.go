package liveness

import (
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	response.SuccessJson(c, "hello world", nil)
}
