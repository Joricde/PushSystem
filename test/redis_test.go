package test

import (
	"PushSystem/config"
	"PushSystem/util"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client
var c = context.Background()

func connectRedis() {
	cfg := config.Conf.Redis
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     cfg.Address,  // use default Addr
		Password: cfg.Password, // no password set
		DB:       0,            // use default DB
	})
	result, err := RedisDB.Ping(c).Result()
	if err != nil {
		return
	}
	util.Logger.Debug("redis ctx: " + result)
	fmt.Println("redis ctx: " + result)
}
