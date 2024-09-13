package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"goal-app/pkg/jwt"
	"goal-app/pkg/render"
)

func JWTAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var err error
		var claims *jwt.Claims

		code = render.OK
		token := c.GetHeader("Authorization")

		if token == "" {
			code = render.AuthLoginRequireError
		} else {
			token = strings.TrimSpace(strings.Replace(token, "Bearer", "", 1))
			claims, err = jwt.ParseToken(token)
			if err != nil {
				code = render.AuthTokenError
			} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
				code = render.AuthTokenError
			}
		}

		if code != render.OK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  render.GetCodeMsg(code, nil),
			})

			c.Abort()
			return
		}

		c.Set(jwt.ContextUserKey, claims)
		c.Set(jwt.ContextUserTokenKey, token)
		c.Next()
	}
}
