package auth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"goal-app/model"
	"goal-app/pkg/jwt"
	"goal-app/pkg/render"
	"goal-app/pkg/utils"
	"gorm.io/gorm"
)

type IAuthService interface {
	Login(c *gin.Context, payload *ReqAdminAuth) (*RespAdminAuth, int, error)
	Logout(c *gin.Context) error
	UpdateUserLastLogin(uuid string) error
	GetAccountMenus(c *gin.Context, userId int64) ([]*RespSystemMenuItem, int, error)
	GetAccountMenuAuthTags(c *gin.Context, userId int64) ([]string, int, error)
}

type authService struct {
}

func NewAuthService() IAuthService {
	return &authService{}
}

// Login 登录
func (s *authService) Login(c *gin.Context, payload *ReqAdminAuth) (*RespAdminAuth, int, error) {
	user, err := model.GetUserByPhone(model.GetDB(), payload.Phone)
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

	result := RespAdminAuthUser{}
	err = copier.Copy(&result, &user)
	if err != nil {
		return nil, render.DBAttributesCopyError, err
	}
	result.Phone = user.PhoneMask()

	data := &RespAdminAuth{
		Token:      token,
		ExpireTime: expireTime.Format(utils.DateTimeLayout),
		User:       result,
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
	return model.UpdateUserLastLoginAt(model.GetDB(), uuid)
}

func (s *authService) GetAccountMenus(c *gin.Context, userId int64) ([]*RespSystemMenuItem, int, error) {
	res := make([]*RespSystemMenuItem, 0)
	roleUsers, _, err := model.GetSystemRoleUserList(
		model.GetDB(), 0, 0, "user_id = ?", []interface{}{userId},
	)
	if err != nil {
		return nil, render.QueryError, err
	}
	if roleUsers == nil || len(roleUsers) == 0 {
		return res, render.OK, nil
	}

	roleIds := make([]int64, 0)
	for _, roleUser := range roleUsers {
		roleIds = append(roleIds, roleUser.RoleID)
	}

	roleMenus, _, err := model.GetSystemRoleMenuList(
		model.GetDB(), 0, 0, "role_id in ?", []interface{}{roleIds},
	)
	if err != nil {
		return nil, render.QueryError, err
	}
	if roleMenus == nil || len(roleMenus) == 0 {
		return res, render.OK, nil
	}

	menuIds := make([]int64, 0)
	for _, rm := range roleMenus {
		menuIds = append(menuIds, rm.MenuID)
	}

	menus, _, err := model.GetSystemMenuList(
		model.GetDB(), 0, 0, "id in ? and status = ?", []interface{}{menuIds, "enable"},
	)
	if err != nil {
		return nil, render.QueryError, err
	}

	err = copier.Copy(&res, &menus)
	if err != nil {
		return nil, render.DBAttributesCopyError, err
	}
	return res, render.OK, nil
}

func (s *authService) GetAccountMenuAuthTags(c *gin.Context, userId int64) ([]string, int, error) {
	menus, _, err := s.GetAccountMenus(c, userId)
	if err != nil {
		return nil, render.QueryError, err
	}
	result := make([]string, 0)
	for _, menu := range menus {
		if menu.AuthTag != "" {
			result = append(result, menu.AuthTag)
		}
	}
	return result, render.OK, err
}
