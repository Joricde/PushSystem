package controller

import (
	"PushSystem/model"
	"PushSystem/util"
	"PushSystem/util/status"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Login(ctx *gin.Context) {
	clientUser := new(model.User)
	err := ctx.BindJSON(clientUser)
	if err != nil {
		zap.L().Debug("user: " + ctx.Query("username"))
		zap.L().Debug(err.Error())
		panic(err)
	}
	zap.L().Debug("Normal login")
	user := model.GetUserByUsername(clientUser.Username)
	check := util.PasswordAddSalt(clientUser.Password, user.Salt)
	if check == user.Password {
		_, err = model.SetRedisUser(user)
		if err != nil {
			panic(err)
		}
		token, err := util.CreateToken(user)
		if err != nil {
			panic(err)
		}
		ctx.JSON(status.SUCCESS, token)
	} else {
		ctx.JSON(status.ErrorAuth, clientUser.Username)
	}
	zap.L().Debug("usr+pwd ")
}
func Register(ctx *gin.Context) {
	clientUser := new(model.User)
	err := ctx.BindJSON(&clientUser)
	zap.L().Debug(clientUser.ToString())
	if err != nil {
		zap.L().Debug(err.Error())
		panic(err)
	}
	clientUser.Salt = time.Now().UnixMilli()
	clientUser.Password = util.PasswordAddSalt(clientUser.Password, clientUser.Salt)
	info := model.CreateUser(clientUser)
	if len(info) == 0 {
		ctx.JSON(status.SUCCESS, clientUser)
	} else {
		ctx.JSON(status.InvalidParams, clientUser.Username)
	}
	zap.L().Info(clientUser.Username + clientUser.Password)
}
