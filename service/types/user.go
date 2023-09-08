package types

type (
	ReqGetUserList struct {
		Page             int    `json:"page" form:"page"`
		Size             int    `json:"size" form:"size"`
		LastLoginAtStart string `json:"last_login_at_start" form:"last_login_at_start"`
		LastLoginAtEnd   string `json:"last_login_at_end" form:"last_login_at_end"`
	}

	RespUser struct {
		Phone string `json:"phone"`
	}
)
