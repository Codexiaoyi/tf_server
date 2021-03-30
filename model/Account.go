package model

import "gorm.io/gorm"

//用户账户模型，Role用于角色管理
type Account struct {
	gorm.Model
	Email    string `gorm:"type:varchar(30)" json:"email"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}
