package util

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// CustomerClaims 自定义Claims结构体
type CustomerClaims struct {
	ID       uint
	Username string
	jwt.RegisteredClaims
}

// TODO : 可配置密钥
// 定义密钥
var secretKey = []byte("your-secret-key")

// JWTManager 定义jwt管理器
type JWTManager struct {
	secretKey []byte
	issuer    string
	expire    time.Duration
}

// NewJWTManager 新建 jwt管理器
func NewJWTManager(secretKey, issuer string, expire time.Duration) *JWTManager {
	return &JWTManager{
		secretKey: []byte(secretKey),
		issuer:    issuer,
		expire:    expire,
	}
}

// Generate 生成JWT Token
func Generate(id uint, username string) (string, error) {
	claims := CustomerClaims{
		ID:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// TODO : 可配置过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// Verify 验证token
func ParseToken(tokenString string) (*CustomerClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomerClaims{}, func(token *jwt.Token) (any, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomerClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
