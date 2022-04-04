package service

import (
	"PushSystem/model"
	"PushSystem/util"
)

type UserService struct {
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

func (u UserService) IsCreateUser(user *model.User, pwd *model.UserPwd) bool {
	return model.User{}.CreateUser(user, pwd)
}

func (u UserService) SetRedisUser(user *model.User) (string, error) {
	return model.User{}.SetRedisUser(user)
}

func (u UserService) GetUserByUsername(username string) *model.User {
	return model.User{}.GetUserByUsername(username)
}

func (u UserService) GetUserPwdByUid(uid uint) *model.UserPwd {
	return model.User{}.GetUserPwdByUserID(uid)
}
