package render

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Pagination 分页参数基础结构
type Pagination struct {
	Page      int      `form:"page" default:"1" example:"1"`    // 页码
	Limit     int      `form:"limit" default:"10" example:"10"` // 每页数量
	CreatedAt []string `form:"created_at[]"`                    // 数据创建时间起止区间
}

type RespPageJson struct {
	Page   int         `json:"page"`
	Limit  int         `json:"limit"`
	Total  int64       `json:"total"`
	Result interface{} `json:"result"`
}

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
