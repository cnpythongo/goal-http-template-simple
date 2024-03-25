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

	ReqUpdateUserProfile struct {
		UserId   int64  `json:"user_id"`   // 用户ID
		RealName string `json:"real_name"` // 真实姓名
		IDNumber string `json:"id_number"` // 身份证号
	}

	ReqUpdateUser struct {
		UUID     string `json:"uuid"`     // 用户UUID
		Nickname string `json:"nickname"` // 昵称
		Avatar   string `json:"avatar"`   // 头像
		Email    string `json:"email"`    // 邮箱
	}

	ReqUpdateUserPassword struct {
		UUID        string `json:"uuid"`         // 用户UUID
		OldPassword string `json:"old_password"` // 旧密码
		NewPassword string `json:"new_password"` // 新密码
	}
)
