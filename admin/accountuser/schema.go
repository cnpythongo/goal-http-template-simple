package accountuser

import (
	"goal-app/model"
	"goal-app/pkg/render"
)

type (
	// ReqGetUserList 获取用户列表的请求参数体
	ReqGetUserList struct {
		render.Pagination
		UUID               string                 `form:"uuid" example:"826d6b1aa64d471d822d667e92218158"` // 用户UUID,精确匹配
		Phone              string                 `form:"phone" example:"13800138000"`                     // 手机号,模糊查询
		Email              string                 `form:"email" example:"abc@abc.com"`                     // 邮箱,模糊查询
		Nickname           string                 `form:"nickname" example:"Tom"`                          // 昵称,模糊查询
		Status             []model.UserStatusType `form:"status[]" example:"FREEZE,ACTIVE"`                // 用户状态
		LastLoginTimeStart int64                  `form:"last_login_time_start"`                           // 最近登录开始时间
		LastLoginTimeEnd   int64                  `form:"last_login_time_end"`                             // 最近登录结束时间
		IsAdmin            *int64                 `form:"is_admin" example:"1"`                            // 是否admin, 1 or 0
	}

	// RespUserBasic 用户基础数据结构体
	RespUserBasic struct {
		UUID          string               `json:"uuid" example:"826d6b1aa64d471d822d667e92218158"` // 用户UUID,32位字符串
		Phone         string               `json:"phone" example:"13800138000"`                     // 手机号
		Email         string               `json:"email" example:"abc@abc.com"`                     // 邮箱
		Nickname      string               `json:"nickname" example:"Tom"`                          // 昵称
		Status        model.UserStatusType `json:"status" example:"ACTIVE"`                         // 用户状态
		IsAdmin       bool                 `json:"is_admin" example:"false"`                        // 是否管理员
		CreateTime    int64                `json:"create_time" example:"1724914598"`                // 账号创建时间(unix秒时间戳)
		LastLoginTime int64                `json:"last_login_time" example:"1724914598"`            // 最近登录时间(unix秒时间戳)
	}

	// RespUserDetail 用户详情数据结构体
	RespUserDetail RespUserBasic

	// ReqCreateUser 创建用户的请求结构体
	ReqCreateUser struct {
		Phone           string `json:"phone" binding:"required" example:"13800138000"`       // 手机号
		Password        string `json:"password" binding:"required"  example:"123456"`        // 密码
		PasswordConfirm string `json:"password_confirm" binding:"required" example:"123456"` // 确认密码
		Email           string `json:"email" example:"abc@a.com"`                            // 邮箱
		Nickname        string `json:"nickname" example:"Tom"`                               // 昵称
		IsAdmin         bool   `json:"is_admin" example:"true"`                              // 是否属于管理员账号
	}

	// ReqUpdateUser 更新用户的请求结构体
	ReqUpdateUser struct {
		Email     string           `json:"email,omitempty" example:"abc@abc.com"` // 邮箱
		Nickname  string           `json:"nickname,omitempty" example:"Tom"`      // 昵称
		Avatar    string           `json:"avatar,omitempty" example:"a/b/c.jpg"`  // 用户头像URL
		Gender    model.UserGender `json:"gender,omitempty" example:"3"`          // 性别:3-保密,1-男,2-女
		Signature string           `json:"signature,omitempty" example:"haha"`    // 个性化签名
		Status    string           `json:"status,omitempty" example:"FREEZE"`     // 用户状态
	}

	ReqUpdateUserProfile struct {
		UserId   int64  `json:"user_id"`   // 用户ID
		RealName string `json:"real_name"` // 真实姓名
		IDNumber string `json:"id_number"` // 身份证号
	}
)
