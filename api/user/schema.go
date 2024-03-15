package user

import "github.com/cnpythongo/goal-tools/utils"

type (
	ReqGetUserInfo struct {
		UUID string `uri:"uuid" binding:"required"` // 用户UUID
	}

	// RespUserInfo 用户信息数据结构
	RespUserInfo struct {
		UUID        string           `json:"uuid"`                        // 用户uuid
		Phone       string           `json:"phone" example:"138****8000"` // 带掩码的手机号
		LastLoginAt *utils.LocalTime `json:"last_login_at,omitempty"`     // 最近的登录时间
		Nickname    string           `json:"nickname"`                    // 昵称
		Avatar      string           `json:"avatar"`                      // 头像
	}
)
