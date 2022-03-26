package resp

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400
)

var Flags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	InvalidParams: "请求参数错误",
}

func GetMessage(code int) string {
	msg, ok := Flags[code]
	if ok {
		return msg
	}
	return Flags[ERROR]
}
