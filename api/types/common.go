package types

type (
	// RespEmptyJson 空数据返回结构
	RespEmptyJson struct {
		Code int    `json:"code"` // 结果码：0-成功，其它-失败
		Msg  string `json:"msg"`  // 消息, code不为0时，返回简单的错误描述
	}

	// RespFailJson 失败数据返回结构
	RespFailJson struct {
		Code  int    `json:"code"`  // 结果码：0-成功，其它-失败
		Msg   string `json:"msg"`   // 消息, code不为0时，返回简单的错误描述
		Error string `json:"error"` // 具体的错误信息
	}

	// Pagination 分页参数基础体
	Pagination struct {
		Page           int    `form:"page" default:"1" example:"1"`                    // 页码
		Limit          int    `form:"limit" default:"10" example:"10"`                 // 每页数量
		CreatedAtStart string `form:"created_at_start" example:"2023-09-01 01:30:59"`  // 数据创建时间起始
		CreatedAtEnd   string `form:"last_login_at_end" example:"2023-09-01 22:59:59"` // 数据创建时间截止
	}
)
