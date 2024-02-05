package render

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespJsonData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Json(c *gin.Context, code int, result interface{}) {
	data := RespJsonData{
		Code: code,
		Msg:  GetCodeMsg(code),
		Data: result,
	}
	c.JSON(http.StatusOK, data)
}
