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
	messageService := new(service.MessageService)
	data, _ := messageService.GetAllGroupsByUserID(userID)
	ctx.JSON(resp.SUCCESS, resp.NewSuccessResp(resp.WithData(data)))
	zap.L().Debug("resp")
}

func AddGroup(ctx *gin.Context) {
	userID := ctx.GetUint(config.HeadUserID)
	messageService := new(service.MessageService)
	e := ctx.BindJSON(messageService)
	if e != nil {
		return
	}
	e = messageService.AddGroup(userID, messageService)
	if e != nil {
		return
	}
	ctx.JSON(resp.SUCCESS, resp.NewSuccessResp(resp.WithData("ok")))
}

func UpdateGroup(ctx *gin.Context) {
	ok, messageService := PermissionsIdentify(ctx)
	if ok {
		e := messageService.SetGroupInfo(messageService)
		if e != nil {
			return
		}
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
	}
}

func DeleteGroup(ctx *gin.Context) {
	userID := ctx.GetUint(config.HeadUserID)
	ok, messageService := PermissionsIdentify(ctx)
	if ok {
		e := messageService.DeleteGroupByGroupID(userID, messageService.GroupID)
		if e != nil {
			return
		}
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
	}
}

func SetShareable(ctx *gin.Context) {
	ok, messageService := PermissionsIdentify(ctx)
	if ok {
		e := messageService.SetGroupInfo(messageService)
		if e != nil {
			ctx.JSON(resp.SUCCESS, resp.NewErrorResp())
			return
		}
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
	}
}

func JoinShareGroup(ctx *gin.Context) {

}

func PermissionsIdentify(ctx *gin.Context) (bool, *service.MessageService) {
	userID := ctx.GetUint(config.HeadUserID)
	messageService := new(service.MessageService)
	e := ctx.BindJSON(messageService)
	if e != nil {
		ctx.JSON(resp.SUCCESS, resp.NewErrorResp())
		return false, messageService
	}
	belongToUser, e := messageService.IsBelongToUser(userID, messageService.GroupID)
	if e != nil {
		ctx.JSON(resp.SUCCESS, resp.NewErrorResp())
		return false, messageService
	}
	if !belongToUser {
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp())
		return false, messageService
	}
	return true, messageService
}
