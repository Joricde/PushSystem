package middleware

import (
	"PushSystem/model"
	"PushSystem/resp"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"time"
)

// 限流器，最大允许10并发, 0.1s 速率
var globeLimit = rate.NewLimiter(rate.Every(500*time.Millisecond), 10)

func GlobeLimitRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t, _ := model.GetClientIP(ctx.ClientIP())
		if time.Now().Unix() > t {
			ok := model.SetClientIP(ctx.ClientIP())
			if !ok {
				ctx.JSON(resp.ERROR, resp.NewErrorResp())
				ctx.Abort()
				return
			}
		} else {
			ctx.JSON(resp.ERROR, resp.NewErrorResp(resp.WithMessage("访问频率过快，请稍后访问")))
			ctx.Abort()
		}
		globeLimit.Allow()
		ctx.Next()
	}
}

func ApiLimitRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t, _ := model.GetClientIP(ctx.ClientIP())
		if time.Now().Unix() > t {
			ok := model.SetClientIP(ctx.ClientIP())
			if !ok {
				ctx.JSON(resp.ERROR, resp.NewErrorResp())
				ctx.Abort()
				return
			}
		} else {
			ctx.JSON(resp.ERROR, resp.NewErrorResp(resp.WithMessage("访问频率过快，请稍后访问")))
			ctx.Abort()
		}
		globeLimit.Allow()
		ctx.Next()
	}
}
