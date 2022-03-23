package model

import (
	"PushSystem/config"
	"PushSystem/util"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var RedisDB *redis.Client
var c = context.Background()

func init() {
	cfg := config.Conf.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	DB = db
	migration()
	connectRedis()
}

func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{})
	if err != nil {
		return
	}
}

func connectRedis() {
	cfg := config.Conf.Redis
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     cfg.Address,  // use default Addr
		Password: cfg.Password, // no password set
		DB:       0,            // use default DB
	})
	result, err := RedisDB.Ping(c).Result()
	if err != nil {
		util.Logger.Error(err.Error())
	}
	util.Logger.Debug("redis ctx: " + result)
	fmt.Println("redis ctx: " + result)
}
