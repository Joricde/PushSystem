package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"
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

	fmt.Println(body)
}
