package types

import "github.com/cnpythongo/goal/model"

type (
	ReqAdminAuth struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	RespAdminAuthUser struct {
		UUID        string           `json:"uuid"`
		Nickname    string           `json:"nickname"`
		LastLoginAt *model.LocalTime `json:"last_login_at"`
	}

	RespAdminAuth struct {
		Token string            `json:"token"`
		User  RespAdminAuthUser `json:"user"`
	}
)
