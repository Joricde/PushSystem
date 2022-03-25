package model

import (
	"encoding/json"
	"strconv"
	"time"
)

const ExpTime = time.Minute * 30

func SetRedisUser(user *User) (string, error) {
	key := "userID" + strconv.Itoa(int(user.ID))
	val, _ := json.Marshal(user)
	result, err := RedisDB.Set(c, key, val, ExpTime).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func GetRedisUserByID(id int) (*User, error) {
	key := "userID" + strconv.Itoa(id)
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
