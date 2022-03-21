package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	// Use predefined header gin.PlatformXXX
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return
	}
	// Or set your own trusted request header for another trusted proxy service
	// Don't set it to any suspect request header, it's unsafe

	router.GET("/", func(c *gin.Context) {
		// If you set TrustedPlatform, ClientIP() will resolve the
		// corresponding header and return IP directly
		fmt.Printf("ClientIP: %s\n", c.ClientIP())
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	err = router.Run()
	if err != nil {
		return
	}
}
