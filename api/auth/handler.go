package auth

import (
	"github.com/gin-gonic/gin"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

type IAuthHandler interface {
	// Signup 注册
	Signup(c *gin.Context)
	// Signin 登录
	Signin(c *gin.Context)
	// Logout 退出
	Logout(c *gin.Context)
	// Captcha 生成验证码
	Captcha(c *gin.Context)
}

type authHandler struct {
	svc IAuthService
}

func NewAuthHandler(svc IAuthService) IAuthHandler {
	return &authHandler{svc: svc}
}

// Signin 登录
// @Tags 登录认证
// @Summary 登录
// @Description 前台用户登录接口
// @Accept json
// @Produce json
// @Param data body ReqUserAuth true "请求体"
// @Success 200 {object} render.RespJsonData{data=RespUserAuth} "code不为0时表示有错误"
// @Failure 500
// @Router /auth/signin [post]
func (h *authHandler) Signin(c *gin.Context) {
	var req *ReqUserAuth
	if err := c.ShouldBindJSON(&req); err != nil {
		render.Json(c, render.PayloadError, err)
		return
	}

	match, err := h.svc.CaptchaVerify(req.CaptchaId, req.CaptchaAnswer)
	if !match {
		render.Json(c, render.AuthCaptchaError, nil)
		return
	}

	result, code, err := h.svc.Login(c, req)
	if err != nil {
		render.Json(c, code, err)
		return
	}
	render.Json(c, render.OK, result)
}

// Logout 退出
// @Tags 登录认证
// @Summary 退出
// @Description 前台用户退出
// @Description 前端调用该接口，无需关注结果，自行清理掉请求头的 Authorization，页面跳转至首页
// @Description 后端可以执行清理redis缓存, 设置token黑名单等操作
// @Produce json
// @Security APIAuth
// @Success 200 {object} render.RespJsonData
// @Failure 500
// @Router /auth/logout [post]
func (h *authHandler) Logout(c *gin.Context) {
	go func() {
		err := h.svc.Logout(c)
		if err != nil {
			log.GetLogger().Error(err)
		}
	}()
	render.Json(c, render.OK, "ok")
}

// Signup 注册
// @Tags 登录认证
// @Summary 注册
// @Description 前台用户注册接口
// @Accept json
// @Produce json
// @Param data body ReqAuthSignup true "请求体"
// @Success 200 {object} render.RespJsonData
// @Failure 500
// @Router /auth/signup [post]
func (h *authHandler) Signup(c *gin.Context) {
	var req ReqAuthSignup
	if err := c.ShouldBindJSON(&req); err != nil {
		render.Json(c, render.PayloadError, err)
		return
	}

	if req.Password != req.ConfirmPassword {
		render.Json(c, render.ParamsError, nil)
		return
	}

	match, err := h.svc.CaptchaVerify(req.CaptchaId, req.CaptchaAnswer)
	if !match {
		render.Json(c, render.AuthCaptchaError, nil)
		return
	}

	code, err := h.svc.Signup(&req)
	if err != nil {
		render.Json(c, code, err)
		return
	}

	render.Json(c, render.OK, nil)
}

// Captcha 获取验证码ID和图片base64
// @Tags 通用
// @Summary 获取验证码ID和图片base64
// @Description 获取验证码ID和图片base64
// @Produce json
// @Param data query ReqAuthCaptcha true "请求参数"
// @Success 200 {object} render.RespJsonData{data=RespAuthCaptcha}
// @Failure 500
// @Router /auth/captcha [get]
func (h *authHandler) Captcha(c *gin.Context) {
	var req ReqAuthCaptcha
	if err := c.ShouldBindQuery(&req); err != nil {
		render.Json(c, render.PayloadError, err)
		return
	}
	resp, code, err := h.svc.Captcha(&req)
	if err != nil {
		render.Json(c, code, err)
		return
	}

	render.Json(c, render.OK, resp)
}
