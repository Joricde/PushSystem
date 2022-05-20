package util

import (
	"PushSystem/config"
	"PushSystem/model"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt"
	"io"
	"mime/multipart"
	"strconv"
	"time"
)

type Token struct {
	UserID   uint
	UserName string
	Exp      int64
}

func CountSha256(file *multipart.FileHeader) (string, error) {
	var hash = sha256.New()
	open, e := file.Open()
	if e != nil {
		return "", e
	}
	_, e = io.Copy(hash, open)
	if e != nil {
		return "", e
	}
	b := hash.Sum(nil)
	result := hex.EncodeToString(b)
	e = open.Close()
	return result, e
}

func AddSalt(password string, salt int64) string {
	var hash = sha256.New()
	hash.Write([]byte(password))
	b := hash.Sum([]byte("1"))
	mid := hex.EncodeToString(b) + strconv.FormatInt(salt, 10)
	hash.Write([]byte(mid))
	result := hex.EncodeToString(b)
	return result
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

func ParseUserToken(tokenString string) (*Token, error) {
	claims, e := pareToken(tokenString)
	if e != nil {
		return nil, e
	}
	token := Token{
		UserID:   uint(claims[config.TokenUID].(float64)),
		UserName: claims[config.TokenUsername].(string),
		Exp:      int64(claims[config.TokenEXP].(float64)),
	}
	return &token, e

}

func CreateShareToken(groupID uint) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"groupID": groupID,
	})
	shareToken, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return shareToken, nil
}

func ParseShareToken(shareToken string) (uint, error) {
	claims, err := pareToken(shareToken)
	if err != nil {
		return 0, err
	}
	return uint(claims["groupID"].(float64)), nil

}

func pareToken(token string) (jwt.MapClaims, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
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
		return claims, err
	} else {
		return nil, fmt.Errorf("params token error ")
	}

}
