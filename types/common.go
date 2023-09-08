package types

type (
	// RespEmptyJson 空数据返回结构
	RespEmptyJson struct {
		Code int    `json:"code"` // 结果码：0-成功，其它-失败
		Msg  string `json:"msg"`  // 消息, code不为0时，返回简单的错误描述
	}

	// RespFailJson 失败数据返回结构
	RespFailJson struct {
		RespEmptyJson
		Error string `json:"error"` // 具体的错误信息
	}
)
