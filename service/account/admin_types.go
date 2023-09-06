package account

import "github.com/cnpythongo/goal/model"

type (
	ReqAdminAuth struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	respAdminAuthUser struct {
		UUID        string           `json:"uuid"`
		Nickname    string           `json:"nickname"`
		LastLoginAt *model.LocalTime `json:"last_login_at"`
	}

	RespAdminAuth struct {
		Token string            `json:"token"`
		User  respAdminAuthUser `json:"user"`
	}

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
