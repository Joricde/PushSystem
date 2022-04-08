package service

import (
	"PushSystem/model"
	"PushSystem/util"
	"time"
)

type UserService struct{}

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

func (u UserService) CreateUser(user *model.User, pwd *model.UserPwd) bool {
	return model.User{}.CreateUser(user, pwd)
}

func (u UserService) SetPassword(uid uint, pwd string) bool {
	userPwd := new(model.UserPwd)
	userPwd.Salt = time.Now().UnixMilli()
	userPwd.Password = util.AddSalt(userPwd.Password, userPwd.Salt)
	return model.User{}.UpdateUserPassword(uid, userPwd)
}

func (u UserService) SetWechatKey(uid uint, wechatKey string) bool {
	return model.User{}.UpdateUserWechatKey(uid, wechatKey)
}

func (u UserService) SetUserInfo(user *model.User) bool {
	return model.User{}.UpdateUserInfo(user)
}

func (u UserService) GetUserByUsername(username string) *model.User {
	return model.User{}.GetUserByUsername(username)
}

func (u UserService) SetRedisUser(user *model.User) (string, error) {
	return model.User{}.SetRedisUser(user)
}
