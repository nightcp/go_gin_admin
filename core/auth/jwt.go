package auth

import (
	"admin/core"
	"admin/util"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-module/carbon/v2"
)

type auth struct {
}

var Jwt = &auth{}

// MakeToken 生成令牌
func (a *auth) MakeToken(data CustomClaims, ttl int) (string, error) {
	c := Claims{
		UserID:   data.UserID,
		Username: data.Username,
		Identity: data.Identity,
		ExpireAt: carbon.Now().AddMinutes(ttl).Timestamp(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    core.Config.AppName,
			Subject:   data.Username,
			Audience:  jwt.ClaimStrings{core.Config.AppName},
			ExpiresAt: jwt.NewNumericDate(carbon.Now().AddMinutes(ttl).ToStdTime()),
			NotBefore: jwt.NewNumericDate(carbon.Now().AddSecond().ToStdTime()),
			IssuedAt:  jwt.NewNumericDate(carbon.Now().ToStdTime()),
			ID:        util.StrUtil.MakeRandomStr(6),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return t.SignedString([]byte(core.Config.AppName))
}

// ParseToken 解析令牌
func (a *auth) ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(core.Config.AppName), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
