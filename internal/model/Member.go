package model

import "gorm.io/gorm"

//团队成员模型，描述团队与成员关系
//添加两个索引，方便查询
type Member struct {
	gorm.Model
	TeamId   int    `gorm:"type:int;not null;index:idx_teamid;index:idx_teamid_email,priority:1" json:"teamid"`
	Email    string `gorm:"type:varchar(30);not null;default:'';index:idx_email;index:idx_teamid_email,priority:2" json:"email"`
	IsLeader bool   `gorm:"type:int;not null;default:false" json:"isleader"`
}
