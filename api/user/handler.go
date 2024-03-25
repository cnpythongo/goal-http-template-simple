package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"goal-app/pkg/jwt"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type IUserHandler interface {
	Me(c *gin.Context)
	GetUserInfoByUUID(c *gin.Context)
	Profile(c *gin.Context)
	UpdateProfile(c *gin.Context)
}

type userHandler struct {
	svc IUserService
}

func NewUserHandler(svc IUserService) IUserHandler {
	return &userHandler{svc: svc}
}

func (h *userHandler) getLoginCtxUser(c *gin.Context) (*jwt.Claims, int) {
	var user *jwt.Claims
	var code int

	defer func() {
		if err := recover(); err != nil {
			log.GetLogger().Error(err)
			user = nil
			code = render.Error
		}
	}()

	ctxUser, ok := c.Get(jwt.ContextUserKey)
	if !ok {
		return nil, render.AuthLoginRequireError
	}
	user = ctxUser.(*jwt.Claims)
	code = render.OK
	return user, code
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
func (h *userHandler) Me(c *gin.Context) {
	ctxUser, errCode := h.getLoginCtxUser(c)
	if errCode != render.OK {
		render.Json(c, errCode, nil)
		return
	}

	user, code, err := h.svc.GetUserByID(ctxUser.ID)
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
func (h *userHandler) GetUserInfoByUUID(c *gin.Context) {
	var req ReqGetUserInfo
	if err := c.ShouldBindUri(&req); err != nil {
		render.Json(c, render.ParamsError, nil)
		return
	}

	user, code, err := h.svc.GetUserByUUID(req.UUID)
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

// Profile 用户个人资料
// @Tags 用户
// @Summary 用户个人资料
// @Description 当前登录用户的个人资料
// @Produce json
// @Success 200 {object} render.RespJsonData
// @Failure 400 {object} render.RespJsonData
// @Security APIAuth
// @Router /users/me/profile [get]
func (h *userHandler) Profile(c *gin.Context) {
	var ctxUser interface{}
	var ok bool
	if ctxUser, ok = c.Get(jwt.ContextUserKey); !ok {
		render.Json(c, render.AuthLoginRequireError, nil)
		return
	}
	user := ctxUser.(*jwt.Claims)
	pf, code, err := h.svc.GetUserProfile(user.ID)
	if err != nil {
		render.Json(c, code, err)
		return
	}

	render.Json(c, render.OK, pf)
}

// UpdateProfile 更新用户个人资料
// @Tags 用户
// @Summary 更新用户个人资料
// @Description 更新当前登录用户的个人资料
// @Accept json
// @Produce json
// @Success 200 {object} render.RespJsonData
// @Failure 400 {object} render.RespJsonData
// @Security APIAuth
// @Router /users/me/profile/update [post]
func (h *userHandler) UpdateProfile(c *gin.Context) {
	var req ReqUpdateUserProfile
	if err := c.ShouldBindJSON(&req); err != nil {
		render.Json(c, render.ParamsError, err)
		return
	}

	user, errCode := h.getLoginCtxUser(c)
	if errCode != render.OK {
		render.Json(c, errCode, nil)
		return
	}

	req.UserId = user.ID
	code, err := h.svc.UpdateUserProfile(&req)
	if err != nil {
		render.Json(c, code, err)
		return
	}

	render.Json(c, render.OK, "ok")
}
