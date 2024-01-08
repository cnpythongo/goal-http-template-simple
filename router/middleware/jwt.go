package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/cnpythongo/goal/pkg/jwt"
	"github.com/cnpythongo/goal/pkg/response"
)

func JWTAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var err error
		var claims *jwt.Claims

		code = response.SuccessCode
		token := c.GetHeader("Authorization")

		if token == "" {
			code = response.AuthLoginRequireError
		} else {
			token = strings.TrimSpace(strings.Replace(token, "Bearer", "", 1))
			claims, err = jwt.ParseToken(token)
			if err != nil {
				code = response.AuthTokenError
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = response.AuthTokenError
			}
		}

		if code != response.SuccessCode {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  response.GetCodeMsg(code),
			})

			c.Abort()
			return
		}

		c.Set(jwt.ContextUserKey, claims)
		c.Set(jwt.ContextUserTokenKey, token)
		c.Next()
	}
}
