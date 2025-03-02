package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtSecret = []byte("jwt_secret")

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

// 签发tokens
func GenerateToken(id uint, userName string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		ID:        id,
		UserName:  userName,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Alvin-Mall",
		},
	}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaim.SignedString(JwtSecret)
	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

// 签发Email的token
func GenerateEmailToken(userID, Operation uint, email, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := EmailClaims{
		UserID:        userID,
		Email:         email,
		Password:      password,
		OperationType: Operation,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Alvin-Mall",
		},
	}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaim.SignedString(JwtSecret)
	return token, err
}

func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if tokenClaims != nil {
		if Claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return Claims, nil
		}
	}
	return nil, err
}
