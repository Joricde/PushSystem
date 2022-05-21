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
			user.GET("wechat_qr", controller.GetWechatQR)
			user.POST("wechat_check", controller.CheckWechatLogin)
			user.POST("register", controller.Register)
			user.POST("login", controller.Login)
			user.GET("check_name", controller.CheckUsernameExist)

		}
		authed := api.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("group", controller.GetGroup)
			authed.GET("group/join", controller.JoinShareGroup)
			authed.POST("group", controller.AddGroup)
			authed.PUT("group", controller.UpdateGroup)
			authed.PUT("group/share", controller.SetShareable)
			authed.DELETE("group", controller.DeleteGroup)

		}
		{
			authed.GET("/task", controller.GetTasks)
			authed.POST("/task", controller.AddTask)
			authed.POST("/task/upload", controller.UploadFile)
			authed.GET("/task/download", controller.DownloadFile)
			authed.PUT("/task", controller.UpdateTask)
			authed.DELETE("/task", controller.DeleteTask)
		}
		//authed.POST("group")
		{
			api.GET("/task/ws", controller.WebsocketConnect)
		}
	}
	return router

}
