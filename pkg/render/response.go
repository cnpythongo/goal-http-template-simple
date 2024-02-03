package render

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Json(c *gin.Context, code int, result interface{}) {
	data := gin.H{
		"code": code,
		"msg":  GetCodeMsg(code),
	}
	httpCode := http.StatusOK
	if code != OK {
		data["error"] = result
		httpCode = http.StatusBadRequest
	} else {
		data["data"] = result
	}
	c.JSON(httpCode, data)
}
