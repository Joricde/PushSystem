package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type qrResp struct {
	code    int
	message string
	data    struct {
		ticket        string
		expireSeconds int
		qrUrl         string
		token         int
	}
}

func getQRTicket() *qrResp {
	resp, _ := http.Get("https://sctapi.ftqq.com/user/signin")
	body, _ := io.ReadAll(resp.Body)
	r := qrResp{}
	_ = json.Unmarshal(body, &r)
	return &r
}

func GetQRLogin() {
	json := a
	payload := strings.NewReader(json)
	resp, _ := http.NewRequest(http.MethodPost,
		"https://iam.cn-east-3.myhuaweicloud.com/v3/auth/tokens",
		payload)

	body, _ := io.ReadAll(resp.Body)
	println(body)
}
