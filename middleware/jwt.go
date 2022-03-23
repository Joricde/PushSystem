package middleware

import (
	"awesomeProject/util"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {

			claims, err := util.ParseToken(token, secret)
			if err != nil {
				code = util.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims["exp"].(int64) {
				code = util.ErrorAuthCheckTokenTimeout
			}
		}
		if code != util.SUCCESS {
			c.JSON(400, gin.H{
				"status": code,
				"msg":    util.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
