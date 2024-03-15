package auth

import (
	"errors"
	"fmt"
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/gin-gonic/gin"
	"goal-app/model"
	"goal-app/pkg/jwt"
	"goal-app/pkg/render"
	"gorm.io/gorm"
)

type IAuthService interface {
	Login(c *gin.Context, payload *ReqAdminAuth) (*RespAdminAuth, int, error)
	Logout(c *gin.Context) error
	UpdateUserLastLogin(uuid string) error
}

type authService struct {
	db *gorm.DB
}

func NewAuthService() IAuthService {
	db := model.GetDB()
	return &authService{
		db: db,
	}
}

// Login 登录
func (s *authService) Login(c *gin.Context, payload *ReqAdminAuth) (*RespAdminAuth, int, error) {
	user, err := model.GetUserByPhone(s.db, payload.Phone)
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
		err = s.UpdateUserLastLogin(user.UUID)
	}()
	return data, render.OK, nil
}

// Logout 退出系统
func (s *authService) Logout(c *gin.Context) error {
	if value, ok := c.Get(jwt.ContextUserKey); ok {
		claims := value.(*jwt.Claims)
		userId := claims.ID
		// todo: 清理会话缓存之类的一些操作
		fmt.Println(userId)
	}
	if token, ok := c.Get(jwt.ContextUserTokenKey); ok {
		// todo: 清理会话缓存之类的一些操作
		fmt.Println(token)
	}
	return nil
}

func (s *authService) UpdateUserLastLogin(uuid string) error {
	return model.UpdateUserLastLoginAt(s.db, uuid)
}
