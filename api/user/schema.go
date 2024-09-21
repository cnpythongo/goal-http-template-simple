package user

import (
	"goal-app/model"
)

type (
	UserInfoReq struct {
		UUID string `uri:"uuid" binding:"required"` // 用户UUID
	}

	// UserInfoResp 用户信息数据结构
	UserInfoResp struct {
		UUID          string `json:"uuid"`                        // 用户uuid
		Phone         string `json:"phone" example:"138****8000"` // 带掩码的手机号
		LastLoginTime int64  `json:"last_login_time"`             // 最近的登录时间(unix秒时间戳)
		Nickname      string `json:"nickname"`                    // 昵称
		Avatar        string `json:"avatar"`                      // 头像
	}

	UpdateUserProfileReq struct {
		UserId   uint64 `json:"user_id"`   // 用户ID
		RealName string `json:"real_name"` // 真实姓名
		IDNumber string `json:"id_number"` // 身份证号
	}

	UpdateUserReq struct {
		UUID      string           `json:"uuid"`                                 // 用户UUID
		Nickname  string           `json:"nickname,omitempty" example:"Tom"`     // 昵称
		Avatar    string           `json:"avatar,omitempty" example:"a/b/c.jpg"` // 用户头像URL
		Gender    model.UserGender `json:"gender,omitempty" example:"3"`         // 性别:3-保密,1-男,2-女
		Signature string           `json:"signature,omitempty" example:"haha"`   // 个性化签名
	}

	UpdateUserPasswordReq struct {
		UUID        string `json:"uuid"`         // 用户UUID
		OldPassword string `json:"old_password"` // 旧密码
		NewPassword string `json:"new_password"` // 新密码
	}
)
