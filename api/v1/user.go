package v1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"tfserver/internal/application/command"
	"tfserver/internal/errmsg"
	"tfserver/internal/repository"
	"tfserver/internal/repository/cache"
	"tfserver/internal/response"
	"tfserver/pkg/oss"
	"time"

	"github.com/Codexiaoyi/go-mapper"
	"github.com/gin-gonic/gin"
)

//获取用户信息
func GetUserInfo(c *gin.Context) {
	email := c.GetString("email")

	user, err := repository.QueryUserByEmail(email)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	if user.Email != "" && user.Email == email {
		data := make(map[string]interface{})
		data["user"] = user
		response.ResponseWithData(c, errmsg.SUCCESS, data)
	} else {
		response.Response(c, errmsg.ERROR_USER_NOT_EXIST)
	}

}

//更新用户信息
func UpdateUserInfo(c *gin.Context) {
	var command command.UpdateUserInfo
	_ = c.ShouldBindJSON(&command)
	email := c.GetString("email")

	//先查询用户是否存在
	user, err := repository.QueryUserByEmail(email)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}
	if user.ID > 0 {
		//用户存在
		mapErr := mapper.StructMapByFieldName(&command, &user)
		if mapErr != nil {
			response.Response(c, errmsg.ERROR)
			return
		}
		updateErr := repository.UpdateUser(int(user.ID), &user)
		if updateErr != nil {
			response.Response(c, errmsg.ERROR)
			return
		}
		response.Response(c, errmsg.SUCCESS)
	} else {
		//用户不存在，不更新
		response.Response(c, errmsg.ERROR_USER_NOT_EXIST)
	}
}

//上传用户头像
func UploadUserAvatar(c *gin.Context) {
	email := c.GetString("email")
	field := fmt.Sprintf("avatar_%s", email)

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

	exist := cache.CDb.IsExist("user", field)
	if exist {
		//缓存失效
		cache.CDb.Delete("user", field)
	}

	response.Response(c, errmsg.SUCCESS)
}

//加载自己头像
func GetUserAvatar(c *gin.Context) {
	email := c.GetString("email")
	field := fmt.Sprintf("avatar_%s", email)

	exist := cache.CDb.IsExist("user", field)
	if exist {
		//缓存有
		file, err := cache.CDb.Get("user", field)
		if err == nil {
			c.Writer.Write(file)
			return
		}
	}

	//查询头像地址
	user, err := repository.QueryUserByEmail(email)
	if err != nil || user.ID <= 0 || user.Avatar == "" {
		c.Status(http.StatusNoContent)
		return
	}

	file, err := oss.Download(user.Avatar)
	if err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	cache.CDb.Set("user", field, file)

	c.Writer.Write(file)
}
