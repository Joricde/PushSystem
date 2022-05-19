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
		token := ctx.GetHeader("Authorization")
		if token == "" {
			r := resp.NewInvalidResp(resp.WithMessage("token为空"))
			ctx.JSON(code, r)
			ctx.Abort()
		} else {
			claims, err := util.ParseUserToken(token)
			if err != nil {
				r := resp.NewInvalidResp(resp.WithMessage("token错误"))
				ctx.JSON(code, r)
				ctx.Abort()
			} else if time.Now().Unix() > claims.Exp {
				r := resp.NewInvalidResp(resp.WithMessage("token已超时"))
				ctx.JSON(code, r)
				ctx.Abort()
			} else {
				code = resp.SUCCESS
				ctx.Set(config.HeadUserID, claims.UserID)
				ctx.Set(config.HeadUsername, claims.UserName)
				ctx.Next()
			}
		}
	}
}
