package types

import "github.com/cnpythongo/goal-tools/utils"

type (
	ReqAdminAuth struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	RespAdminAuthUser struct {
		UUID        string           `json:"uuid"`
		Phone       string           `json:"phone"`
		LastLoginAt *utils.LocalTime `json:"last_login_at"`
	}

	RespAdminAuth struct {
		Token      string            `json:"token"`
		ExpireTime string            `json:"expire_time"`
		User       RespAdminAuthUser `json:"user"`
	}
)
