package account

// 查询用户结构体
type ReqGetUserListPayload struct {
	Page             int    `json:"page" form:"page" binding:"required"`
	Size             int    `json:"size" form:"size" binding:"required"`
	LastLoginAtStart string `json:"last_login_at_start" form:"last_login_at_start"`
	LastLoginAtEnd   string `json:"last_login_at_end" form:"last_login_at_end"`
}
