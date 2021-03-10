package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtInfo struct {
	RoleId string
	Msg    string
}

func ParseToken(token string, secret string) (*JwtInfo, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	roleId := claim.Claims.(jwt.MapClaims)["uid"].(string)
	msg := claim.Claims.(jwt.MapClaims)["msg"].(string)
	return &JwtInfo{RoleId: roleId, Msg: msg}, nil
}

func BuildToken(jwtInfo *JwtInfo, secret string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": jwtInfo.RoleId,
		"msg": jwtInfo.Msg,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
