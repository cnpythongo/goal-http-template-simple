package user

import (
	"github.com/gin-gonic/gin"
	"goal-app/model"
	"goal-app/pkg/render"
	"gorm.io/gorm"
)

type IUserService interface {
	GetUserByPhone(phone string) (*model.User, error)
	GetUserByUUID(uuid string) (*model.User, int, error)
	GetUserByEmail(email string) (*model.User, int, error)
	GetUserByID(id int64) (*model.User, int, error)
}

type userService struct {
	ctx *gin.Context
	db  *gorm.DB
}

func NewUserService() IUserService {
	db := model.GetDB()
	return &userService{
		db: db,
	}
}

func (s *userService) GetUserByPhone(phone string) (*model.User, error) {
	return model.GetUserByPhone(s.db, phone)
}

func (s *userService) GetUserByUUID(uuid string) (*model.User, int, error) {
	user, err := model.GetUserByConditions(s.db, map[string]interface{}{"uuid": uuid})
	if err != nil {
		return nil, render.QueryError, err
	}
	return user, render.OK, nil
}

func (s *userService) GetUserByEmail(email string) (*model.User, int, error) {
	user, err := model.GetUserByConditions(s.db, map[string]interface{}{"email": email})
	if err != nil {
		return nil, render.QueryError, err
	}
	return user, render.OK, nil
}

func (s *userService) GetUserByID(id int64) (*model.User, int, error) {
	user, err := model.GetUserByConditions(s.db, map[string]interface{}{"id": id})
	if err != nil {
		return nil, render.QueryError, err
	}
	return user, render.OK, nil
}
