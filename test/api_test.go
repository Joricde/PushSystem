package test

import (
	"io"
	"net/http"
	"strings"
)

var a = `
{
"auth": {
"identity": {
"methods": [
"password"
],
"password": {
"user": {
"domain": {
"name": "hid_qxncpzi1-p7u3-e"
},
"name": "HuaweiIAM"
}
}
},
"scope": {
"project": {
"id": "c60dfa0a6b074691b7db27c46d810222"
}
}
}
}`

func GetSMNToken() {
	json := a
	payload := strings.NewReader(json)
	req, _ := http.NewRequest(http.MethodPost,
		"https://iam.cn-east-3.myhuaweicloud.com/v3/auth/tokens",
		payload)

	// Set the auth for the request.
	req.SetBasicAuth("admin", "Admin@123")
	body, _ := io.ReadAll(req.Body)
	println(body)
}

func main() {
	GetSMNToken()
}
