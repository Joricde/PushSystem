package service

import (
	"PushSystem/model"
	"PushSystem/util"
)

type UserService struct {
}

func (u UserService) IsUsernameExist(username string) bool {
	newUser := new(model.User)
	newUser.GetUserByUsername(username)
	if username == newUser.Username {
		return true
	} else {
		return false
	}
}

func (u UserService) IsUserPassword(username string, password string) bool {
	newUser := new(model.User)
	user := newUser.GetUserByUsername(username)
	passwordEscape := util.AddSalt(password, user.Salt)
	if passwordEscape == user.Password {
		return true
	} else {
		return false
	}
}

func (u UserService) SetRedisUser(user *model.User) (string, error) {
	return model.User{}.SetRedisUser(user)
}

func (u UserService) GetUserByUsername(username string) *model.User {
	return model.User{}.GetUserByUsername(username)
}

func (u UserService) IsCreateUser(user *model.User) bool {
	return model.User{}.CreateUser(user)
}
