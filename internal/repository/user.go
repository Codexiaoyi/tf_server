package repository

import (
	"tfserver/internal/model"

	"gorm.io/gorm"
)

//查询用户
func QueryUserByEmail(email string) (model.User, error) {
	var user model.User
	err := db.Limit(1).Where("email = ?", email).Find(&user).Error
	return user, err
}

//更新用户信息
func UpdateUser(id int, user *model.User) error {
	return BeginTransaction(db, func(tx *gorm.DB) error {
		return tx.Model(&user).Where("ID = ?", user.ID).Updates(map[string]interface{}{
			"user_name": user.UserName,
			"gender":    user.Gender,
			"email":     user.Email,
			"avatar":    user.Avatar,
			"year":      user.Year,
			"month":     user.Month,
			"day":       user.Day,
			"province":  user.Province,
			"city":      user.City,
			"street":    user.Street,
		}).Error
	})
}
