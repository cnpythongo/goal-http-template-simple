package auth

import (
	"errors"
	"fmt"
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/gin-gonic/gin"
	"goal-app/admin/accountuser"
	"goal-app/model"
	"goal-app/pkg/jwt"
	"goal-app/pkg/render"
	"gorm.io/gorm"
)

type IAuthService interface {
	Login(c *gin.Context, payload *ReqAdminAuth) (*RespAdminAuth, int, error)
	Logout(c *gin.Context) error
}

type authService struct {
	db      *gorm.DB
	userSvc accountuser.IUserService
}

// Login 登录
func (a *authService) Login(c *gin.Context, payload *ReqAdminAuth) (*RespAdminAuth, int, error) {
	user, err := a.userSvc.GetUserByPhone(payload.Phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, render.AccountUserOrPwdError, err
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
	data := &RespAdminAuth{
		Token:      token,
		ExpireTime: expireTime.Format(utils.DateTimeLayout),
		User: RespAdminAuthUser{
			UUID:        user.UUID,
			Phone:       user.PhoneMask(),
			LastLoginAt: user.LastLoginAt,
			Nickname:    user.Nickname,
			Avatar:      user.Avatar,
		},
	}
	go func() {
		err = a.userSvc.UpdateUserLastLogin(user.UUID)
	}()
	return data, render.OK, nil
}

// Logout 退出系统
func (a *authService) Logout(ctx *gin.Context) error {
	if value, ok := ctx.Get(jwt.ContextUserKey); ok {
		claims := value.(*jwt.Claims)
		userId := claims.ID
		// todo: 清理会话缓存之类的一些操作
		fmt.Println(userId)
	}
	if token, ok := ctx.Get(jwt.ContextUserTokenKey); ok {
		// todo: 清理会话缓存之类的一些操作
		fmt.Println(token)
	}
	return nil
}

func NewAuthService() IAuthService {
	db := model.GetDB()
	return &authService{
		db:      db,
		userSvc: accountuser.NewUserService(),
	}
}
