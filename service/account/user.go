package account

import (
	"errors"
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/jwt"
	"github.com/cnpythongo/goal/pkg/response"
	"gorm.io/gorm"
)

type IUserService interface {
	AdminLogin(payload *ReqAdminAuth) (*RespAdminAuth, int)
	ApiLogin(phone, password string) (map[string]interface{}, int)

	GetUserList(page, size int, conditions map[string]interface{}) ([]*model.User, int, error)
	GetUserByPhone(phone string) (*model.User, error)
	GetUserByUUID(uuid string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(payload *model.User) (*model.User, error)
	DeleteUserByUUID(uuid string) error
}

type userService struct {
	db *gorm.DB
}

func (s *userService) AdminLogin(payload *ReqAdminAuth) (*RespAdminAuth, int) {
	user, err := model.GetUserByPhone(s.db, payload.Phone)
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
	data := &RespAdminAuth{
		Token: token,
		User: respAdminAuthUser{
			UUID:        user.UUID,
			Nickname:    user.Nickname,
			LastLoginAt: user.LastLoginAt,
		},
	}
	return data, response.SuccessCode
}

func (s *userService) ApiLogin(phone, password string) (map[string]interface{}, int) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) DeleteUserByUUID(uuid string) error {
	//TODO implement me
	panic("implement me")
}

func (s *userService) GetUserByPhone(phone string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) GetUserByUUID(uuid string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) CreateUser(payload *model.User) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) GetUserList(page, size int, conditions map[string]interface{}) ([]*model.User, int, error) {
	_, _, _ = model.GetUserList(s.db, page, size, conditions)
	return nil, 0, nil
}

func NewUserService() IUserService {
	db := model.GetDB()
	return &userService{db: db}
}
