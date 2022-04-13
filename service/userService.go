package service

import (
	"PushSystem/model"
	"PushSystem/util"
	"fmt"
	"time"
)

type UserService struct {
	Username  string
	Nickname  string
	Phone     int64
	Email     string `json:"email"`
	WechatID  int64  `json:"wechat_id"`
	Password  string
	Salt      int64
	WechatKey string `json:"wechat_key"`
}

func (u UserService) IsUsernameExist(username string) bool {
	newUser := new(model.User)
	newUser = newUser.GetUserByUsername(username)
	if username == newUser.Username {
		return true
	} else {
		return false
	}
}

func (u UserService) IsUserPassword(uid uint, password string) bool {
	newUser := new(model.User)
	userPwd := newUser.GetUserPwdByUserID(uid)
	passwordEscape := util.AddSalt(password, userPwd.Salt)
	if passwordEscape == userPwd.Password {
		return true
	} else {
		return false
	}
}

func (u UserService) CreateUser() bool {
	user := model.User{
		Username: u.Username,
		Nickname: u.Nickname,
		Phone:    u.Phone,
		Email:    u.Email,
		WechatID: u.WechatID,
		UserPwd: model.UserPwd{
			Password:  u.Password,
			Salt:      u.Salt,
			WechatKey: u.WechatKey,
		},
	}
	return user.CreateUser(&user)
}

func (u UserService) SetPassword(uid uint, pwd string) bool {
	userPwd := new(model.UserPwd)
	userPwd.Salt = time.Now().UnixMilli()
	userPwd.Password = util.AddSalt(pwd, userPwd.Salt)
	return model.User{}.UpdateUserPassword(uid, userPwd)
}

func (u UserService) SetWechatKey(uid uint, wechatKey string) bool {
	return model.User{}.UpdateUserWechatKey(uid, wechatKey)
}

func (u UserService) SetUserInfoByID(uid uint) bool {
	user := model.User{
		Username: u.Username,
		Nickname: u.Nickname,
		Phone:    u.Phone,
		Email:    u.Email,
		WechatID: u.WechatID,
	}
	return user.UpdateUserInfo(&user)
}

func (u UserService) GetUserByUsername(username string) *model.User {
	return model.User{}.GetUserByUsername(username)
}

func (u UserService) GetUserByWechatID(wechatID int64) *model.User {
	return model.User{}.GetUserByWechatID(wechatID)

}

func (u UserService) SetRedisUser(user *model.User) bool {
	return model.User{}.SetRedisUser(user)
}

func (u UserService) ToString() string {
	return fmt.Sprintf("%+v", u)
}
