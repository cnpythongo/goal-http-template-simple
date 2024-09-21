package auth

type (
	// ReqAdminAuth 用户登录请求结构体
	ReqAdminAuth struct {
		Phone    string `json:"phone" binding:"required" example:"13800138000"` // 手机号
		Password string `json:"password" binding:"required" example:"123456"`   // 密码
	}

	// RespAdminAuthUser 用户基本数据结构
	RespAdminAuthUser struct {
		UUID          string `json:"uuid"`                                 // 用户uuid
		Phone         string `json:"phone" example:"138****8000"`          // 带掩码的手机号
		LastLoginTime int64  `json:"last_login_time" example:"1724914598"` // 最近的登录时间(unix秒时间戳)
		Nickname      string `json:"nickname"`                             // 昵称
		Avatar        string `json:"avatar"`                               // 头像
	}

	// RespAdminAuth 用户登录接口返回数据结构
	RespAdminAuth struct {
		Token      string            `json:"token"`       // 令牌
		ExpireTime string            `json:"expire_time"` // 过期时间
		User       RespAdminAuthUser `json:"user"`        // 用户基本信息
	}
)
