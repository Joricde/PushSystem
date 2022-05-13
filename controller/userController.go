package controller

import (
	"PushSystem/api"
	"PushSystem/config"
	"PushSystem/resp"
	"PushSystem/service"
	"PushSystem/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func Login(ctx *gin.Context) {
	var userService = new(service.UserService)
	if !isBindUser(ctx, userService) {
		return
	}
	user := userService.GetUserByUsername(userService.Username)
	if user.ID == 0 {
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(resp.WithMessage("用户名或密码错误")))
		return
	} else if userService.IsUserPassword(user.ID, userService.Password) {
		token, err := util.CreateToken(user)
		if err != nil {
			zap.L().Error(err.Error())
		}
		t := map[string]string{
			"token": token,
		}
		ok := userService.SetRedisUser(user)
		if !ok {

			return
		}
		r := resp.NewSuccessResp(resp.WithData(t))
		ctx.JSON(resp.SUCCESS, r)
	} else {
		r := resp.NewInvalidResp(resp.WithMessage("用户名或密码错误"))
		ctx.JSON(resp.SUCCESS, r)
	}
}

func GetWechatQR(ctx *gin.Context) {
	r := api.GetWechatQR()
	if r.Data.ExpireSeconds > 0 {
		response := resp.NewSuccessResp(resp.WithData(r.Data))
		ctx.JSON(resp.SUCCESS, response)
	} else {
		ctx.JSON(resp.ERROR, resp.NewErrorResp())
	}
}

func CheckWechatLogin(ctx *gin.Context) {
	var userService = new(service.UserService)
	token := ctx.PostForm("wechat_token")
	zap.L().Debug("token: " + token)
	result := api.CheckQRLogin(token)
	if result.Uid > 0 {
		zap.L().Debug("GetUserByWechatID  " + strconv.FormatInt(result.Uid, 16))
		user := userService.GetUserByWechatID(result.Uid)
		createUser := false
		zap.L().Debug("GetUserByWechatID " + user.ToString())
		if user.ID == 0 {
			for {
				userService.Salt = time.Now().UnixMilli()
				userService.Username = strconv.FormatInt(userService.Salt, 16)
				if !userService.IsUsernameExist(userService.Username) {
					break
				}
			}
			userService.WechatID = result.Uid
			userService.WechatKey = result.SendKey
			userService.Password = util.AddSalt(userService.Username, userService.Salt)
			if userService.CreateUser() {
				createUser = true
			} else {
				ctx.JSON(resp.ERROR, resp.NewErrorResp(resp.WithMessage("创建用户失败")))
				return
			}
		}
		user = userService.GetUserByWechatID(result.Uid)
		token, err := util.CreateToken(user)
		if err != nil {
			zap.L().Error(err.Error())
		}
		t := map[string]string{
			"token": token,
		}
		ok := userService.SetRedisUser(user)
		if !ok {
			return
		}
		r := resp.NewSuccessResp(resp.WithData(t), resp.WithMessage("登录成功"))
		if createUser {
			r = resp.NewSuccessResp(resp.WithData(t), resp.WithMessage("已经创建用户,请及时修改用户名与密码"))
		}
		ctx.JSON(resp.SUCCESS, r)

	} else {
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(
			resp.WithMessage("获取微信信息失败")))
	}
}

func Register(ctx *gin.Context) {
	var userService = new(service.UserService)
	if !isBindUser(ctx, userService) {
		return
	}
	userService.Salt = time.Now().UnixMilli()
	userService.Password = util.AddSalt(userService.Password, userService.Salt)
	zap.L().Debug(userService.ToString())
	if userService.CreateUser() {
		zap.L().Debug("create " + userService.ToString())
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp(resp.WithData(
			map[string]string{
				"result": resp.GetMessage(resp.SUCCESS),
			})))
	} else {
		ctx.JSON(resp.ERROR, resp.NewErrorResp())
	}
}

func RegisterFromWechat(ctx *gin.Context) {

}

func CheckUsernameExist(ctx *gin.Context) {
	var userService = new(service.UserService)
	username := ctx.Query("username")
	zap.L().Debug(username)
	if username == "" {
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(resp.WithMessage("用户名不能为空")))
		return
	} else {
		if userService.IsUsernameExist(username) {
			ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(resp.WithMessage("用户名已存在")))
		} else {
			ctx.JSON(resp.SUCCESS, resp.NewSuccessResp(resp.WithMessage("ok")))
		}
	}
}

func ChangeUserInfo(ctx *gin.Context) {
	var userService = new(service.UserService)
	if !isBindUser(ctx, userService) {
		return
	}
	zap.L().Debug(userService.ToString())
	uid := ctx.GetUint(config.TokenUID)
	if userService.SetUserInfoByID(uid) {
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
	} else {
		ctx.JSON(resp.SUCCESS, resp.NewErrorResp())
	}
}

func CheckUsePwd(ctx *gin.Context) {
	var userService = new(service.UserService)
	pwd := ctx.PostForm("password")
	uid := ctx.GetUint(config.TokenUID)
	if pwd == "" {
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(resp.WithMessage("密码不能为空")))
		return
	} else {
		if userService.IsUserPassword(uid, pwd) {
			ctx.JSON(resp.SUCCESS, resp.NewSuccessResp(resp.WithMessage("ok")))
		} else {
			ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(resp.WithMessage("密码错误")))
		}
	}
}

func ChangeUserPWD(ctx *gin.Context) {
	var userService = new(service.UserService)
	pwd := ctx.PostForm("password")
	uid := ctx.GetUint(config.TokenUID)
	if userService.SetPassword(uid, pwd) {
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
	} else {
		ctx.JSON(resp.SUCCESS, resp.NewErrorResp())
	}
}

func ChangeWechatKey(ctx *gin.Context) {
	var userService = new(service.UserService)
	WechatKey := ctx.PostForm("wechat_key")
	uid := ctx.GetUint(config.TokenUID)
	if userService.SetWechatKey(uid, WechatKey) {
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
	} else {
		ctx.JSON(resp.SUCCESS, resp.NewErrorResp())
	}
}

func RetrievePwd(ctx *gin.Context) {

}

func isBindUser(ctx *gin.Context, user *service.UserService) bool {
	err := ctx.BindJSON(user)
	zap.L().Debug(fmt.Sprintln(user.Username))
	if err != nil {
		zap.L().Error(err.Error())
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp())
		return false
	} else if user.Username == "" {
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(resp.WithMessage("用户名为空")))
		return false
	} else if user.Password == "" {
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(resp.WithMessage("密码为空")))
		return false
	}
	return true
}
