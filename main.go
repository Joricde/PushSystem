package main

import (
	"PushSystem/config"
	_ "PushSystem/model"
	"PushSystem/routers"
	"PushSystem/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	run()
}

func run() {
	gin.SetMode(gin.DebugMode)
	r := routers.SetupRouter()
	r.TrustedPlatform = gin.PlatformCloudflare
	r.Use(util.GinLogger(), util.GinRecovery(true))
	err := r.Run(fmt.Sprintf(":%d", config.Conf.AppConfig.Port))
	if err != nil {
		panic(err)
	}
}
