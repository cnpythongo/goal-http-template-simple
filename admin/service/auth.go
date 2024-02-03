package service

import (
	"errors"
	"fmt"
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/gin-gonic/gin"
	"goal-app/admin/types"
	"goal-app/model"
	"goal-app/pkg/jwt"
	"goal-app/pkg/render"
	"gorm.io/gorm"
)

type IAdminAuthService interface {
	Login(payload *types.ReqAdminAuth) (*types.RespAdminAuth, int, error)
	Logout() error
}

type adminAuthService struct {
	ctx     *gin.Context
	userSvc IAccountUserService
}

// Login 登录
func (a *adminAuthService) Login(payload *types.ReqAdminAuth) (*types.RespAdminAuth, int, error) {
	user, err := a.userSvc.GetUserByPhone(payload.Phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, render.DataExistError, err
		}
		return nil, render.QueryError, err
	}

	if user.Status == model.UserStatusFreeze {
		return nil, render.AccountUserFreezeError, err
	}
	if !user.IsAdmin {
		return nil, render.AuthForbiddenError, err
	}

	if !utils.VerifyPassword(payload.Password, user.Password, user.Salt) {
		return nil, render.AuthError, err
	}

	token, expireTime, err := jwt.GenerateToken(user.ID, user.UUID, user.Phone)
	if err != nil {
		return nil, render.AuthTokenGenerateError, err
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
		err = a.userSvc.UpdateUserLastLogin(user.UUID)
	}()
	return data, render.OK, nil
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
		userSvc: NewAccountUserService(),
	}
}
