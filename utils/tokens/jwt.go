package tokens

import (
	"errors"
	"github.com/Godvictory/douyin/internal/conf"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte(conf.Conf.JwtSecret)

type MyClaims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GetToken 生成token
func GetToken(id int64, username string) (string, error) {
	expireTime := time.Now().Add(time.Hour * 24 * 90) // 三个月过期
	SetClaims := MyClaims{
		id,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "ByteHunters",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	return reqClaim.SignedString(JwtKey)
}

// CheckToken 验证token
func CheckToken(token string) (*MyClaims, error) {
	key, err := jwt.ParseWithClaims(token, &MyClaims{}, func(*jwt.Token) (any, error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := key.Claims.(*MyClaims); ok && key.Valid {
		return claims, nil
	} else {
		return nil, errors.New("你的token已过期")
	}
}
