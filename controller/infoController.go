package controller

import (
	"PushSystem/model"
	"PushSystem/util"
	"PushSystem/util/status"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetMsg(ctx *gin.Context) {
	user := model.User{}
	err := ctx.BindJSON(user)
	if err != nil {
		zap.L().Debug(err.Error())
		return
	}
	username := ctx.Query("username")
	password := ctx.Query("password")
	zap.L().Debug("login")
	byUsr := model.GetUserByUsername(user.Username)
	if err != nil {
		return
	}
	if byUsr.Password == password {
		token, err2 := util.CreateToken(byUsr)
		if err2 != nil {
			return
		}
		ctx.JSON(status.SUCCESS, token)
	}
	zap.L().Debug("usr+pwd " + username + password)
}
