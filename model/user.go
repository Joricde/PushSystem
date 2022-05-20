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
	Nickname string `gorm:"type:varchar(64);default:user" json:"nickname"`
	Phone    int64  `gorm:"index; default:null" json:"phone"`
	Email    string `gorm:"type:varchar(64); index" json:"email"`
	WechatID int64  `gorm:"index; default:null" json:"wechat_id"`
	Password Password
	Groups   []Group `gorm:"many2many:user_groups"`
	Dialogue []Dialogue
}

//func (User) BeforeCreate(tx *gorm.DB) (err error) {
//	err = DB.SetupJoinTable(&User{}, "Group", UserGroup{})
//	if err != nil {
//		zap.L().Error("create join table err " + err.Error())
//		fmt.Println("create join table err " + err.Error())
//	}
//	return err
//}

func (u User) Create(user *User) bool {
	err := DB.Create(user).Error
	if err != nil {
		zap.L().Error("create user err: " + err.Error())
		DB.Rollback()
		return false
	}
	zap.L().Debug("create user " + user.Username)
	return true
}

func (u User) UpdateInfo(user *User) bool {
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

func (u User) RetrievePasswordByUserID(userID uint) error {
	e := DB.Where("id = ? ", userID).First(&u).Error
	return e
}

func (u User) DeleteByUserID(userID uint) bool {
	e := DB.Where("id = ?", userID).Delete(&User{}).Error
	if e != nil {
		zap.L().Error(e.Error())
		return false
	} else {
		return true
	}
}

func (u User) RetrieveByUsername(username string) *User {
	user := new(User)
	DB.Where(&User{Username: username}).First(user)
	return user
}

func (u User) RetrieveByPhone(phone int64) *User {
	user := new(User)
	DB.Where(&User{Phone: phone}).First(user)
	return user
}

func (u User) RetrieveByEmail(email string) *User {
	user := new(User)
	DB.Where(&User{Email: email}).First(user)
	return user
}

func (u User) RetrieveByWechatID(wechatId int64) *User {
	user := new(User)
	DB.Where(&User{WechatID: wechatId}).First(user)
	return user
}

func (u User) ToString() string {
	return fmt.Sprintf("%+v", u)
}

func (g UserGroup) ToString() string {
	return fmt.Sprintf("%+v", g)
}
