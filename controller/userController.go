package controller

import (
	"PushSystem/model"
	"PushSystem/resp"
	"PushSystem/service"
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
		ctx.JSON(resp.InvalidParams, resp.NewInvalidResp())
		return
	}
	var userService = new(service.UserService)
	zap.L().Debug("Normal login")
	if userService.IsUserPassword(clientUser.Username, clientUser.Password) {
		user := userService.GetUserByUsername(clientUser.Username)
		if err != nil {
			panic(err)
		}
		token, err := util.CreateToken(user)
		if err != nil {
			zap.L().Error(err.Error())
			panic(err)
		}
		t := map[string]string{
			"token": token,
		}
		r := resp.NewSuccessResp(resp.WithData(t))
		ctx.JSON(resp.SUCCESS, r)
	} else {
		r := resp.New(resp.InvalidParams, resp.GetMessage(resp.InvalidParams), nil)
		ctx.JSON(resp.InvalidParams, r)
	}
	zap.L().Debug("user login ")
}
func Register(ctx *gin.Context) {
	clientUser := new(model.User)
	err := ctx.BindJSON(&clientUser)
	zap.L().Debug(clientUser.ToString())
	if err != nil {
		zap.L().Error(err.Error())
		ctx.JSON(resp.InvalidParams, resp.NewInvalidResp())
		return
	}
	var userService = new(service.UserService)
	clientUser.Salt = time.Now().UnixMilli()
	clientUser.Password = util.AddSalt(clientUser.Password, clientUser.Salt)
	if userService.IsCreateUser(clientUser) {
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

func ChangeInfo(ctx *gin.Context) {

}
