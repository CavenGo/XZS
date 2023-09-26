package common

const (
	Ok                  MyCode = 1
	AccessTokenError    MyCode = 400
	UNAUTHORIZED        MyCode = 401
	AuthError           MyCode = 402
	InnerError          MyCode = 500
	ParameterValidError MyCode = 501
	AccessDenied        MyCode = 502
)

var MsgFlags = map[MyCode]string{
	Ok:                  "成功",
	AccessTokenError:    "用户登录令牌失效",
	UNAUTHORIZED:        "用户未登录",
	AuthError:           "用户名或密码错误",
	InnerError:          "系统内部错误",
	ParameterValidError: "参数验证错误",
	AccessDenied:        "用户没有权限访问",
}

func (m MyCode) Msg() string {
	msg, ok := MsgFlags[m]
	if ok {
		return msg
	}
	return MsgFlags[InnerError]
}
