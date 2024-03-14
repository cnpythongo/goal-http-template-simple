package auth

import (
	"github.com/gin-gonic/gin"
	"goal-app/admin/accountuser"
	"goal-app/model"
	"goal-app/pkg/render"
	"gorm.io/gorm"
)

type IAuthService interface {
	Login(c *gin.Context, payload *ReqUserAuth) (*RespUserAuth, int, error)
	Logout(c *gin.Context) error
}

type authService struct {
	db      *gorm.DB
	userSvc accountuser.IUserService
}

// Login 登录
func (a *authService) Login(c *gin.Context, payload *ReqUserAuth) (*RespUserAuth, int, error) {
	return nil, render.OK, nil
}

// Logout 退出系统
func (a *authService) Logout(ctx *gin.Context) error {
	return nil
}

func NewAuthService() IAuthService {
	db := model.GetDB()
	return &authService{
		db: db,
	}
}
