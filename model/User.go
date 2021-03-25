package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(25)" json:"username"`
	Gender   int    `gorm:"type:int" json:"gender"`
	Email    string `gorm:"type:varchar(30)" json:"email"`
	Avatar   string `gorm:"type:varchar(100)" json:"avatar"`
}
