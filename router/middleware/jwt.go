package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/cnpythongo/goal/pkg/jwt"
	"github.com/cnpythongo/goal/pkg/response"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = response.SuccessCode
		token := c.Query("token")
		if token == "" {
			code = response.ParamsError
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = response.AuthTokenError
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = response.AuthTokenTimeoutError
			}
		}

		if code != response.SuccessCode {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  response.GetCodeMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
