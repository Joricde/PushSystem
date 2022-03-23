package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateToken(uid, secret string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(tokenString string, secret string) (jwt.MapClaims, error) {
	claim, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if claims, ok := claim.Claims.(jwt.MapClaims); ok && claim.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
		return claims, err
	}
	return nil, err
}
