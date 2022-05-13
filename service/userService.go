package service

import (
	"PushSystem/model"
	"PushSystem/util"
	"fmt"
	"go.uber.org/zap"
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

var UserModel = new(model.User)

func (u UserService) IsUsernameExist(username string) bool {
	newUser := new(model.User)
	newUser = UserModel.GetUserByUsername(username)
	zap.L().Debug("username: " + username + " nweUsername: " + newUser.Username)
	zap.L().Debug(newUser.ToString())
	if username == newUser.Username {
		return true
	} else {
		return false
	}
}

func (u UserService) IsUserPassword(uid uint, password string) bool {
	userPwd := UserModel.GetPasswordByUserID(uid)
	passwordEscape := util.AddSalt(password, userPwd.Salt)
	if passwordEscape == userPwd.PasswordHash {
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
		Password: model.Password{
			PasswordHash: u.Password,
			Salt:         u.Salt,
			WechatKey:    u.WechatKey,
		},
	}
	return user.CreateUser(&user)
}

func (u UserService) SetPassword(uid uint, pwd string) bool {
	userPwd := new(model.Password)
	userPwd.Salt = time.Now().UnixMilli()
	userPwd.PasswordHash = util.AddSalt(pwd, userPwd.Salt)
	return UserModel.UpdatePassword(uid, userPwd)
}

func (u UserService) SetWechatKey(uid uint, wechatKey string) bool {
	return UserModel.UpdateWechatKey(uid, wechatKey)
}

func (u UserService) SetUserInfoByID(uid uint) bool {
	user := model.User{
		Username: u.Username,
		Nickname: u.Nickname,
		Phone:    u.Phone,
		Email:    u.Email,
		WechatID: u.WechatID,
	}
	return UserModel.UpdateUserInfo(&user)
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
