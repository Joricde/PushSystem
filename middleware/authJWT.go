package middleware

import (
	"PushSystem/config"
	"PushSystem/resp"
	"PushSystem/util"
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
				code = resp.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > int64(claims[config.TokenEXP].(float64)) {
				code = resp.ErrorAuthCheckTokenTimeout
			}
			c.Set(config.HeadUSERID, claims[config.TokenUID])
		}
		if code != resp.SUCCESS {
			c.JSON(400, gin.H{
				"status": code,
				"msg":    resp.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
