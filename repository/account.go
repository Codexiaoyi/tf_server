package repository

import "tfserver/model"

//增加新账户
func AddAccount(account *model.Account) error {
	err := db.Create(&account).Error
	return err
}

//查询账户密码
func QueryPasswordByEmail(email string) (string, error) {
	var account model.Account
	err := db.Limit(1).Select("password").Where("email = ?", email).Take(&account).Error
	return account.Password, err
}

//检查账户是否存在
func CheckAccount(email string) bool {
	var account model.Account
	db.Select("id").Where("email = ?", email).First(&account)
	return account.ID > 0
}
