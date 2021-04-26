package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"tfserver/application/command"
	"tfserver/application/query"
	"tfserver/config"
	"tfserver/model"
	"tfserver/repository"
	"tfserver/repository/cache"
	"tfserver/util/errmsg"
	"tfserver/util/oss"
	"tfserver/util/response"
	"time"

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

//获取自己用户相册内的所有缩略图列表
func GetUserAlbumMedias(c *gin.Context) {
	var query query.GetUserAlbumMedias
	c.ShouldBindJSON(&query)

	//先查询相册是否存在
	album, err := repository.QueryUserAlbumByAlbumId(query.AlbumId)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	if album.ID <= 0 {
		response.Response(c, errmsg.ERROR_ALBUM_NOT_EXIST)
		return
	}

	medias, err := repository.QueryUserAlbumMedias(query.AlbumId)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	data := make(map[string]interface{})
	data["medias"] = medias
	response.ResponseWithData(c, errmsg.SUCCESS, data)
}

//获取缩略图
func GetThumbnail(c *gin.Context) {
	var query query.GetThumbnail
	c.ShouldBindJSON(&query)

	exist := cache.CDb.IsExist("album_thumbnail", query.Url)
	if exist {
		//缓存有
		file, err := cache.CDb.Get("album_thumbnail", query.Url)
		if err == nil {
			c.Writer.Write(file)
			return
		}
	}

	if query.Url == "" {
		c.Status(http.StatusNoContent)
		return
	}

	file, err := oss.Download(query.Url)
	if err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	cache.CDb.Set("album_thumbnail", query.Url, file)

	c.Writer.Write(file)
}

//指定用户相册封面
func SetUserAlbumCover(c *gin.Context) {
	var command command.SetUserAlbumCover
	c.ShouldBindJSON(&command)

	err := repository.SetUserAlbumCover(command.AlbumId, command.MediaId)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	response.Response(c, errmsg.SUCCESS)
}

//获取文件上传凭证及路径，用于前端直接上传到oss
func GetUploadKey(c *gin.Context) {
	email := c.GetString("email")
	var command command.GetKeyAndUrl
	c.ShouldBindJSON(&command)

	split := strings.Split(command.Name, ".")
	ext := split[len(split)-1] //获取扩展名
	//获取当前时间戳
	timeUnix := strconv.FormatInt(time.Now().Unix(), 10)
	url := fmt.Sprintf("user/%s/album/media/%s.%s", email, timeUnix, ext)
	data := make(map[string]interface{})
	data["region"] = config.OssEndpointToWeb
	data["accessKeyId"] = config.OssAccessKeyIdToWeb
	data["accessKeySecret"] = config.OssAccessKeySecretToWeb
	data["bucket"] = config.OssBucketNameToWeb
	data["url"] = url //所有的url
	response.ResponseWithData(c, errmsg.SUCCESS, data)
}

//文件上传成功后回调
func UploadMedia(c *gin.Context) {
	email := c.GetString("email")
	var command command.UploadSuccess
	c.ShouldBindJSON(&command)

	split := strings.Split(command.Url, "/")
	name := split[len(split)-1]
	thumbnailUrl := fmt.Sprintf("user/%s/album/media/thumbnail/%s", email, name)

	media := model.UserAlbumMedia{
		AlbumId:      command.AlbumId,
		IsVideo:      command.IsVideo,
		Url:          command.Url,
		ThumbnailUrl: thumbnailUrl,
	}

	//写入数据库
	err := repository.AddImage(&media)
	if err != nil {
		response.Response(c, errmsg.ERROR)
		return
	}

	response.Response(c, errmsg.SUCCESS)
}
