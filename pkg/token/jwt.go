package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtInfo struct {
	RoleId string
}

func NewJwtInfo(roleId string) *JwtInfo {
	return &JwtInfo{RoleId: roleId}
}

func ParseToken(token string, secret string) (*JwtInfo, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	roleId := claim.Claims.(jwt.MapClaims)["uid"].(string)
	return &JwtInfo{RoleId: roleId}, nil
}

func BuildToken(jwtInfo *JwtInfo, secret string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": jwtInfo.RoleId,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
