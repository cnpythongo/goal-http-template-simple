package account

import (
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IUserService interface {
	GetUserByPhone(phone string) (*model.User, error)
	GetUserByUUID(uuid string) (*model.User, int, error)
}

type userService struct {
	ctx *gin.Context
	db  *gorm.DB
}

func (s *userService) GetUserByPhone(phone string) (*model.User, error) {
	return model.GetUserByPhone(s.db, phone)
}

func (s *userService) GetUserByUUID(uuid string) (*model.User, int, error) {
	user, err := model.GetUserByConditions(s.db, map[string]interface{}{"uuid": uuid})
	if err != nil {
		return nil, response.DBQueryError, err
	}
	return user, response.SuccessCode, nil
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserService(ctx *gin.Context) IUserService {
	db := model.GetDB()
	return &userService{
		ctx: ctx,
		db:  db,
	}
}
