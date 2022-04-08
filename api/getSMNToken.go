package api

import (
	"io"
	"net/http"
	"strings"
)

var a = ``

func GetSMNToken() {
	json := a
	payload := strings.NewReader(json)
	req, _ := http.NewRequest(http.MethodPost,
		"https://iam.cn-east-3.myhuaweicloud.com/v3/auth/tokens",
		payload)

	// Set the auth for the request.
	body, _ := io.ReadAll(req.Body)
	println(body)
}
