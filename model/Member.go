package model

import "gorm.io/gorm"

//团队成员模型，描述团队与成员关系
//添加两个索引，方便查询
type Member struct {
	gorm.Model
	TeamId   int    `gorm:"type:int;index:idx_teamid" json:"teamid"`
	Email    string `gorm:"type:varchar(30);index:idx_email" json:"email"`
	IsLeader bool   `gorm:"type:int" json:"isleader"`
}
