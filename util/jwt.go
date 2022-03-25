package util

import (
	"PushSystem/config"
	"PushSystem/model"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateToken(user *model.User) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		config.TokenUID: user.ID,
		config.TokenEXP: time.Now().Add(config.ExpTime).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	claim, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if claims, ok := claim.Claims.(jwt.MapClaims); ok && claim.Valid {
		return claims, err
	}
	return nil, err
}
