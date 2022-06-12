package api

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
)

func SendWechat(wechatKey, title, context string) bool {
	type Re struct {
		Code    int
		Message string
		Data    struct {
			PushID  string
			ReadKey string
			Error   string
			Errno   int
		}
	}
	r, _ := http.PostForm(fmt.Sprintf("https://sctapi.ftqq.com/%s.send", wechatKey),
		url.Values{"title": {title},
			"desp": {context}})
	body, _ := io.ReadAll(r.Body)
	c := new(Re)
	err := json.Unmarshal(body, c)
	if err != nil {
		zap.L().Error(err.Error())
		return false
	}
	zap.L().Debug(fmt.Sprintln(c))
	if len(c.Data.PushID) > 0 {
		return true
	} else {
		return false
	}

}
