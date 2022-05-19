package model

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Password struct {
	gorm.Model
	UserID       uint   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PasswordHash string `gorm:"type:varchar(256);not null"`
	Salt         int64  `gorm:"type:bigint"`
	WechatKey    string `gorm:"type:varchar(128)" json:"wechat_key"`
}

func (u User) UpdatePassword(userID uint, pwd *Password) bool {
	e := DB.Model(Password{}).Where("user_id = ? ", userID).Updates(Password{
		PasswordHash: pwd.PasswordHash,
		Salt:         pwd.Salt,
	}).Error
	if e != nil {
		zap.L().Error(e.Error())
		return false
	} else {
		return true
	}
}

func (u User) UpdateWechatKey(userID uint, wechatKey string) bool {
	e := DB.Model(Password{}).Where("user_id = ? ", userID).Updates(Password{
		WechatKey: wechatKey,
	}).Error
	if e != nil {
		zap.L().Error(e.Error())
		return false
	} else {
		return true
	}
}

func (u User) GetPasswordByUserID(userID uint) *Password {
	password := new(Password)
	DB.Where("user_id = ?", userID).First(password)
	return password
}

func (p Password) ToString() string {
	return fmt.Sprintf("%+v", p)
}
