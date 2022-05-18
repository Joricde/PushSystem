package util

import (
	"PushSystem/config"
	"PushSystem/model"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type Token struct {
	UserID   uint
	UserName string
	Exp      int64
}

func CreateToken(user *model.User) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		config.TokenUID:      user.ID,
		config.TokenUsername: user.Username,
		config.TokenEXP:      time.Now().Add(config.ExpTime).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(tokenString string) (*Token, error) {
	claim, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := claim.Claims.(jwt.MapClaims)
	if ok && claim.Valid {
		token := Token{
			UserID:   uint(claims[config.TokenUID].(float64)),
			UserName: claims[config.TokenUsername].(string),
			Exp:      int64(claims[config.TokenEXP].(float64)),
		}
		return &token, err
	} else {
		return nil, fmt.Errorf("params token error ")
	}

}
