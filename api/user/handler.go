package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"goal-app/pkg/jwt"
	"goal-app/pkg/render"
)

type IUserHandler interface {
	Me(c *gin.Context)
	GetUserInfoByUUID(c *gin.Context)
}

type userHandler struct {
	svc IUserService
}

func NewUserHandler(svc IUserService) IUserHandler {
	return &userHandler{svc: svc}
}

// Me 当前登录用户的信息
// @Tags 用户
// @Summary 获取当前登录用户的信息
// @Description 获取当前登录用户的信息
// @Produce json
// @Success 200 {object} render.RespJsonData{data=RespUserInfo} "code不为0时表示有错误"
// @Failure 500
// @Security APIAuth
// @Router /users/me [get]
func (u *userHandler) Me(c *gin.Context) {
	var ctxUser interface{}
	var ok bool
	if ctxUser, ok = c.Get(jwt.ContextUserKey); !ok {
		render.Json(c, render.AuthLoginRequireError, nil)
		return
	}

	claims := ctxUser.(*jwt.Claims)
	userId := claims.ID

	user, code, err := u.svc.GetUserByID(userId)
	if code != render.OK {
		render.Json(c, code, err)
		return
	}

	result := &RespUserInfo{}
	err = copier.Copy(&result, &user)
	if err != nil {
		render.Json(c, render.DBAttributesCopyError, nil)
		return
	}

	result.Phone = user.PhoneMask()
	render.Json(c, render.OK, result)
}

// GetUserInfoByUUID 获取用户的信息
// @Tags 用户
// @Summary 获取用户的信息
// @Description 通过用户UUID获取用户的信息
// @Produce json
// @Param uuid path string true "用户UUID"
// @Success 200 {object} render.RespJsonData{data=RespUserInfo} "code不为0时表示有错误"
// @Failure 500
// @Security APIAuth
// @Router /users/{uuid} [get]
func (u *userHandler) GetUserInfoByUUID(c *gin.Context) {
	var req ReqGetUserInfo
	if err := c.ShouldBindUri(&req); err != nil {
		render.Json(c, render.ParamsError, nil)
		return
	}

	user, code, err := u.svc.GetUserByUUID(req.UUID)
	if code != render.OK {
		render.Json(c, code, err)
		return
	}

	result := &RespUserInfo{}
	err = copier.Copy(&result, &user)
	if err != nil {
		render.Json(c, render.DBAttributesCopyError, nil)
		return
	}

	result.Phone = user.PhoneMask()
	result.LastLoginAt = nil
	render.Json(c, render.OK, result)
}
