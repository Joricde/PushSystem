package model

import (
	"PushSystem/config"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func (u User) SetRedisUser(user *User) (string, error) {
	key := config.RedisUserID + strconv.Itoa(int(user.ID))
	val, _ := json.Marshal(user)
	result, err := RedisDB.Set(c, key, val, config.ExpTime).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (u User) GetRedisUserByID(id uint) (*User, error) {
	key := config.RedisUserID + strconv.Itoa(int(id))
	jsUser, err := RedisDB.Get(c, key).Result()
	if err != nil {
		return nil, err
	}
	user := new(User)
	err = json.Unmarshal([]byte(jsUser), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func SetClientIP(key string) error {
	t := time.Now().Add(time.Second).Unix()
	result, err := RedisDB.Set(c, key, t, config.ExpTime).Result()
	if err != nil {
		zap.L().Error(result + err.Error())
	}
	return err
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
