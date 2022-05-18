package controller

import (
	"PushSystem/config"
	"PushSystem/resp"
	"PushSystem/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetGroup(ctx *gin.Context) {
	userID := ctx.GetUint(config.HeadUserID)
	data, _ := service.MessageService{}.GetAllGroupsByUserID(userID)
	ctx.JSON(resp.SUCCESS, resp.NewSuccessResp(resp.WithData(data)))
	zap.L().Debug("resp")
}

func AddGroup(ctx *gin.Context) {
	messageService := new(service.MessageService)
	userID := ctx.GetUint(config.HeadUserID)
	err := ctx.BindJSON(messageService)
	if err != nil {
		return
	}
	err = messageService.AddGroup(userID, messageService)
	if err != nil {
		return
	}
	ctx.JSON(resp.SUCCESS, resp.NewSuccessResp(resp.WithData("ok")))
}
