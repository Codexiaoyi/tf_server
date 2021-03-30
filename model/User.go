package model

import "gorm.io/gorm"

//用户模型，保存用户基本信息
//Email字段创建索引
type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(25)" json:"username"`
	Gender   int    `gorm:"type:int" json:"gender"`
	Email    string `gorm:"type:varchar(30);index" json:"email"`
	Avatar   string `gorm:"type:varchar(100)" json:"avatar"`
	Birthday
	Address
}

//生日
type Birthday struct {
	Year  int `gorm:"type:int" json:"year"`
	Month int `gorm:"type:int" json:"month"`
	Day   int `gorm:"type:int" json:"day"`
}

//住址
type Address struct {
	Province string `gorm:"type:varchar(25)" json:"province"`
	City     string `gorm:"type:varchar(25)" json:"city"`
	Street   string `gorm:"type:varchar(25)" json:"street"`
}
