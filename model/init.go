package model

import (
	"PushSystem/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

var RedisDB *redis.Client
var c = context.Background()

func init() {
	cfg := config.Conf.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	DB = db
	migration()
	connectRedis()
}

func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}, &Task{}, &UserPwd{})
	if err != nil {
		zap.L().Error(err.Error())
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
		zap.L().Error(err.Error())
	}
	zap.L().Debug("redis ctx: " + result)
	fmt.Println("redis ctx: " + result)
}
