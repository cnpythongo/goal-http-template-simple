package auth

import "github.com/cnpythongo/goal-tools/utils"

type (
	// ReqUserAuth 用户登录接口请求数据结构
	ReqUserAuth struct {
		Phone    string `json:"phone" binding:"required" example:"13800138000"` // 手机号
		Password string `json:"password" binding:"required" example:"123456"`   // 密码
	}

	// RespUserInfo 用户登录接口返回的用户信息数据结构
	RespUserInfo struct {
		UUID        string           `json:"uuid"`                        // 用户uuid
		Phone       string           `json:"phone" example:"138****8000"` // 带掩码的手机号
		LastLoginAt *utils.LocalTime `json:"last_login_at"`               // 最近的登录时间
		Nickname    string           `json:"nickname"`                    // 昵称
		Avatar      string           `json:"avatar"`                      // 头像
	}

	// RespUserAuth 用户登录接口返回数据结构
	RespUserAuth struct {
		Token      string       `json:"token"`       // 令牌
		ExpireTime string       `json:"expire_time"` // 过期时间
		User       RespUserInfo `json:"user"`        // 用户基本信息
	}
)
