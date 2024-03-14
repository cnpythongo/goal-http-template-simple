package user

import "github.com/gin-gonic/gin"

type IUserHandler interface {
	Me(c *gin.Context)
	GetUserByUUID(c *gin.Context)
}

type userHandler struct {
	svc IUserService
}

func NewUserHandler(svc IUserService) IUserHandler {
	return &userHandler{svc: svc}
}

func (u *userHandler) Me(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *userHandler) GetUserByUUID(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
