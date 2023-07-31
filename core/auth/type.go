package auth

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Identity string `json:"identity"`
	ExpireAt int64  `json:"expire_at"`
	jwt.RegisteredClaims
}

type CustomClaims struct {
	UserID   uint
	Username string
	Identity string
}
