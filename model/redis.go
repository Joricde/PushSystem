package model

import (
	"PushSystem/config"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"math/rand"
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

func SetUserDynamicKey(userID uint) (int, error) {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000 // 六位数动态口令
	t := time.Second * 1000
	//TODO 部署时修改时间
	key := config.RedisUserID + ":dynamicKey:" + strconv.Itoa(int(userID))
	_, err := RedisDB.Set(c, key, code, t).Result()
	return code, err
}

func GetUserDynamicKey(userID uint) (int, error) {
	key := config.RedisUserID + ":dynamicKey:" + strconv.Itoa(int(userID))
	r, err := RedisDB.Get(c, key).Result()
	if err != nil {
		zap.L().Error("2" + err.Error())
		return 0, err
	}
	result, err := strconv.Atoi(r)
	return result, err
}
