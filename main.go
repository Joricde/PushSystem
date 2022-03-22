package main

import (
	"awesomeProject/config"
	"awesomeProject/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.TrustedPlatform = gin.PlatformCloudflare
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	err := r.Run(fmt.Sprintf(":%d", config.Conf.AppConfig.Port))
	if err != nil {
		panic(err)
	}
}
