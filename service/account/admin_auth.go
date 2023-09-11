package account

import (
	"errors"
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/jwt"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/cnpythongo/goal/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IAdminAuthService interface {
	Login(payload *types.ReqAdminAuth) (*types.RespAdminAuth, int)
}

type adminAuthService struct {
	ctx     *gin.Context
	userSvc IUserService
}

func NewAdminAuthService(ctx *gin.Context) IAdminAuthService {
	return &adminAuthService{
		ctx:     ctx,
		userSvc: NewUserService(ctx),
	}
}

func (a *adminAuthService) Login(payload *types.ReqAdminAuth) (*types.RespAdminAuth, int) {
	user, err := a.userSvc.GetUserByPhone(payload.Phone)
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
