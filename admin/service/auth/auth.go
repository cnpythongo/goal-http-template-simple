package auth

import (
	"errors"
	"fmt"
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/admin/service/account"
	"github.com/cnpythongo/goal/admin/types"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/jwt"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IAdminAuthService interface {
	Login(payload *types.ReqAdminAuth) (*types.RespAdminAuth, int, error)
	Logout() error
}

type adminAuthService struct {
	ctx     *gin.Context
	userSvc account.IUserService
}

// Login 登录
func (a *adminAuthService) Login(payload *types.ReqAdminAuth) (*types.RespAdminAuth, int, error) {
	user, err := a.userSvc.GetUserByPhone(payload.Phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.AccountUserNotExistError, err
		}
		return nil, response.AccountQueryUserError, err
	}

	if user.Status == model.FREEZE {
		return nil, response.AccountUserFreezeError, err
	}
	if !user.IsAdmin {
		return nil, response.AuthForbiddenError, err
	}

	if !utils.VerifyPassword(payload.Password, user.Password, user.Salt) {
		return nil, response.AuthError, err
	}

	token, expireTime, err := jwt.GenerateToken(user.ID, user.UUID, user.Phone)
	if err != nil {
		return nil, response.AuthTokenGenerateError, err
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
	return data, response.SuccessCode, nil
}

// Logout 退出系统
func (a *adminAuthService) Logout() error {
	ctx := a.ctx
	value, ok := ctx.Get(jwt.ContextUserKey)
	if ok {
		claims := value.(*jwt.Claims)
		userId := claims.ID
		token, ok2 := ctx.Get(jwt.ContextUserTokenKey)
		if ok2 {
			fmt.Println(userId)
			fmt.Println(token)
		}
	}
	return nil
}

func NewAdminAuthService(ctx *gin.Context) IAdminAuthService {
	return &adminAuthService{
		ctx:     ctx,
		userSvc: account.NewUserService(ctx),
	}
}
