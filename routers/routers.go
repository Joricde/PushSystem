package routers

import (
	"awesomeProject/config"
	"awesomeProject/controller"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if config.Conf.AppConfig.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Static("/static", "static")

	api := router.Group("api/")
	{
		api.POST("login", controller.Login)
		api.POST("register", controller.Register)
		authed := api.Group("/")
		authed.Use(middleware.JWT())

	}
	return router
}
