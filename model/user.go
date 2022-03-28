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
	Email     string `gorm:"type:varchar(64) unique;index" json:"email"`
	WechatID  int64  `gorm:"index" json:"wechat_id"`
	WechatKey string `gorm:"type:varchar(128)" json:"wechat_key"`
}

func (u User) CreateUser(user *User) bool {
	b := true
	err := DB.Create(user).Error
	if err != nil {
		zap.L().Debug("create user err: " + err.Error())
		DB.Rollback()
		b = false
	}
	zap.L().Debug("create user " + user.Username)
	DB.Commit()
	return b
}

func (u User) GetUserByUsername(username string) *User {
	user := new(User)
	DB.Where(&User{Username: username}).First(user)
	zap.L().Debug(fmt.Sprintln(user))
	return user
}

func (u User) GetUserByPhone(phone int64) *User {
	user := new(User)
	DB.Where(&User{Phone: phone}).First(user)
	zap.L().Debug(fmt.Sprintln(user))
	return user
}

func (u User) GetUserByEmail(email string) *User {
	user := new(User)
	DB.Where(&User{Email: email}).First(user)
	zap.L().Debug(fmt.Sprintln(user))
	return user
}

func (u User) GetUserByWechatID(wechatId int64) *User {
	user := new(User)
	DB.Where(&User{WechatID: wechatId}).First(user)
	zap.L().Debug(fmt.Sprintln(user))
	return user
}

func (u User) UpdateUserInfo(user User) bool {
	newUser := new(User)
	e := DB.Model(newUser).Updates(User{
		Username:  user.Username,
		Nickname:  user.Nickname,
		Phone:     user.Phone,
		Email:     user.Email,
		WechatID:  user.WechatID,
		WechatKey: user.WechatKey,
	}).Error
	zap.L().Debug(fmt.Sprintln(user))
	if len(e.Error()) > 0 {
		zap.L().Error(e.Error())
		return false
	} else {
		return true
	}
}

func (u User) UpdateUserPassword(user User) bool {
	newUser := new(User)
	e := DB.Model(newUser).Updates(User{
		Password: user.Password,
		Salt:     user.Salt,
	}).Error
	zap.L().Debug(fmt.Sprintln(user))
	if len(e.Error()) > 0 {
		zap.L().Error(e.Error())
		return false
	} else {
		return true
	}
}

func (u User) DeleteUserByID(uid int64) bool {
	e := DB.Where("id = ?", uid).Delete(&User{}).Error
	zap.L().Debug("delete user, ok")
	if e != nil {
		zap.L().Debug(e.Error())
		return false
	} else {
		return true
	}
}

func (u User) ToString() string {
	return fmt.Sprintf("%+v", u)
}
