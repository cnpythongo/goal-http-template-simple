package auth

import (
	"github.com/gin-gonic/gin"
)

// @Summary 用户登录
// @Schemes
// @Description 前端应用的用户登录接口
// @Tags 登录
// @Accept json
// @Produce json
// @Success 200 {json} Helloworld
// @Router /account/login [post]
func Login(c *gin.Context) {
	panic("implement me")
}

func Logout(c *gin.Context) {
	panic("implement me")
}
