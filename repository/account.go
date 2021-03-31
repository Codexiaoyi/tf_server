package repository

import (
	"tfserver/model"

	"gorm.io/gorm"
)

//查询账户密码
func QueryAccountByEmail(email string) (model.Account, error) {
	var account model.Account
	err := db.Limit(1).Where("email = ?", email).Find(&account).Error
	return account, err
}

//检查账户是否存在
func CheckAccount(email string) bool {
	var account model.Account
	db.Select("id").Where("email = ?", email).First(&account)
	return account.ID > 0
}

//增加新账户
func AddAccount(account *model.Account) error {
	return BeginTransaction(db, func(tx *gorm.DB) error {
		var err error
		err = tx.Create(&account).Error
		if err != nil {
			return err
		}
		//添加默认账户信息
		err = tx.Create(&model.User{
			Email: account.Email,
		}).Error
		return err
	})
}

//修改账号信息，目前仅改密码
func UpdateAccount(account *model.Account) error {
	return BeginTransaction(db, func(tx *gorm.DB) error {
		err := db.Model(&account).Where("ID = ?", account.ID).Updates(map[string]interface{}{
			"password": account.Password,
		}).Error
		return err
	})
}
