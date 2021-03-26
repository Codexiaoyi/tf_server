package v1

import (
	"tfserver/application/command"
	"tfserver/application/query"
	"tfserver/repository"
	"tfserver/util/errmsg"
	"tfserver/util/response"

	"github.com/gin-gonic/gin"
)

//获取用户信息
func GetUserInfo(c *gin.Context) {
	var query query.GetUserInfo
	_ = c.ShouldBindJSON(&query)
	status := errmsg.ERROR

	user, err := repository.QueryUserByEmail(query.Email)
	if err == nil {
		data := make(map[string]interface{})
		data["user"] = user
		response.ResponseWithData(c, status, data)
		return
	}

	response.Response(c, status)
}

//更新用户信息
func UpdateUserInfo(c *gin.Context) {
	var command command.UpdateUserInfo
	_ = c.ShouldBindJSON(&command)
	status := errmsg.ERROR

	//先查询用户是否存在
	user, err := repository.QueryUserByEmail(command.Email)
	if err == nil && user.ID > 0 {
		updateErr := repository.UpdateUser(int(user.ID), &user)
		if updateErr == nil {
			//更新成功
			status = errmsg.SUCCESS
		}
	} else {
		//用户不存在，不更新
		status = errmsg.ERROR_USER_NOT_EXIST
	}

	response.Response(c, status)
}
