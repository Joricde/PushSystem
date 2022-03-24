package model

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// User Model
type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(50);not null;unique;index" json:"username"`
	Nickname  string `gorm:"type:varchar(50)" json:"nickname"`
	Password  string `gorm:"type:varchar(256);not null" json:"password"`
	Salt      int64
	Phone     int64  `gorm:"index"`
	Email     string `gorm:"type:varchar(64)"`
	WechatId  int    `gorm:"index" `
	WechatKey string `gorm:"type:varchar(128)"`
}

func CreateUser(user *User) string {
	newUser := new(User)
	DB.Where("username= ? ", user.Username).First(newUser)
	if user.Username == newUser.Username {
		return ""
	}
	err := DB.Create(user).Error
	if err != nil {
		zap.L().Debug("create user err: " + err.Error())
		DB.Rollback()
	}
	zap.L().Debug("create user : " + err.Error())
	DB.Commit()
	return err.Error()
}

func GetUserByUsername(username string) *User {
	user := new(User)
	DB.Where(&User{Username: username}).First(user)
	zap.L().Debug(fmt.Sprintln(user))
	return user
}

func (u User) ToString() string {
	return fmt.Sprintf("%+v", u)
}
