package v1

import (
	"io/ioutil"
	"path"
	"strconv"
	"tfserver/application/command"
	"tfserver/repository"
	"tfserver/util/errmsg"
	"tfserver/util/log"
	"tfserver/util/oss"
	"tfserver/util/response"
	"time"

	"github.com/Codexiaoyi/go-mapper"
	"github.com/gin-gonic/gin"
)

//获取用户信息
func GetUserInfo(c *gin.Context) {
	email := c.GetString("email")
	status := errmsg.ERROR

	user, err := repository.QueryUserByEmail(email)
	if err == nil {
		if user.Email != "" && user.Email == email {
			data := make(map[string]interface{})
			data["user"] = user
			status = errmsg.SUCCESS
			response.ResponseWithData(c, status, data)
			return
		} else {
			status = errmsg.ERROR_USER_NOT_EXIST
		}
	}

	response.Response(c, status)
}

//更新用户信息
func UpdateUserInfo(c *gin.Context) {
	var command command.UpdateUserInfo
	_ = c.ShouldBindJSON(&command)
	email := c.GetString("email")
	status := errmsg.ERROR

	//先查询用户是否存在
	user, err := repository.QueryUserByEmail(email)
	if err == nil && user.ID > 0 {
		//用户存在
		mapErr := mapper.StructMapByFieldName(&command, &user)
		if mapErr == nil {
			updateErr := repository.UpdateUser(int(user.ID), &user)
			if updateErr == nil {
				//更新成功
				status = errmsg.SUCCESS
			}
		} else {
			log.ErrorLog("Dto map failed!", mapErr.Error())
		}
	} else {
		//用户不存在，不更新
		status = errmsg.ERROR_USER_NOT_EXIST
	}

	response.Response(c, status)
}

//上传用户头像
func UploadUserAvatar(c *gin.Context) {
	email := c.GetString("email")
	file, err := c.FormFile("avatar")
	if err != nil {
		response.Response(c, errmsg.FILE_UPLOAD_FAILED)
		return
	}

	src, err := file.Open()
	if err != nil {
		response.Response(c, errmsg.FILE_UPLOAD_FAILED)
		return
	}
	defer src.Close()

	//获取当前时间戳
	timeUnix := strconv.FormatInt(time.Now().Unix(), 10)
	//组成新文件名
	path := "user/" + email + "/avatar/" + timeUnix + path.Ext(file.Filename)

	//先保存地址到数据库
	user, err := repository.QueryUserByEmail(email)
	if err != nil || user.ID <= 0 {
		response.Response(c, errmsg.ERROR_USER_NOT_EXIST)
		return
	}
	//用户存在
	user.Avatar = path
	updateErr := repository.UpdateUser(int(user.ID), &user)
	if updateErr != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	//获取byte数组
	out, _ := ioutil.ReadAll(src)
	err = oss.Upload(out, path)
	if err != nil {
		response.Response(c, errmsg.FILE_UPLOAD_FAILED)
		return
	}

	response.Response(c, errmsg.SUCCESS)
}
