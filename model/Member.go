package model

import "gorm.io/gorm"

//团队成员模型，描述团队与成员关系
type Member struct {
	gorm.Model
	TeamId   int    `gorm:"type:int" json:"teamid"`
	Email    string `gorm:"type:varchar(30)" json:"email"`
	IsLeader bool   `gorm:"type:int" json:"isleader"`
}
