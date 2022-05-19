package routers

import (
	"PushSystem/config"
	"PushSystem/controller"
	"PushSystem/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if config.Conf.AppConfig.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Static("/static", "static")
	//router.Use(middleware.GlobeLimitRequest())
	api := router.Group("api/v1")
	{
		user := api.Group("user")
		{
			user.POST("login", controller.Login)
			user.GET("wechat_qr", controller.GetWechatQR)
			user.POST("wechat_check", controller.CheckWechatLogin)
			user.GET("check_name", controller.CheckUsernameExist)
			user.POST("register", controller.Register)
		}
		authed := api.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("group", controller.GetGroup)
			authed.POST("group", controller.AddGroup)
			authed.PUT("group", controller.UpdateGroup)
			authed.DELETE("group", controller.DeleteGroup)
			authed.PUT("group/share", controller.SetShareable)
			authed.GET("group/join", controller.JoinShareGroup)
		}
		//authed.POST("group")
	}
	return router

}
