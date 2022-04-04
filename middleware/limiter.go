package middleware

import (
	"PushSystem/model"
	"PushSystem/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"time"
)

// 限流器，最大允许1000并发, 0.5s 速率
var globeLimit = rate.NewLimiter(rate.Every(500*time.Millisecond), 1000)

func GlobeLimitRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t, _ := model.GetClientIP(ctx.ClientIP())
		if time.Now().Unix() > t {
			err := model.SetClientIP(ctx.ClientIP())
			if err != nil {
				zap.L().Error(err.Error())
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
			err := model.SetClientIP(ctx.ClientIP())
			if err != nil {
				zap.L().Error(err.Error())
			}
		} else {
			ctx.JSON(resp.ERROR, resp.NewErrorResp(resp.WithMessage("访问频率过快，请稍后访问")))
			ctx.Abort()
		}
		globeLimit.Allow()
		ctx.Next()
	}
}
