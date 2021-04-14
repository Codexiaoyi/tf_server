package repository

import (
	"tfserver/model"

	"gorm.io/gorm"
)

//查询用户相册是否存在
func CheckAlbum(albumId int) bool {
	var album model.UserAlbum
	db.Select("id").Where("ID = ?", albumId).First(&album)
	return album.ID > 0
}

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
