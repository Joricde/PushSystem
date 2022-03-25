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
		ctx.JSON(resp.ERROR, resp.Response{
			Code:    resp.ERROR,
			Message: resp.GetMsg(resp.ERROR),
			Data:    nil,
		})
		return
	}
	ctxGetData, ok := ctx.Get(config.HeadUSERID)
	userID, _ := ctxGetData.(uint)
	if ok {
		if userID != user.ID {
			ctx.JSON(resp.SUCCESS, resp.Response{
				Code:    resp.InvalidParams,
				Message: resp.GetMsg(resp.InvalidParams),
				Data:    nil,
			})
			return
		}
	}
	data := model.GetAllTaskByUserID(userID)
	ctx.JSON(resp.SUCCESS, resp.Response{
		Code:    resp.SUCCESS,
		Message: resp.GetMsg(resp.SUCCESS),
		Data:    data,
	})
	zap.L().Debug("resp")
}
