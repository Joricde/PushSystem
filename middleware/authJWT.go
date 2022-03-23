package middleware

import (
	"PushSystem/util"
	"PushSystem/util/msg"
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

			claims, err := util.ParseToken(token)
			if err != nil {
				code = msg.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims["exp"].(int64) {
				code = msg.ErrorAuthCheckTokenTimeout
			}
		}
		if code != msg.SUCCESS {
			c.JSON(400, gin.H{
				"status": code,
				"msg":    msg.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
