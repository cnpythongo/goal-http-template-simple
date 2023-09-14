package types

import (
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/model"
)

type (
	// ReqGetUserList 获取用户列表的请求参数体
	ReqGetUserList struct {
		Page             int                    `json:"page,default=1" form:"page,default=1" default:"1" example:"1"`                 // 页码
		Size             int                    `json:"size,default=10" form:"size,default=10" default:"10" example:"10"`             // 每页数量
		Phone            string                 `json:"phone" form:"phone" example:"13800138000"`                                     // 手机号,模糊查询
		Email            string                 `json:"email" form:"email" example:"abc@abc.com"`                                     // 邮箱,模糊查旬
		Nickname         string                 `json:"nickname" form:"nickname" example:"Tom"`                                       // 昵称,模糊查询
		Status           []model.UserStatusType `json:"status" form:"status" example:"FREEZE,ACTIVE"`                                 // 用户状态,多种状态过滤使用逗号分隔
		LastLoginAtStart string                 `json:"last_login_at_start" form:"last_login_at_start" example:"2023-09-01 01:30:59"` // 最近登录时间起始
		LastLoginAtEnd   string                 `json:"last_login_at_end" form:"last_login_at_end" example:"2023-09-01 22:59:59"`     // 最近登录时间截止
	}

	// RespUserBasic 用户基础数据结构体
	RespUserBasic struct {
		UUID        string               `json:"uuid" example:"826d6b1aa64d471d822d667e92218158"` // 用户UUID,32位字符串
		Phone       string               `json:"phone" example:"13800138000"`                     // 手机号
		Nickname    string               `json:"nickname" example:"Tom"`                          // 昵称
		LastLoginAt *utils.LocalTime     `json:"last_login_at" example:"2023-09-01 13:30:59"`     // 最近登录时间
		Status      model.UserStatusType `json:"status" example:"ACTIVE"`                         // 用户状态
		IsAdmin     bool                 `json:"is_admin" example:"false"`                        // 是否管理员
	}

	// RespUserDetail 用户详情数据结构体
	RespUserDetail struct {
		RespUserBasic
	}

	// RespGetUserList 获取用户列表的响应数据结构
	RespGetUserList struct {
		Page   int              `json:"page"`   // 当前页
		Total  int              `json:"total"`  // 总页数
		Count  int              `json:"count"`  // 总记录数
		Result []*RespUserBasic `json:"result"` // 当前结果集
	}

	// ReqCreateUser 创建用户的请求结构体
	ReqCreateUser struct {
		Phone           string `json:"phone" binding:"required" example:"13800138000"` // 手机号
		Email           string `json:"email" example:"abc@a.com"`                      // 邮箱
		Nickname        string `json:"nickname" example:"Tom"`                         // 昵称
		Password        string `json:"password" example:"123456"`                      // 密码
		PasswordConfirm string `json:"password_confirm" example:"123456"`              // 确认密码
	}

	// ReqUpdateUser 更新用户的请求结构体
	ReqUpdateUser struct {
		Email     string `json:"email,omitempty" example:"abc@abc.com"` // 邮箱
		Nickname  string `json:"nickname,omitempty" example:"Tom"`      // 昵称
		Avatar    string `json:"avatar,omitempty" example:"a/b/c.jpg"`  // 用户头像URL
		Gender    int64  `json:"gender,omitempty" example:"3"`          // 性别:3-保密,1-男,2-女
		Signature string `json:"signature,omitempty" example:"haha"`    // 个性化签名
		Status    string `json:"status,omitempty" example:"FREEZE"`     // 用户状态
	}
)
