package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"goal-app/pkg/config"
)

var (
	jwtSecret           = []byte(config.GetConfig().App.Secret)
	ContextUserKey      = "GoalUser"
	ContextUserTokenKey = "GoalUserToken"
)

type Claims struct {
	ID   int64  `json:"id"`
	UUID string `json:"uuid"`
	// Phone string `json:"phone"`
	jwt.StandardClaims
}

func GenerateToken(id int64, uid, phone string) (string, time.Time, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(15 * 24 * time.Hour)

	claims := Claims{
		id,
		uid,
		// phone,
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
