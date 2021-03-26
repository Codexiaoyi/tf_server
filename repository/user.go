package repository

import (
	"tfserver/model"
)

//新增用户
func AddUser(data *model.User) error {
	err := db.Create(&data).Error
	return err
}

//查询用户
func QueryUserByEmail(email string) (model.User, error) {
	var user model.User
	err := db.Limit(1).Where("email = ?", email).Find(&user).Error
	return user, err
}

//更新用户信息
func UpdateUser(id int, data *model.User) error {
	var user model.User

	err = db.Model(&user).Where("ID = ?", user.ID).Updates(map[string]interface{}{
		"username": data.UserName,
		"gender":   data.Gender,
		"email":    data.Email,
		"avatar":   data.Avatar,
	}).Error

	return err
}
