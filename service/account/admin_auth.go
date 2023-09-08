package account

import (
	"errors"
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/jwt"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/cnpythongo/goal/types"
	"gorm.io/gorm"
)

type IAdminAuthService interface {
	AdminLogin(payload *types.ReqAdminAuth) (*types.RespAdminAuth, int)
}

type adminAuthService struct {
	db      *gorm.DB
	userSvc IUserService
}

func NewAdminAuthService() IAdminAuthService {
	db := model.GetDB()
	return &adminAuthService{
		db:      db,
		userSvc: NewUserService(),
	}
}

func (a *adminAuthService) AdminLogin(payload *types.ReqAdminAuth) (*types.RespAdminAuth, int) {
	user, err := model.GetUserByPhone(a.db, payload.Phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.AccountUserNotExistError
		}
		return nil, response.AccountQueryUserError
	}

	if user.Status == model.FREEZE {
		return nil, response.AccountUserFreezeError
	}
	if !user.IsAdmin {
		return nil, response.AuthForbiddenError
	}

	if !utils.VerifyPassword(payload.Password, user.Password, user.Salt) {
		return nil, response.AuthError
	}

	token, expireTime, err := jwt.GenerateToken(user.ID, user.UUID, user.Phone)
	if err != nil {
		return nil, response.AuthTokenGenerateError
	}
	data := &types.RespAdminAuth{
		Token:      token,
		ExpireTime: expireTime.Format(utils.DateTimeLayout),
		User: types.RespAdminAuthUser{
			UUID:        user.UUID,
			Phone:       user.PhoneMask(),
			LastLoginAt: user.LastLoginAt,
		},
	}
	go func() {
		err = a.userSvc.UpdateUserLastLogin(user.ID)
	}()
	return data, response.SuccessCode
}
