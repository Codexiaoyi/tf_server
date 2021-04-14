package v1

import (
	"tfserver/application/command"
	"tfserver/application/query"
	"tfserver/model"
	"tfserver/repository"
	"tfserver/util/errmsg"
	"tfserver/util/response"

	"github.com/Codexiaoyi/go-mapper"
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

//获取用户相册信息
func GetUserAlbumInfo(c *gin.Context) {
	var query query.GetUserAlbumInfo
	c.ShouldBindJSON(&query)

	album, err := repository.QueryUserAlbumByAlbumId(query.AlbumId)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	if album.ID <= 0 {
		response.Response(c, errmsg.ERROR_ALBUM_NOT_EXIST)
		return
	}

	data := make(map[string]interface{})
	data["album"] = album
	response.ResponseWithData(c, errmsg.SUCCESS, data)
}

//创建新相册
func CreateNewUserAlbum(c *gin.Context) {
	email := c.GetString("email")
	var command command.CreateAlbum
	c.ShouldBindJSON(&command)

	var album model.UserAlbum
	err := mapper.StructMapByFieldName(&command, &album)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	album.Email = email
	err = repository.CreateNewUserAlbum(&album)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	response.Response(c, errmsg.SUCCESS)
}

//修改相册信息
func UpdateUserAlbumInfo(c *gin.Context) {
	var command command.UpdateAlbumInfo
	c.ShouldBindJSON(&command)

	//先查询相册是否存在
	album, err := repository.QueryUserAlbumByAlbumId(command.AlbumId)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	if album.ID <= 0 {
		response.Response(c, errmsg.ERROR_ALBUM_NOT_EXIST)
		return
	}

	err = mapper.StructMapByFieldName(&command, &album)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	err = repository.UpdateUserAlbumInfo(int(album.ID), &album)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	response.Response(c, errmsg.SUCCESS)
}
