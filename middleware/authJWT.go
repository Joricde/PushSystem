package middleware

import (
	"PushSystem/model"
	"PushSystem/util"
	"PushSystem/util/status"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
				code = status.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims["exp"].(int64) {
				result, err := model.GetRedisUserByID(claims["id"].(int))
				if err != nil {
					code = status.ErrorAuthCheckTokenTimeout
				} else {
					_, err = model.SetRedisUser(result)
					if err != nil {
						zap.L().Error(err.Error())
					}
					util.RenewToken(result)
				}
			}
		}
		if code != status.SUCCESS {
			c.JSON(400, gin.H{
				"status": code,
				"msg":    status.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
