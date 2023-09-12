package types

type (
	// RespEmptyJson 空数据返回结构
	RespEmptyJson struct {
		Code int    `json:"code" example:"0"` // 结果码：0-成功，其它-失败
		Msg  string `json:"msg" example:"ok"` // 消息, code不为0时，返回简单的错误描述
	}

	// RespFailJson 失败数据返回结构
	RespFailJson struct {
		Code  int    `json:"code" example:"1000"`       // 结果码：0-成功，其它-失败
		Msg   string `json:"msg" example:"error"`       // 消息, code不为0时，返回简单的错误描述
		Error string `json:"error" example:"not found"` // 具体的错误信息
	}
)
