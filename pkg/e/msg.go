package e

var MsgFlags = map[int]string{
	Success:                "ok",
	Error:                  "fail",
	InvalidParams:          "参数错误",
	ErrorExistUser:         "用户已存在",
	ErrorFailEncryption:    "加密失败",
	ErrorExistUserNotFound: "用户不存在",
	ErrorNotCompare:        "密码不匹配",
	ErrorAuthToken:         "token 认证失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
