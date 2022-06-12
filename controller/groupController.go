package controller

import (
	"PushSystem/config"
	"PushSystem/resp"
	"PushSystem/service"
	"PushSystem/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
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
		e := messageService.SetGroupTitle(messageService)
		if e != nil {
			zap.L().Debug(e.Error())
			return
		}
		zap.L().Debug("response")
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
	}
}

func UpdateGroupSort(ctx *gin.Context) {
	ok, messageService := PermissionsIdentify(ctx)
	if ok {
		userID := ctx.GetUint(config.HeadUserID)
		e := messageService.SetGroupSort(userID, messageService.GroupID, messageService.Sort)
		if e != nil {
			zap.L().Debug(e.Error())
			return
		}
		zap.L().Debug("response")
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
	}
}

func DeleteGroup(ctx *gin.Context) {
	userID := ctx.GetUint(config.HeadUserID)
	ok, messageService := PermissionsIdentify(ctx)
	if ok {
		zap.L().Debug(fmt.Sprint(messageService))
		e := messageService.DeleteGroupByGroupID(userID, messageService.GroupID)
		if e != nil {
			zap.L().Debug(e.Error())
			return
		}
		zap.L().Debug("response")
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
	}
}

func SetShareable(ctx *gin.Context) {
	ok, messageService := PermissionsIdentify(ctx)
	if ok {
		userID := ctx.GetUint(config.HeadUserID)
		p, e := messageService.IsGroupCreator(userID, messageService.GroupID)
		if e != nil {
			zap.L().Debug("response")
			ctx.JSON(resp.SUCCESS, resp.NewErrorResp())
			return
		}
		if p {
			m, e := messageService.SetGroupShare(messageService.GroupID, messageService.IsShare)
			if e != nil {
				zap.L().Debug("response")
				ctx.JSON(resp.SUCCESS, resp.NewErrorResp())
				return
			}
			if m.IsShare {
				shareToken, e := util.CreateShareToken(m.GroupID)
				if e != nil {
					ctx.JSON(resp.SUCCESS, resp.NewErrorResp())
				}
				r := map[string]string{
					"shareToken": shareToken,
				}
				zap.L().Debug("response")
				ctx.JSON(resp.SUCCESS, resp.NewSuccessResp(resp.WithData(r)))
			} else {
				ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
			}
		} else {
			ctx.JSON(resp.SUCCESS, resp.NewErrorResp(resp.WithMessage("非创建者")))
		}
	}
}

func JoinShareGroup(ctx *gin.Context) {
	shareToken := ctx.Query("shareToken")
	zap.L().Debug(shareToken)
	userID := ctx.GetUint(config.HeadUserID)
	groupID, err := util.ParseShareToken(shareToken)
	sort := ctx.GetInt("sort")
	if err != nil {
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(
			resp.WithMessage("加入共享组失败，请检查链接是否正确")))
		return
	}
	ms := new(service.MessageService)
	IsNewJoin, err := ms.JoinShareGroup(userID, groupID, sort)
	if err != nil {
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(
			resp.WithMessage("加入共享组失败，请检查链接是否正确")))
		return
	} else if IsNewJoin {
		r := map[string]string{
			"groupID": strconv.Itoa(int(groupID)),
		}
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp(resp.WithData(r)))
	} else {
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(resp.WithMessage("已加入该组")))
	}
}

func PermissionsIdentify(ctx *gin.Context) (bool, *service.MessageService) {
	userID := ctx.GetUint(config.HeadUserID)
	messageService := new(service.MessageService)
	e := ctx.BindJSON(messageService)
	if e != nil {
		zap.L().Debug("response")
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp())
		return false, messageService
	}
	belongToUser, e := messageService.IsBelongToUser(userID, messageService.GroupID)
	if e != nil {
		zap.L().Debug("response")
		ctx.JSON(resp.SUCCESS, resp.NewErrorResp())
		return false, messageService
	}
	if !belongToUser {
		zap.L().Debug("response")
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp())
		return false, messageService
	}
	return true, messageService
}
