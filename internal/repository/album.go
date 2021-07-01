package repository

import (
	"errors"
	"tfserver/internal/model"

	"gorm.io/gorm"
)

//获取用户相册信息
func QueryUserAlbumByAlbumId(albumId int) (model.UserAlbum, error) {
	var album model.UserAlbum
	err := db.Limit(1).Where("ID = ?", albumId).Find(&album).Error
	return album, err
}

//根据email查询所有相册
func QueryUserAlbumsByEmail(email string) ([]model.UserAlbum, error) {
	albums := make([]model.UserAlbum, 0)
	err := db.Where("email = ?", email).Find(&albums).Error
	return albums, err
}

//查询用户相册内媒体数组
func QueryUserAlbumMedias(albumId int) ([]model.UserAlbumMedia, error) {
	medias := make([]model.UserAlbumMedia, 0)
	err := db.Where("album_id = ?", albumId).Find(&medias).Error
	return medias, err
}

//新建相册
func CreateNewUserAlbum(album *model.UserAlbum) error {
	return BeginTransaction(db, func(tx *gorm.DB) error {
		return tx.Create(&album).Error
	})
}

//修改相册信息
func UpdateUserAlbumInfo(albumId int, album *model.UserAlbum) error {
	return BeginTransaction(db, func(tx *gorm.DB) error {
		return tx.Model(&album).Where("ID = ?", albumId).Updates(map[string]interface{}{
			"public":       album.Public,
			"name":         album.Name,
			"introduction": album.Introduction,
			"cover":        album.Cover,
			"like":         album.Like,
			"collect":      album.Collect,
		}).Error
	})
}

//指定相册封面（将已有的文件）
func SetUserAlbumCover(albumId, mediaId int) error {
	return BeginTransaction(db, func(tx *gorm.DB) error {
		var media model.UserAlbumMedia
		err := tx.Limit(1).Select("album_id", "thumbnail_url").Where("ID = ?", mediaId).Find(&media).Error
		if err != nil || media.AlbumId != albumId {
			return errors.New("not exist")
		}

		return tx.Model(&model.UserAlbum{}).Where("ID = ?", albumId).Update("cover", media.ThumbnailUrl).Error
	})
}

//新建图片文件到相册
func AddImage(medias *model.UserAlbumMedia) error {
	return BeginTransaction(db, func(tx *gorm.DB) error {
		return tx.Create(&medias).Error
	})
}
