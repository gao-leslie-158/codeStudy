package controller

type ResCode int64

const (
	CodeSuccess = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPssword
	CodeServerBusy
	CodeNeedLogin
	CodeInvalidToken
)

var CodeMsgMap = map[ResCode]string{
	CodeSuccess:        "success",
	CodeInvalidParam:   "请求参数错误",
	CodeUserExist:      "用户已经存在",
	CodeUserNotExist:   "用户不存在",
	CodeInvalidPssword: "用户名或密码错误",
	CodeServerBusy:     "服务繁忙",
	CodeNeedLogin:      "需要登录",
	CodeInvalidToken:   "无效的Token",
}

func (c ResCode) Msg() string {
	msg, ok := CodeMsgMap[c]
	if !ok {
		msg = CodeMsgMap[CodeServerBusy]
	}
	return msg
}
