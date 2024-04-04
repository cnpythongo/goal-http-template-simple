package imageflix

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goal-app/model"
	"gorm.io/gorm"
)

type IImageFlixCreditService interface {
	GetUserImageFlixCredit(userId int64) (*model.ImageFlixCredit, error)
	UpdateUserImageFlixCredit(userId, value int64) error
}

type imageFlixCreditService struct {
	ctx *gin.Context
	db  *gorm.DB
}

func NewImageFlixCreditService() IImageFlixCreditService {
	db := model.GetDB()
	return &imageFlixCreditService{
		db: db,
	}
}

func (s *imageFlixCreditService) GetUserImageFlixCredit(userId int64) (*model.ImageFlixCredit, error) {
	result, err := model.GetImageFlixCreditByUser(s.db, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.NewImageFlixCredit(), nil
	}
	return result, err
}

func (s *imageFlixCreditService) UpdateUserImageFlixCredit(userId, point int64) error {
	return model.UpdateImageFlixCreditByUser(s.db, userId, point)
}
