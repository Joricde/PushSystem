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
		ctx.JSON(resp.InvalidParams, resp.NewInvalidResp())
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
		r := resp.NewSuccessResp(resp.WithData(t))
		ctx.JSON(resp.SUCCESS, r)
	} else {
		r := resp.New(resp.InvalidParams, resp.GetMsg(resp.ErrorAuth), nil)
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
	clientUser.Salt = time.Now().UnixMilli()
	clientUser.Password = util.AddSalt(clientUser.Password, clientUser.Salt)
	result := model.CreateUser(clientUser)
	if len(result) == 0 {
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp(resp.WithData(
			map[string]string{
				"result": resp.GetMsg(resp.SUCCESS),
			})))
	} else {
		ctx.JSON(resp.ERROR, resp.NewERRORResp())
	}
	zap.L().Debug(clientUser.Username + clientUser.Password)
}

func ChangeInfo() {

}
