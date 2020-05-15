package utils

const (
	SUCCESS = 200
	ERROR   = 500
)

var ErrorMap = map[int]string{
	SUCCESS: "成功",
	ERROR:   "失败",
}

func GetMsg(code int) string {
	msg, ok := ErrorMap[code]
	if ok {
		return msg
	}

	return ErrorMap[ERROR]
}
