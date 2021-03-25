package repository

import "tfserver/model"

//增加新账户
func AddAccount(account *model.Account) bool {
	err := db.Create(&account).Error
	return err == nil
}

//查询账户
func QueryPasswordByEmail(email string) (model.Account, bool) {
	var account model.Account
	err := db.Limit(1).Select("password").Where("email = ?", email).Find(&account)
	if err != nil {
		return account, false
	}
	return account, true
}
