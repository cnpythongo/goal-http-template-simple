package auth

import (
	"errors"
	"fmt"
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"goal-app/model"
	"goal-app/pkg/jwt"
	"goal-app/pkg/render"
	"gorm.io/gorm"
)

type IAuthService interface {
	Login(c *gin.Context, payload *ReqUserAuth) (*RespUserAuth, int, error)
	Logout(c *gin.Context) error
	Signup(payload *ReqUserSignup) (int, error)
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
func (s *authService) Login(c *gin.Context, payload *ReqUserAuth) (*RespUserAuth, int, error) {
	user, err := model.GetUserByEmail(s.db, payload.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, render.AccountUserOrPwdError, err
		}
		return nil, render.QueryError, err
	}

	if user.Status == model.UserStatusFreeze {
		return nil, render.AccountUserFreezeError, err
	}

	if !utils.VerifyPassword(payload.Password, user.Password, user.Salt) {
		return nil, render.AuthError, err
	}

	token, expireTime, err := jwt.GenerateToken(user.ID, user.UUID, "")
	if err != nil {
		return nil, render.AuthTokenGenerateError, err
	}
	result := RespUserInfo{}
	err = copier.Copy(&result, &user)
	if err != nil {
		return nil, render.DBAttributesCopyError, err
	}

	data := &RespUserAuth{
		Token:      token,
		ExpireTime: expireTime.Format(utils.DateTimeLayout),
		User:       result,
	}
	go func() {
		err = model.UpdateUserLastLoginAt(s.db, user.UUID)
	}()
	return data, render.OK, nil
}

// Logout 退出
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

func (s *authService) Signup(payload *ReqUserSignup) (int, error) {
	user, err := model.GetUserByEmail(s.db, payload.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return render.QueryError, err
	}

	if user != nil {
		return render.DataExistError, err
	}

	_, err = model.CreateUser(s.db, &model.User{
		UUID:     utils.UUID(),
		Password: payload.Password,
		Email:    payload.Email,
		Status:   model.UserStatusInactive,
		IsAdmin:  false,
		Gender:   model.UserGenderUnknown,
	})
	if err != nil {
		return render.CreateError, err
	}

	// todo: 如果是邮箱注册，还需要发送激活邮件
	return render.OK, nil
}
