package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string
	Gender   int
	Email    string
	Avatar   string
}
