package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type CustomClaims struct {
	UserId   string `json:"user_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

// 生成jwt
func GenerateToken(userId string, userName string) (string, error) {

	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := CustomClaims{
		userId,
		userName,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间戳
			IssuedAt:  nowTime.Unix(),    //当前时间戳
			Issuer:    "goapiGame",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 验证jwt
func ParseToken(tokenString string) (*CustomClaims, error) {

	tokenClaims, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// 更新Token，有问题待查
func RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	// 拿到token基础数据
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil

	})

	// 校验token当前还有效
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		// 修改Claims的过期时间(int64)
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return GenerateToken(claims.UserId, claims.UserName)
	}

	return "", fmt.Errorf("token获取失败:%v", err)
}
