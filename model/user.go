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
	Salt      int64  `gorm:"type:text"`
	Phone     int64  `gorm:"index" json:"phone"`
	Email     string `gorm:"type:varchar(64)" json:"email"`
	WechatId  int    `gorm:"index" json:"wechat_id"`
	WechatKey string `gorm:"type:varchar(128)" json:"wechat_key"`
}

func CreateUser(user *User) string {
	newUser := new(User)
	r := ""
	DB.Where("username= ? ", user.Username).First(newUser)
	if user.Username == newUser.Username {
		r = "User already exists"
	} else {
		err := DB.Create(user).Error
		if err != nil {
			r = "create user err: " + err.Error()
			zap.L().Debug(r)
			DB.Rollback()
		}
	}
	zap.L().Debug("create user " + user.Username)
	DB.Commit()
	return r
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
