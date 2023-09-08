package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/cnpythongo/goal/pkg/config"
)

var jwtSecret = []byte(config.GetConfig().App.Secret)

type Claims struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	jwt.StandardClaims
}

func GenerateToken(phone, password string) (string, time.Time, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(15 * 24 * time.Hour)

	claims := Claims{
		phone,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "goal-app",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, expireTime, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
