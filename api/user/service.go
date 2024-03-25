package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goal-app/model"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"gorm.io/gorm"
)

type IUserService interface {
	GetUserByPhone(phone string) (*model.User, error)
	GetUserByUUID(uuid string) (*model.User, int, error)
	GetUserByEmail(email string) (*model.User, int, error)
	GetUserByID(id int64) (*model.User, int, error)
	GetUserProfile(userId int64) (*model.UserProfile, int, error)
	UpdateUserProfile(payload *ReqUpdateUserProfile) (int, error)
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

func (s *userService) GetUserProfile(userId int64) (*model.UserProfile, int, error) {
	pf, err := model.GetUserProfileByUserId(s.db, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, render.DataNotExistError, err
		}
		log.GetLogger().Errorf("model.GetUserProfileByUserId Error ==> %v", err)
		return nil, render.QueryError, err
	}
	return pf, render.OK, nil
}

func (s *userService) UpdateUserProfile(payload *ReqUpdateUserProfile) (int, error) {
	pf, code, err := s.GetUserProfile(payload.UserId)
	if err != nil {
		return code, err
	}

	_, err = model.UpdateUserProfile(s.db, pf)
	if err != nil {
		log.GetLogger().Errorf("model.UpdateUserProfile Error ==> %v", err)
		return render.UpdateError, err
	}
	return render.OK, nil
}
