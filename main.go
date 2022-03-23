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
	//zap.L().Debug(fmt.Sprintln(config.Conf))
	r := routers.SetupRouter()
	r.TrustedPlatform = gin.PlatformCloudflare
	r.Use(util.GinLogger(), util.GinRecovery(true))
	err := r.Run(fmt.Sprintf(":%d", config.Conf.AppConfig.Port))
	if err != nil {
		panic(err)
	}
}
