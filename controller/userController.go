package controller

import (
	"PushSystem/model"
	"PushSystem/util"
	"PushSystem/util/msg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	zap.L().Debug("login")
	byUsr, err := model.CheckUsername(username)
	if err != nil {
		return
	}
	if byUsr.Password == password {
		token, err2 := util.CreateToken(username)
		if err2 != nil {
			return
		}
		ctx.JSON(msg.SUCCESS, token)
	}
	zap.L().Debug("usr+pwd " + username + password)
}
func Register(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	u := model.User{
		Username: username,
		Password: password,
	}
	_, err := model.CreateUser(&u)
	if err != nil {
		return
	}
	ctx.JSON(msg.SUCCESS, username)
	zap.L().Info(username + password)
}
