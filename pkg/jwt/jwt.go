package jwt

import (
	"errors"
	"tfserver/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour

type TFClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

//获取token
func GetToken(email string) (string, error) {
	//创建声明
	claims := TFClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    config.JwtIssuer,
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(config.JwtKey)
}

//解析token
func ParseToken(tokenString string) (*TFClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &TFClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return config.JwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*TFClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
