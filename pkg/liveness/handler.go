package liveness

import (
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	response.SuccessJsonResp(c, "hello world", nil)
}
