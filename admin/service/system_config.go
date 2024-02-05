package service

import (
	"goal-app/model"
	"gorm.io/gorm"
)

type ISystemConfigService interface {
	GetSystemConfigList()
}

type systemConfigService struct {
	db *gorm.DB
}

func NewSystemConfigService() ISystemConfigService {
	db := model.GetDB()
	return &systemConfigService{
		db: db,
	}
}
func (s *systemConfigService) GetSystemConfigList() {
	return
}
