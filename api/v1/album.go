package v1

import (
	"tfserver/repository"
	"tfserver/util/errmsg"
	"tfserver/util/response"

	"github.com/gin-gonic/gin"
)

//获取用户所有相册
func GetUserAlbums(c *gin.Context) {
	email := c.GetString("email")

	albums, err := repository.QueryUserAlbumsByEmail(email)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	data := make(map[string]interface{})
	data["albums"] = albums
	response.ResponseWithData(c, errmsg.SUCCESS, data)
}

//创建新相册
func CreateNewUserAlbum(c *gin.Context) {

}

//修改相册信息
func UpdateUserAlbumInfo(c *gin.Context) {

}

//删除相册
func DeleteUserAlbum(c *gin.Context) {

}
