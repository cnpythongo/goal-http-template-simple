package account

import (
	"errors"
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/jwt"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/cnpythongo/goal/service/types"
	"gorm.io/gorm"
)

type IAdminAuthService interface {
	AdminLogin(payload *types.ReqAdminAuth) (*types.RespAdminAuth, int)
}

type adminAuthService struct {
	db *gorm.DB
}

func NewAdminAuthService() IAdminAuthService {
	db := model.GetDB()
	return &adminAuthService{db: db}
}

func (a *adminAuthService) AdminLogin(payload *types.ReqAdminAuth) (*types.RespAdminAuth, int) {
	user, err := model.GetUserByPhone(a.db, payload.Phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.AccountUserNotExistError
		}
		return nil, response.AccountQueryUserError
	}
	if user.Status == model.INACTIVE {
		return nil, response.AccountUserInactiveError
	} else if user.Status == model.FREEZE {
		return nil, response.AccountUserFreezeError
	}

	if !utils.VerifyPassword(payload.Password, user.Password, user.Salt) {
		return nil, response.AuthError
	}

	token, err := jwt.GenerateToken(user.Phone, user.Password)
	if err != nil {
		return nil, response.AuthTokenGenerateError
	}
	data := &types.RespAdminAuth{
		Token: token,
		User: types.RespAdminAuthUser{
			UUID:        user.UUID,
			Nickname:    user.Nickname,
			LastLoginAt: user.LastLoginAt,
		},
	}
	return data, response.SuccessCode
}
