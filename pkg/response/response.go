package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func jsonResp(c *gin.Context, code int, extends map[string]interface{}) {
	data := gin.H{
		"code": code,
		"msg":  GetCodeMsg(code),
	}
	if extends != nil {
		for key, value := range extends {
			data[key] = value
		}
	}
	c.JSON(http.StatusOK, data)
}

func SuccessJsonResp(c *gin.Context, result interface{}, extends map[string]interface{}) {
	if extends != nil {
		extends["data"] = result
	} else if result != nil {
		extends = map[string]interface{}{
			"data": result,
		}
	}
	jsonResp(c, SuccessCode, extends)
}

func FailJsonResp(c *gin.Context, code int, err interface{}) {
	jsonResp(c, code, map[string]interface{}{
		"error": err,
	})
}
