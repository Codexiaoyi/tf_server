package command

import "tfserver/model"

type UpdateUserInfo struct {
	UserName string
	Gender   int
	Avatar   string
	model.Birthday
	model.Address
}
