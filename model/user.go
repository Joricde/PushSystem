package model

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// User Model
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);not null;unique;index"`
	nickname string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(256);not null"`
	Phone    int64  `gorm:"index"`
	Email    string `gorm:"varchar(64)"`
}

func CreateUser(user *User) (*User, error) {
	u := new(User)
	DB.Select("username").First(&u)
	if user.Username != u.Username {
		err := DB.Create(user).Error
		zap.L().Debug("create username")
		if err != nil {
			return user, err
		}
		return user, nil
	}
	return new(User), nil
}

func CheckUsername(usr string) (username *User, err error) {
	zap.L().Debug("check username by username")
	DB.Where(&User{Username: usr}).First(&username)
	zap.L().Debug(fmt.Sprintln(username))
	if err != nil {
		return username, err
	}
	return username, nil
}
