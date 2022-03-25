package controller

import (
	"PushSystem/model"
	"PushSystem/resp"
	"PushSystem/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Login(ctx *gin.Context) {
	clientUser := new(model.User)
	err := ctx.BindJSON(clientUser)
	if err != nil {
		zap.L().Debug("user: " + ctx.Query("username"))
		zap.L().Error(err.Error())
		ctx.JSON(resp.InvalidParams, resp.Response{
			Code:    resp.InvalidParams,
			Message: resp.GetMsg(resp.InvalidParams),
			Data:    nil,
		})
		return
	}
	zap.L().Debug("Normal login")
	user := model.GetUserByUsername(clientUser.Username)
	check := util.AddSalt(clientUser.Password, user.Salt)

	if check == user.Password {
		_, err = model.SetRedisUser(user)
		if err != nil {
			panic(err)
		}
		token, err := util.CreateToken(user)
		if err != nil {
			panic(err)
		}
		t := map[string]string{
			"token": token,
		}
		ctx.JSON(resp.SUCCESS, resp.Response{
			Code:    resp.SUCCESS,
			Message: resp.GetMsg(resp.SUCCESS),
			Data:    t,
		})
	} else {
		ctx.JSON(resp.ErrorAuth, resp.Response{
			Code:    resp.ErrorAuth,
			Message: resp.GetMsg(resp.ErrorAuth),
			Data:    nil,
		})
	}
	zap.L().Debug("user login ")
}
func Register(ctx *gin.Context) {
	clientUser := new(model.User)
	err := ctx.BindJSON(&clientUser)
	zap.L().Debug(clientUser.ToString())
	if err != nil {
		zap.L().Error(err.Error())
		ctx.JSON(resp.InvalidParams, resp.Response{
			Code:    resp.InvalidParams,
			Message: resp.GetMsg(resp.InvalidParams),
			Data:    nil,
		})
		return
	}
	clientUser.Salt = time.Now().UnixMilli()
	clientUser.Password = util.AddSalt(clientUser.Password, clientUser.Salt)
	info := model.CreateUser(clientUser)
	r := map[string]string{
		"result": resp.GetMsg(resp.SUCCESS),
	}
	if len(info) == 0 {
		ctx.JSON(resp.SUCCESS, resp.Response{
			Code:    resp.SUCCESS,
			Message: resp.GetMsg(resp.SUCCESS),
			Data:    r,
		})
	} else {
		ctx.JSON(resp.ERROR, resp.Response{
			Code:    resp.ERROR,
			Message: resp.GetMsg(resp.ERROR),
			Data:    nil,
		})
	}
	zap.L().Debug(clientUser.Username + clientUser.Password)
}
