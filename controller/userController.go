package controller

import (
	"PushSystem/api"
	"PushSystem/config"
	"PushSystem/model"
	"PushSystem/resp"
	"PushSystem/service"
	"PushSystem/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type UserInfo struct {
	User model.User
	Pwd  model.UserPwd
}

var userService = new(service.UserService)

func Login(ctx *gin.Context) {
	userInfo := new(UserInfo)
	if !isBindUser(ctx, userInfo) {
		return
	}
	user := userService.GetUserByUsername(userInfo.User.Username)
	if user.ID == 0 {
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(resp.WithMessage("用户名或密码错误")))
		return
	} else if userService.IsUserPassword(user.ID, userInfo.Pwd.Password) {
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
	token := ctx.PostForm("wechat_token")
	zap.L().Debug("token: " + token)
	result := api.CheckQRLogin(token)
	if result.Uid > 0 {
		user := service.UserService{}.GetUserByWechatID(result.Uid)
		if user.ID > 0 {
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
			ctx.JSON(resp.SUCCESS, resp.NewErrorResp(
				resp.WithMessage("用户不存在，请注册并绑定微信")))
		}
	} else {
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(
			resp.WithMessage("获取微信信息失败")))
	}
}

func Register(ctx *gin.Context) {
	userInfo := new(UserInfo)
	if !isBindUser(ctx, userInfo) {
		return
	}
	userInfo.Pwd.Salt = time.Now().UnixMilli()
	userInfo.Pwd.Password = util.AddSalt(userInfo.Pwd.Password, userInfo.Pwd.Salt)
	zap.L().Debug(userInfo.ToString())
	if userService.CreateUser(&userInfo.User, &userInfo.Pwd) {
		zap.L().Debug("create " + userInfo.ToString())
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp(resp.WithData(
			map[string]string{
				"result": resp.GetMessage(resp.SUCCESS),
			})))
	} else {
		ctx.JSON(resp.ERROR, resp.NewErrorResp())
	}
}

func CheckUsernameExist(ctx *gin.Context) {
	username := ctx.Query("username")
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
	userInfo := new(UserInfo)
	if !isBindUser(ctx, userInfo) {
		return
	}
	zap.L().Debug(userInfo.ToString())
	uid := ctx.GetUint(config.TokenUID)
	userInfo.User.ID = uid
	if userService.SetUserInfo(&userInfo.User) {
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
	} else {
		ctx.JSON(resp.SUCCESS, resp.NewErrorResp())
	}
}

func CheckUsePwd(ctx *gin.Context) {
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
	pwd := ctx.PostForm("password")
	uid := ctx.GetUint(config.TokenUID)
	if userService.SetPassword(uid, pwd) {
		ctx.JSON(resp.SUCCESS, resp.NewSuccessResp())
	} else {
		ctx.JSON(resp.SUCCESS, resp.NewErrorResp())
	}
}

func ChangeWechatKey(ctx *gin.Context) {
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

func isBindUser(ctx *gin.Context, user *UserInfo) bool {
	err := ctx.BindJSON(user)
	zap.L().Debug(fmt.Sprintln(user.User.Username))
	if err != nil {
		zap.L().Error(err.Error())
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp())
		return false
	} else if user.User.Username == "" {
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(resp.WithMessage("用户名为空")))
		return false
	} else if user.Pwd.Password == "" {
		ctx.JSON(resp.SUCCESS, resp.NewInvalidResp(resp.WithMessage("密码为空")))
		return false
	}
	return true
}

func (u UserInfo) ToString() string {
	return fmt.Sprintf("%+v", u)
}
