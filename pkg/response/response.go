package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func EmptyJsonResp(c *gin.Context, code int) {
	data := gin.H{
		"code": code,
		"msg":  GetCodeMsg(code),
	}
	c.JSON(http.StatusOK, data)
}

func SuccessJsonResp(c *gin.Context, result interface{}, extends map[string]interface{}) {
	ret := gin.H{
		"code": SuccessCode,
		"msg":  GetCodeMsg(SuccessCode),
		"data": result,
	}
	if extends != nil {
		for key, value := range extends {
			ret[key] = value
		}
	}
	c.JSON(http.StatusOK, ret)
}

func FailJsonResp(c *gin.Context, code int, err interface{}) {
	data := gin.H{
		"code":  code,
		"msg":   GetCodeMsg(code),
		"error": err,
	}
	c.JSON(http.StatusOK, data)
}
