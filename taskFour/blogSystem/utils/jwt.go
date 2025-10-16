// utils/jwt.go
package utils

import (
	. "blogSystem/config"
	"blogSystem/models"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// JWT 签名密钥（生产环境应从环境变量获取）
var jwtSecret = []byte("lxf")

// Claims 定义 JWT 的声明结构
type Claims struct {
	UserID   uint   `json:"userID"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	UserKey  string `json:"userKey"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT Token
func GenerateToken(user *models.User, userKey string) (*Claims, string, error) {
	var email string
	if user.Email != nil {
		email = *user.Email
	}
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		Phone:    user.Phone,
		Email:    email,
		UserKey:  userKey,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // 24小时有效期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-jwt-demo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	return claims, signedToken, err
}

// ParseToken 解析并验证 JWT Token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func RefreshToken(claims *Claims) error {
	//redis存jwt信息
	key := claims.UserKey + ":" + claims.Username
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claimJson, err := json.Marshal(*claims)
	if err != nil {
		return err
	}
	if err = GetRDB().Set(GetContext(), key, claimJson, 1*time.Hour).Err(); err != nil {
		return err
	}
	return nil
}

func VerifyToken(claims *Claims) error {
	expiresAt := claims.ExpiresAt
	//issuedAt := claims.IssuedAt
	//当token过期时间戳-当前时间戳（也就是token剩余有效时间<=20分钟）时，刷新token
	sub := expiresAt.Sub(time.Now())
	//log.Println("sub:", sub.Seconds())
	if sub.Seconds() <= 1200 {
		err := RefreshToken(claims)
		if err != nil {
			return err
		}
	}
	return nil
}
