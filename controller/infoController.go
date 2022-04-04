package controller

import (
	"PushSystem/config"
	"PushSystem/model"
	"PushSystem/resp"
	"PushSystem/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetMsg(ctx *gin.Context) {
	user := new(model.User)
	err := ctx.BindJSON(user)
	if err != nil {
		zap.L().Error(err.Error())
		ctx.JSON(resp.SUCCESS, resp.NewErrorResp())
		return
	}
	userID := ctx.GetUint(config.HeadUserID)
	username := ctx.GetString(config.HeadUsername)
	if username != user.Username {
		ctx.JSON(resp.InvalidParams, resp.NewInvalidResp(resp.WithMessage("无权访问")))
		return
	}
	data := service.MsgService{}.GetAllTaskByUserID(userID)
	ctx.JSON(resp.SUCCESS, resp.NewSuccessResp(resp.WithData(data)))
	zap.L().Debug("resp")
}
