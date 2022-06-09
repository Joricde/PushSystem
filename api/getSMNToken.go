package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var a = ``

func GetSMNToken() {
	json := a
	payload := strings.NewReader(json)
	req, _ := http.NewRequest(http.MethodPost,
		"***/v3/auth/tokens",
		payload)

	// Set the auth for the request.
	body, _ := io.ReadAll(req.Body)
	time.Now().Format("2006-01-02:15:04")
	fmt.Println(body)
}
