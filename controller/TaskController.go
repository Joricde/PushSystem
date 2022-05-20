package controller

import (
	"PushSystem/config"
	"PushSystem/resp"
	"PushSystem/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var GroupService service.MessageService

func GetTasks(ctx *gin.Context) {
	ok, taskService := TaskPermissionsIdentify(ctx)
	if ok {
		taskServices, e := taskService.GetAllTasksByGroupID(taskService.GroupID)
		if e != nil {
			zap.L().Debug(e.Error())
			return
		}
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp(resp.WithData(taskServices)))
	}
	zap.L().Debug("resp")
}

func AddTask(ctx *gin.Context) {
	ok, taskService := TaskPermissionsIdentify(ctx)
	if ok {
		zap.L().Debug("resp")
		e := taskService.AddTask(taskService)
		if e != nil {
			return
		}
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
	}
}

func TaskPermissionsIdentify(ctx *gin.Context) (bool, *service.TaskService) {
	userID := ctx.GetUint(config.HeadUserID)
	taskService := new(service.TaskService)
	e := ctx.BindJSON(taskService)
	if e != nil {
		zap.L().Debug(e.Error())
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp())
		return false, taskService
	}
	belongToUser, e := GroupService.IsBelongToUser(userID, taskService.GroupID)
	if e != nil {
		zap.L().Debug("response")
		ctx.JSON(resp.SUCCESS, resp.NewErrorResp())
		return false, taskService
	}
	if !belongToUser {
		zap.L().Debug("response")
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp())
		return false, taskService
	}
	return true, taskService
}
