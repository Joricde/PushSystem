package model

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// User Model
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(64);not null;unique;index" json:"username"`
	Nickname string `gorm:"type:varchar(64)" json:"nickname"`
	Phone    int64  `gorm:"index" json:"phone"`
	Email    string `gorm:"type:varchar(64); index" json:"email"`
	WechatID int64  `gorm:"index" json:"wechat_id"`
}

type UserPwd struct {
	gorm.Model
	User      User
	UserID    uint
	Password  string `gorm:"type:varchar(256);not null" json:"password"`
	Salt      int64  `gorm:"type:text"`
	WechatKey string `gorm:"type:varchar(128)" json:"wechat_key"`
}

func (u User) CreateUser(user *User, pwd *UserPwd) bool {
	err := DB.Create(user).Error
	b := true
	if err != nil {
		zap.L().Error("create user err: " + err.Error())
		DB.Rollback()
		b = false
	} else {
		pwd.UserID = u.GetUserByUsername(user.Username).ID
		err = DB.Create(pwd).Error
		if err != nil {
			zap.L().Debug("create user err: " + err.Error())
			DB.Rollback()
			b = false
		}
	}
	zap.L().Debug("create user " + user.Username)
	DB.Commit()
	return b
}

func (u User) GetUserPwdByUserID(uid uint) *UserPwd {
	userPwd := new(UserPwd)
	DB.Where("user_id = ?", uid).First(userPwd)
	return userPwd
}

func (u User) GetUserByUsername(username string) *User {
	user := new(User)
	DB.Where(&User{Username: username}).First(user)
	return user
}

func (u User) GetUserByPhone(phone int64) *User {
	user := new(User)
	DB.Where(&User{Phone: phone}).First(user)
	return user
}

func (u User) GetUserByEmail(email string) *User {
	user := new(User)
	DB.Where(&User{Email: email}).First(user)
	return user
}

func (u User) GetUserByWechatID(wechatId int64) *User {
	user := new(User)
	DB.Where(&User{WechatID: wechatId}).First(user)
	return user
}

func (u User) UpdateUserInfo(user User) bool {
	newUser := new(User)
	e := DB.Model(newUser).Updates(User{
		Username: user.Username,
		Nickname: user.Nickname,
		Phone:    user.Phone,
		Email:    user.Email,
		WechatID: user.WechatID,
	}).Error
	if len(e.Error()) > 0 {
		zap.L().Error(e.Error())
		return false
	} else {
		return true
	}
}

func (u User) UpdateUserPassword(uid uint, pwd UserPwd) bool {
	e := DB.Model(UserPwd{}).Where("user_id = ? ", uid).Updates(UserPwd{
		Password: pwd.Password,
		Salt:     pwd.Salt,
	}).Error
	if len(e.Error()) > 0 {
		zap.L().Error(e.Error())
		return false
	} else {
		return true
	}
}

func (u User) DeleteUserByID(uid uint) bool {
	e := DB.Where("id = ?", uid).Delete(&User{}).Error
	if e != nil {
		zap.L().Error(e.Error())
		return false
	} else {
		return true
	}
}

func (u User) ToString() string {
	return fmt.Sprintf("%+v", u)
}
