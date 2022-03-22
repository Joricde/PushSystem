package main

import (
	"awesomeProject/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.TrustedPlatform = gin.PlatformCloudflare
	err := r.Run(fmt.Sprintf(":%d", config.Conf.AppConfig.Port))
	if err != nil {
		panic(err)
	}
}
