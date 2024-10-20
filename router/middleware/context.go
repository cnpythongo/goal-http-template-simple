package middleware

import (
	"github.com/gin-gonic/gin"
	"goal-app/pkg/jwt"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
)

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
