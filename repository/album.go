package repository

import (
	"tfserver/model"

	"gorm.io/gorm"
)

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
func UpdateUserAlbumInfo(album *model.UserAlbum) error {
	return BeginTransaction(db, func(tx *gorm.DB) error {
		return tx.Model(&album).Where("ID = ?", album.ID).Updates(map[string]interface{}{
			"team_name":    album.Public,
			"avatar":       album.Name,
			"introduction": album.Introduction,
			"cover":        album.Cover,
			"like":         album.Like,
			"collect":      album.Collect,
		}).Error
	})
}
