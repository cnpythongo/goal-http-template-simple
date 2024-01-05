package types

type (
	// ReqGetUserProfileList 获取用户资料列表参数请求体
	ReqGetUserProfileList struct {
		Pagination
		UserId   int    `json:"user_id" form:"user_id" example:"123"`             // 用户ID
		UserUUID string `json:"user_uuid" form:"user_uuid" example:"abcef123"`    // 用户UUID
		Phone    string `json:"phone" form:"phone" example:"13800138000"`         // 用户手机号
		RealName string `json:"real_name" form:"real_name" example:"张三"`          // 真实姓名，模糊查询
		IDNumber string `json:"id_number" form:"id_number" example:"12312312312"` // 身份证号，模糊查询
	}

	// ReqCreateUserProfile 创建用户资料请求体
	ReqCreateUserProfile struct {
		UserId   int    `json:"user_id" binding:"required" example:"123"`           // 用户ID
		RealName string `json:"real_name" binding:"required" example:"张三"`          // 真实姓名
		IDNumber string `json:"id_number" binding:"required" example:"12312312312"` // 身份证号
	}

	// ReqUpdateUserProfile 更新用户资料请求体
	ReqUpdateUserProfile struct {
		UserId   int    `json:"user_id" binding:"required" example:"123"`  // 用户ID
		RealName string `json:"real_name,omitempty" example:"张三"`          // 真实姓名
		IDNumber string `json:"id_number,omitempty" example:"12312312312"` // 身份证号
	}
)
