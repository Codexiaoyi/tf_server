package repository

import (
	"tfserver/model"
)

//新增用户
func AddUser(data *model.User) bool {
	err := db.Create(&data).Error
	return err == nil
}

//查询用户
func QueryUserByEmail(email string) (model.User, bool) {
	var user model.User
	err := db.Limit(1).Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, false
	}
	return user, true
}

//更新用户信息
func UpdateUser(id int, data *model.User) bool {
	var user model.User

	err = db.Model(&user).Updates(map[string]interface{}{
		"username": data.UserName,
		"gender":   data.Gender,
		"email":    data.Email,
		"avatar":   data.Avatar,
	}).Error

	if err != nil {
		return false
	}

	return true
}
