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
		code = resp.SUCCESS
		token := c.GetHeader("token")
		if token == "" {
			code = resp.InvalidParams
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = resp.InvalidParams
			} else {
				if time.Now().Unix() > int64(claims[config.TokenEXP].(float64)) {
					code = resp.InvalidParams
				}
			}
			c.Set(config.HeadUSERID, claims[config.TokenUID])
		}
		if code != resp.SUCCESS {
			r := resp.NewInvalidResp(resp.WithMessage(resp.GetMsg(code)))
			c.JSON(code, r)
			c.Abort()
			return
		}
		c.Next()
	}
}
