package account

import (
	"github.com/cnpythongo/goal/model"
	"gorm.io/gorm"
)

type IUserService interface {
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
