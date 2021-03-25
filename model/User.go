package model

import "gorm.io/gorm"

//用户模型，保存用户基本信息
//Email字段创建索引
type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(25)" json:"username"`
	Gender   int    `gorm:"type:int" json:"gender"`
	Email    string `gorm:"type:varchar(30)" json:"email" "index"`
	Avatar   string `gorm:"type:varchar(100)" json:"avatar"`
}
