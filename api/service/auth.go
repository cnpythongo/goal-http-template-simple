package service

import (
	"github.com/gin-gonic/gin"
)

type IApiAuthService interface {
	Login(phone, password string) (map[string]interface{}, int)
}

type apiAuthService struct {
	ctx *gin.Context
}

func NewApiAuthService(ctx *gin.Context) IApiAuthService {
	return &apiAuthService{ctx: ctx}
}

func (a *apiAuthService) Login(phone, password string) (map[string]interface{}, int) {
	//TODO implement me
	panic("implement me")
}
