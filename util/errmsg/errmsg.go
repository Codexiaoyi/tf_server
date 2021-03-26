package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code= 1000... 账号模块的错误
	ERROR_ACCOUNT_EXIST     = 1001
	ERROR_ACCOUNT_NOT_EXIST = 1002
	ERROR_PASSWORD_ERROR    = 1003
	// code= 2000... 用户模块错误
	ERROR_USER_NOT_EXIST = 2001
)

var codeMsg = map[int]string{
	SUCCESS:                 "OK",
	ERROR:                   "FAIL",
	ERROR_ACCOUNT_EXIST:     "账号已存在",
	ERROR_ACCOUNT_NOT_EXIST: "账号不存在",
	ERROR_PASSWORD_ERROR:    "密码错误",
	ERROR_USER_NOT_EXIST:    "用户不存在",
}

//通过错误码获取错误信息
func GetErrMsg(code int) string {
	return codeMsg[code]
}
