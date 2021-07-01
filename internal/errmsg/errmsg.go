package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
	// token 错误
	TOKEN_ERROR        = 600
	TOKEN_NOT_FOUND    = 601
	TOKEN_FORMAT_ERROR = 602
	TOKEN_NOT_VALID    = 603

	//文件上传错误
	FILE_UPLOAD_FAILED   = 801
	FILE_DOWNLOAD_FAILED = 802

	// code= 1000... 账号模块的错误
	ERROR_ACCOUNT_EXIST     = 1001
	ERROR_ACCOUNT_NOT_EXIST = 1002
	ERROR_PASSWORD_ERROR    = 1003
	// code= 2000... 用户模块错误
	ERROR_USER_NOT_EXIST = 2001
	// code= 3000 ... 团队模块
	ERROR_TEAM_NOT_EXIST         = 3001
	ERROR_MEMBER_NOT_IN_TEAM     = 3002
	ERROR_MEMBER_ALREADY_IN_TEAM = 3003
	ERROR_MEMBER_IS_LEADER       = 3004
	ERROR_MEMBER_IS_NOT_LEADER   = 3005
	ERROR_MEMBER_IS_ME           = 3006
	// code= 4000 ...相册模块
	ERROR_ALBUM_NOT_EXIST = 4001
)

var codeMsg = map[int]string{
	SUCCESS:                      "OK",
	ERROR:                        "FAIL",
	TOKEN_ERROR:                  "token错误",
	TOKEN_NOT_FOUND:              "无token",
	TOKEN_FORMAT_ERROR:           "token格式错误",
	TOKEN_NOT_VALID:              "token无效或已过期",
	FILE_UPLOAD_FAILED:           "文件上传失败",
	FILE_DOWNLOAD_FAILED:         "文件下载失败",
	ERROR_ACCOUNT_EXIST:          "账号已存在",
	ERROR_ACCOUNT_NOT_EXIST:      "账号不存在",
	ERROR_PASSWORD_ERROR:         "密码错误",
	ERROR_USER_NOT_EXIST:         "用户不存在",
	ERROR_TEAM_NOT_EXIST:         "团队不存在",
	ERROR_MEMBER_NOT_IN_TEAM:     "成员不在团队内",
	ERROR_MEMBER_ALREADY_IN_TEAM: "成员已经在团队内",
	ERROR_MEMBER_IS_LEADER:       "队长需先转让或解散团队",
	ERROR_MEMBER_IS_NOT_LEADER:   "不是队长，无权限操作",
	ERROR_MEMBER_IS_ME:           "不能对自己操作",
	ERROR_ALBUM_NOT_EXIST:        "相册不存在",
}

//通过错误码获取错误信息
func GetErrMsg(code int) string {
	return codeMsg[code]
}
