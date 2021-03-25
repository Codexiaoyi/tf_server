package repository

import (
	"tfserver/model"
)

//新增用户
func CreateUser(data *model.User) bool {
	err := db.Create(&data).Error
	return err == nil
}

// 查询用户
func GetUser(id int) (model.User, bool) {
	var user model.User
	err := db.Limit(1).Where("ID = ?", id).Find(&user).Error
	if err != nil {
		return user, false
	}
	return user, true
}
