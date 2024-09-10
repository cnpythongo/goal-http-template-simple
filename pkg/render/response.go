package render

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Pagination 分页参数基础结构
type Pagination struct {
	Page           int   `form:"page" default:"1" example:"1"`    // 页码
	Limit          int   `form:"limit" default:"10" example:"10"` // 每页数量
	CreatedAtStart int64 `form:"created_at_start"`                // 数据创建开始区间
	CreatedAtEnd   int64 `form:"created_at_end"`                  // 数据创建结束区间
}

type RespPageJson struct {
	Page   int         `json:"page"`
	Limit  int         `json:"limit"`
	Total  int64       `json:"total"`
	Result interface{} `json:"result"`
}

type JsonDataResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Json(c *gin.Context, code int, result interface{}) {
	data := JsonDataResp{
		Code: code,
		Msg:  GetCodeMsg(code, result),
		Data: result,
	}
	c.JSON(http.StatusOK, data)
}
