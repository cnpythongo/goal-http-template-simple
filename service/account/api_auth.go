package account

import (
	"github.com/cnpythongo/goal/model"
	"gorm.io/gorm"
)

type IApiAuthService interface {
	Login(phone, password string) (map[string]interface{}, int)
}

type apiAuthService struct {
	db *gorm.DB
}

func NewApiAuthService() IApiAuthService {
	db := model.GetDB()
	return &apiAuthService{db: db}
}

func (a *apiAuthService) Login(phone, password string) (map[string]interface{}, int) {
	//TODO implement me
	panic("implement me")
}
