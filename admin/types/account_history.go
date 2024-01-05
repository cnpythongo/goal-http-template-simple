package types

type (
	// ReqGetHistoryList 获取登录历史记录列表的请求参数体
	ReqGetHistoryList struct {
		Pagination
		UserId   int    `json:"user_id" form:"user_id" example:"123"`          // 用户ID
		UserUUID string `json:"user_uuid" form:"user_uuid" example:"abcef123"` // 用户UUID
		Phone    string `json:"phone" form:"phone" example:"13800138000"`      // 用户手机号
		Email    string `json:"email" form:"email" example:"abc@abc.com"`      // 用户邮箱
		Nickname string `json:"nickname" form:"nickname" example:"Tom"`        // 用户昵称
	}
)
