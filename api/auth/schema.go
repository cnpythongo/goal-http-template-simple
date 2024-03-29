package auth

import "github.com/cnpythongo/goal-tools/utils"

type (
	CaptchaStruct struct {
		CaptchaId     string `json:"captcha_id" form:"captcha_id" binding:"required"`                     // 验证码ID
		CaptchaAnswer string `json:"captcha_answer" form:"captcha_answer" binding:"required,min=6,max=6"` // 验证码,4位
	}

	// ReqUserAuth 用户登录接口请求数据结构
	ReqUserAuth struct {
		// Phone    string `json:"phone" binding:"required" example:"13800138000"` // 手机号
		Email    string `json:"email" binding:"required" example:"foo@bar.com"` // 邮箱
		Password string `json:"password" binding:"required" example:"123456"`   // 密码
		CaptchaStruct
	}

	// RespUserInfo 用户登录接口返回的用户信息数据结构
	RespUserInfo struct {
		UUID        string           `json:"uuid"`                        // 用户uuid
		Email       string           `json:"email" example:"foo@bar.com"` // 邮箱
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

	// ReqAuthSignup 用户注册接口数据结构
	ReqAuthSignup struct {
		// Phone    string `json:"phone" binding:"required" example:"13800138000"` // 手机号
		Email           string `json:"email" example:"foo@bar.com"`                          // 邮箱
		Password        string `json:"password" binding:"required" example:"123456"`         // 密码
		ConfirmPassword string `json:"confirm_password" binding:"required" example:"123456"` // 确认密码
		CaptchaStruct
	}

	// ReqAuthCaptcha 验证码接口请求参数
	ReqAuthCaptcha struct {
		TS string `form:"ts"` // 时间戳字符串，避免缓存
	}

	// RespAuthCaptcha 验证码接口返回的数据结构
	RespAuthCaptcha struct {
		CaptchaId  string `json:"captcha_id"`  // 验证码ID
		CaptchaImg string `json:"captcha_img"` // base64编码的验证码图片
	}
)
