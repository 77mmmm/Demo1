package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("miaomiaomi")

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 将用户名形成token
func GenerateToken(username string) (string, error) {
	expireAt := time.Now().Add(time.Hour).Unix()
	claims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    "somebody",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 解析前端传来的Token
func ParseToken(token string) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil

	})
	if tokenClaims != nil {
		//tokenClaims.Claims本身是*Entity.MyClaims类型，将其断言
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}
