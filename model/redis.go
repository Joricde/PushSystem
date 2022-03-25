package model

import (
	"PushSystem/config"
	"encoding/json"
	"strconv"
)

func SetRedisUser(user *User) (string, error) {
	key := config.RedisUserID + strconv.Itoa(int(user.ID))
	val, _ := json.Marshal(user)
	result, err := RedisDB.Set(c, key, val, config.ExpTime).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func GetRedisUserByID(id uint) (*User, error) {
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
