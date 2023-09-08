package types

type (
	// ReqGetUserList 获取用户列表的请求参数体
	ReqGetUserList struct {
		Page             int    `json:"page" form:"page" example:"1"`                                                 // 页码
		Size             int    `json:"size" form:"size" example:"10"`                                                // 每页数量
		LastLoginAtStart string `json:"last_login_at_start" form:"last_login_at_start" example:"2023-09-01 01:30:59"` // 最近登录时间起始
		LastLoginAtEnd   string `json:"last_login_at_end" form:"last_login_at_end" example:"2023-09-01 22:59:59"`     // 最近登录时间截止
	}

	// RespUser 用户数据结构体
	RespUser struct {
		Phone string `json:"phone" example:"13800138000"` // 手机号
	}
)
