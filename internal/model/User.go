package model

import "gorm.io/gorm"

//用户模型，保存用户基本信息
//Email字段创建索引
type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(25);not null;default:''" json:"username"`
	Gender   int    `gorm:"type:int;not null" json:"gender"`
	Email    string `gorm:"type:varchar(30);not null;default:'';index" json:"email"`
	Avatar   string `gorm:"type:varchar(100);not null;default:''" json:"avatar"`
	Birthday
	Address
}

//生日
type Birthday struct {
	Year  int `gorm:"type:int;not null" json:"year"`
	Month int `gorm:"type:int;not null" json:"month"`
	Day   int `gorm:"type:int;not null" json:"day"`
}

//住址
type Address struct {
	Province string `gorm:"type:varchar(25);not null;default:''" json:"province"`
	City     string `gorm:"type:varchar(25);not null;default:''" json:"city"`
	Street   string `gorm:"type:varchar(25);not null;default:''" json:"street"`
}
