package auth

import (
	"github.com/gin-gonic/gin"
	"goal-app/api/user"
	"goal-app/pkg/jwt"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"goal-app/pkg/utils"
)

type IAuthHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Menus(c *gin.Context)
}

type authHandler struct {
	svc IAuthService
}

func NewAuthHandler(svc IAuthService) IAuthHandler {
	return &authHandler{svc: svc}
}

func GetLoginCtxUser(c *gin.Context) (*jwt.Claims, int) {
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

// Login 登录
// @Tags 认证
// @Summary 登录
// @Description 后台管理系统登录接口
// @Accept json
// @Produce json
// @Param data body ReqAdminAuth true "请求体"
// @Success 200 {object} render.JsonDataResp{data=RespAdminAuth} "code不为0时表示有错误"
// @Failure 500
// @Router /account/login [post]
func (h *authHandler) Login(c *gin.Context) {
	var payload *ReqAdminAuth
	if err := c.ShouldBindJSON(&payload); err != nil {
		render.Json(c, render.PayloadError, err)
		return
	}

	result, code, err := h.svc.Login(c, payload)
	if code != render.OK {
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, result)
}

// Logout 退出
// @Tags 认证
// @Summary 退出
// @Description 退出后台管理系统
// @Description 前端调用该接口，无需关注结果，自行清理掉请求头的 Authorization，页面跳转至首页
// @Description 后端可以执行清理redis缓存, 设置token黑名单等操作
// @Produce json
// @Security AdminAuth
// @Success 200 {object} render.JsonDataResp
// @Router /account/logout [post]
func (h *authHandler) Logout(c *gin.Context) {
	go func() {
		err := h.svc.Logout(c)
		if err != nil {
			log.GetLogger().Error(err)
		}
	}()
	render.Json(c, render.OK, "ok")
}

// Menus 账号可用菜单
// @Tags 认证
// @Summary 账号可用菜单
// @Description 账号可用菜单
// @Accept json
// @Produce json
// @Success 200 {object} render.JsonDataResp{data=[]RespSystemMenuItem} "code不为0时表示有错误"
// @Failure 500
// @Security AdminAuth
// @Router /account/menus [get]
func (h *authHandler) Menus(c *gin.Context) {
	ctxUser, errCode := user.GetLoginCtxUser(c)
	if errCode != render.OK {
		render.Json(c, errCode, nil)
		return
	}

	menus, code, err := h.svc.GetAccountMenus(c, ctxUser.ID)
	if err != nil {
		render.Json(c, code, err)
		return
	}

	result := make([]*RespSystemMenuItem, 0)
	for _, menu := range menus {
		if utils.StrInArrayIndex(menu.Kind, []string{"dir", "menu"}) != -1 {
			result = append(result, menu)
		}
	}
	render.Json(c, render.OK, result)
}
