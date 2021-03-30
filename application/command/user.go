package command

import "tfserver/model"

type UpdateUserInfo struct {
	UserName string
	Gender   int
	Email    string
	Avatar   string
	model.Birthday
	model.Address
}
