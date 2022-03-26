package middleware

import (
	"PushSystem/config"
	"PushSystem/resp"
	"PushSystem/util"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := resp.InvalidParams
		token := ctx.GetHeader("token")
		if token == "" {
			r := resp.NewInvalidResp(resp.WithMessage("token为空"))
			ctx.JSON(code, r)
			ctx.Abort()
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				r := resp.NewInvalidResp(resp.WithMessage("token错误"))
				ctx.JSON(code, r)
			} else if time.Now().Unix() > int64(claims[config.TokenEXP].(float64)) {
				r := resp.NewInvalidResp(resp.WithMessage("token已超时"))
				ctx.JSON(code, r)
				ctx.Abort()
			} else {
				code = resp.SUCCESS
				ctx.Set(config.HeadUserID, claims[config.TokenUID])
				ctx.Set(config.HeadUsername, claims[config.TokenUsername])
				ctx.Next()
			}
		}
	}
}
