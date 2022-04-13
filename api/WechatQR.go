package api

import (
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type QRResp struct {
	Data struct {
		Ticket        string `json:"ticket"`
		ExpireSeconds int    `json:"expire_seconds"`
		QrUrl         string `json:"qr_url"`
		Token         string `json:"token"`
	}
}

type CheckQRResp struct {
	Uid     int64
	SendKey string
}

func GetWechatQR() *QRResp {
	resp, _ := http.Get("https://sctapi.ftqq.com/user/signin")
	body, _ := io.ReadAll(resp.Body)
	r := new(QRResp)
	err := json.Unmarshal(body, r)
	if err != nil {
		zap.L().Error(err.Error())
	}
	zap.L().Debug(r.Data.QrUrl)
	return r
}

func CheckQRLogin(token string) *CheckQRResp {
	type checkQR struct {
		Code    int
		Message string
		Data    interface{}
	}
	type getCheckQR struct {
		Result struct {
			Uid     string
			SendKey string
		}
	}
	r, _ := http.PostForm("https://sctapi.ftqq.com/sso/check",
		url.Values{"token": {token}})
	body, _ := io.ReadAll(r.Body)
	c := new(checkQR)
	err := json.Unmarshal(body, c)
	if err != nil {
		zap.L().Error(err.Error())
		return new(CheckQRResp)
	}
	midData, ok := c.Data.(map[string]interface{})
	if ok {
		marshal, err := json.Marshal(midData)
		if err != nil {
			zap.L().Error(err.Error())
		}
		midResp := new(getCheckQR)
		err = json.Unmarshal(marshal, midResp)
		if err != nil {
			zap.L().Error(err.Error())
		}
		uid, err := strconv.ParseInt(midResp.Result.Uid, 10, 64)
		if err != nil {
			zap.L().Error(err.Error())
		}
		return &CheckQRResp{Uid: uid,
			SendKey: midResp.Result.SendKey}
	} else {
		return new(CheckQRResp)
	}
}

func GetHWLogin() {
	json := a
	payload := strings.NewReader(json)
	resp, _ := http.NewRequest(http.MethodPost,
		"https://iam.cn-east-3.myhuaweicloud.com/v3/auth/tokens",
		payload)
	body, _ := io.ReadAll(resp.Body)
	println(body)
}
