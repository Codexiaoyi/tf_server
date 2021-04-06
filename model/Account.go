package model

import "gorm.io/gorm"

//用户账户模型，Role用于角色管理
type Account struct {
	gorm.Model
	Email    string `gorm:"type:varchar(30);not null;default:''" json:"email"`
	Password string `gorm:"type:varchar(100);not null;default:''" json:"password"`
	Role     int    `gorm:"type:int;not null" json:"role"`
}
