package util

import (
	"PushSystem/model"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

const expTime = time.Minute * 15

var cxt = context.Background()

func CreateToken(user *model.User) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": user.ID,
		"exp": time.Now().Add(expTime).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	model.RedisDB.Set(cxt, "token"+strconv.Itoa(int(user.ID)), token, expTime*2)
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

func RenewToken(user *model.User) (string, error) {
	return CreateToken(user)
}
