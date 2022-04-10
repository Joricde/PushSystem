package model

import (
	"PushSystem/config"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func (u User) SetRedisUser(user *User) bool {
	key := config.RedisUserID + strconv.Itoa(int(user.ID))
	val, _ := json.Marshal(user)
	_, err := RedisDB.Set(c, key, val, config.ExpTime).Result()
	if err != nil {
		zap.L().Error(err.Error())
		return false
	}
	return true
}

func (u User) GetRedisUserByID(id uint) *User {
	key := config.RedisUserID + strconv.Itoa(int(id))
	jsUser, err := RedisDB.Get(c, key).Result()
	user := new(User)
	switch {
	case err == redis.Nil:
		user.ID = 0
	case err != nil:
		zap.L().Error(err.Error())
		user.ID = -1
	default:
		_ = json.Unmarshal([]byte(jsUser), user)
	}
	return user
}

func SetClientIP(key string) bool {
	t := time.Now().Add(time.Second).Unix()
	result, err := RedisDB.Set(c, key, t, config.ExpTime).Result()
	if err != nil {
		zap.L().Error(result + err.Error())
		return false
	}
	return true
}

func GetClientIP(key string) (int64, error) {
	r, err := RedisDB.Get(c, key).Result()
	if err == redis.Nil {
		return -1, nil
	} else if err != nil {
		zap.L().Error(r)
	}
	result, _ := strconv.ParseInt(r, 10, 64)
	return result, err
}
