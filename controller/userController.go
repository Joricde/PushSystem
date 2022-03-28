package controller

import (
	"PushSystem/model"
	"PushSystem/resp"
	"PushSystem/service"
	"PushSystem/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type UserInfo struct {
	User model.User
	Pwd  model.UserPwd
}

func (i UserInfo) ToString() string {
	return fmt.Sprintf("%+v", i)
}

func Login(ctx *gin.Context) {
	userInfo := new(UserInfo)
	err := ctx.BindJSON(userInfo)
	zap.L().Debug(fmt.Sprintln(userInfo))
	if err != nil {
		zap.L().Error(err.Error())
		ctx.JSON(resp.InvalidParams, resp.NewInvalidResp())
		return
	}
	var userService = new(service.UserService)
	user := userService.GetUserByUsername(userInfo.User.Username)
	if user.ID == 0 {
		ctx.JSON(resp.InvalidParams, resp.NewInvalidResp(resp.WithMessage("用户名不存在")))
		return
	} else if userService.IsUserPassword(user.ID, userInfo.Pwd.Password) {
		token, err := util.CreateToken(user)
		if err != nil {
			zap.L().Error(err.Error())
		}
		t := map[string]string{
			"token": token,
		}
		_, err = userService.SetRedisUser(user)
		if err != nil {
			return
		}
		r := resp.NewSuccessResp(resp.WithData(t))
		ctx.JSON(resp.SUCCESS, r)
	} else {
		r := resp.NewInvalidResp(resp.WithMessage("用户名或密码错误"))
		ctx.JSON(resp.InvalidParams, r)
	}
}

func Register(ctx *gin.Context) {
	userInfo := new(UserInfo)
	if !isBindUser(ctx, userInfo) {
		return
	}
	var userService = new(service.UserService)
	userInfo.Pwd.Salt = time.Now().UnixMilli()
	userInfo.Pwd.Password = util.AddSalt(userInfo.Pwd.Password, userInfo.Pwd.Salt)
	zap.L().Debug(userInfo.ToString())
	if userService.IsCreateUser(&userInfo.User, &userInfo.Pwd) {
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp(resp.WithData(
			map[string]string{
				"result": resp.GetMessage(resp.SUCCESS),
			})))
	} else {
		ctx.JSON(resp.ERROR, resp.NewErrorResp())
	}
}

func CheckUsernameExist(ctx *gin.Context) {
	var userService = new(service.UserService)
	username := ctx.Query("username")
	if username != "" {
		if userService.IsUsernameExist(username) {
			ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(resp.WithMessage("用户名已存在")))
		}
	} else {
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp(resp.WithMessage("ok")))
	}
}

func isBindUser(ctx *gin.Context, user *UserInfo) bool {
	err := ctx.BindJSON(user)
	zap.L().Debug(fmt.Sprintln(user))
	if err != nil {
		zap.L().Error(err.Error())
		ctx.JSON(resp.InvalidParams, resp.NewInvalidResp())
		return false
	} else if user.User.Username == "" {
		ctx.JSON(resp.InvalidParams, resp.NewInvalidResp(resp.WithMessage("用户名为空")))
		return false
	} else if user.Pwd.Password == "" {
		ctx.JSON(resp.InvalidParams, resp.NewInvalidResp(resp.WithMessage("密码为空")))
		return false
	}
	return true
}
