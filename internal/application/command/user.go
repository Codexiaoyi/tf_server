package command

import "tfserver/internal/model"

type UpdateUserInfo struct {
	UserName string
	Gender   int
	Avatar   string
	model.Birthday
	model.Address
}
