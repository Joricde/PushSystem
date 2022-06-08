package controller

import (
	"PushSystem/config"
	"PushSystem/resp"
	"PushSystem/service"
	"PushSystem/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

var GroupService service.MessageService

func GetTasks(ctx *gin.Context) {
	ok, taskService := TaskPermissionsIdentify(ctx)
	if ok {
		taskServices, e := taskService.GetAllTasksByGroupID(taskService.GroupID)
		if e != nil {
			zap.L().Debug(e.Error())
			errorResp(ctx)
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
			errorResp(ctx)
			return
		}
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
	}
}

func UpdateTask(ctx *gin.Context) {
	ok, taskService := TaskPermissionsIdentify(ctx)
	if ok {
		zap.L().Debug("resp")
		e := taskService.UpdateTask(taskService)
		if e != nil {
			errorResp(ctx)
			return
		}
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
	}
}

func DeleteTask(ctx *gin.Context) {
	ok, taskService := TaskPermissionsIdentify(ctx)
	if ok {
		zap.L().Debug("resp")
		e := taskService.DeleteTask(taskService.TaskID)
		if e != nil {
			return
		}
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
	}
}

func UploadFile(ctx *gin.Context) {
	userID := ctx.GetUint(config.HeadUserID)
	g, e := strconv.ParseUint(ctx.PostForm("GroupID"), 10, 16)
	if e != nil {
		zap.L().Debug("response")
		InvalidResp(ctx)
		return
	}
	GroupID := uint(g)
	t, e := strconv.ParseUint(ctx.PostForm("TaskID"), 10, 16)
	if e != nil {
		zap.L().Debug("response")
		errorResp(ctx)
		return
	}
	TaskID := uint(t)
	belongToUser, e := GroupService.IsBelongToUser(userID, GroupID)
	if e != nil {
		zap.L().Debug("response")
		InvalidResp(ctx)
		return
	}
	if belongToUser {
		file, e := ctx.FormFile("file")
		if e != nil {
			errorResp(ctx)
			zap.L().Debug(e.Error())
			return
		}
		taskService := new(service.TaskService)
		taskService.GroupID = GroupID
		taskService.TaskID = TaskID
		taskService.AppendixName = filepath.Base(file.Filename)
		result, e := util.CountSha256(file)
		if e != nil {
			zap.L().Debug(e.Error())
			errorResp(ctx)
			return
		}
		taskService.AppendixHash = result
		isExistFile := false
		savePath := config.Conf.FilePath + "/" + result
		_, err := os.Stat(savePath)
		if err == nil {
			isExistFile = true
		} else if os.IsNotExist(err) {
			isExistFile = false
		} else {
			zap.L().Debug(e.Error())
			errorResp(ctx)
			return
		}
		if !isExistFile {
			e = ctx.SaveUploadedFile(file, savePath)
			if e != nil {
				zap.L().Debug(e.Error())
				errorResp(ctx)
				return
			}
		}
		e = taskService.UpdateTaskAppendix(taskService)
		if e != nil {
			zap.L().Debug(e.Error())
			errorResp(ctx)
			return
		}
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
	}
}

func DownloadFile(ctx *gin.Context) {
	ok, taskService := TaskPermissionsIdentify(ctx)
	if ok {
		zap.L().Debug("resp")
		t, e := taskService.GetTasksByGTaskID(taskService.TaskID)
		if e != nil {
			zap.L().Debug(e.Error())
			return
		}
		savePath := config.Conf.FilePath + "/" + t.AppendixHash
		filename := url.QueryEscape(t.AppendixName)
		ctx.Writer.Header().Add("Content-Disposition",
			fmt.Sprintf("attachment;filename=%s", filename))
		ctx.File(savePath)
	}
}

func TaskPermissionsIdentify(ctx *gin.Context) (bool, *service.TaskService) {
	userID := ctx.GetUint(config.HeadUserID)
	taskService := new(service.TaskService)
	e := ctx.BindJSON(taskService)
	if e != nil {
		zap.L().Debug(e.Error())
		InvalidResp(ctx)
		return false, taskService
	}
	belongToUser, e := GroupService.IsBelongToUser(userID, taskService.GroupID)
	if e != nil {
		zap.L().Debug("response")
		errorResp(ctx)
		return false, taskService
	}
	if !belongToUser {
		zap.L().Debug("response")
		InvalidResp(ctx)
		return false, taskService
	}
	return true, taskService
}

func errorResp(ctx *gin.Context) {
	ctx.JSON(resp.SUCCESS, resp.NewErrorResp())
}

func InvalidResp(ctx *gin.Context) {
	ctx.JSON(resp.SUCCESS, resp.NewInvalidResp())
}
