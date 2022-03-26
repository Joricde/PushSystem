package controller

import (
	"PushSystem/config"
	"PushSystem/model"
	"PushSystem/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetMsg(ctx *gin.Context) {
	user := new(model.User)
	err := ctx.BindJSON(user)
	if err != nil {
		zap.L().Debug(err.Error())
		ctx.JSON(resp.ERROR, resp.NewErrorRResp())
		return
	}
	userID := ctx.GetUint(config.HeadUserID)
	username := ctx.GetString(config.HeadUsername)
	if username != user.Username {
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(resp.WithMessage("无权访问")))
		return
	}
	data := model.GetAllTaskByUserID(userID)
	ctx.JSON(resp.SUCCESS, resp.NewSuccessResp(resp.WithData(data)))
	zap.L().Debug("resp")
}
