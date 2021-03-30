package model

import "gorm.io/gorm"

//团队模型
type Team struct {
	gorm.Model
	TeamName     string `gorm:"type:varchar(20)" json:"teamname"`
	Avatar       string `gorm:"type:varchar(100)" json:"avatar"`
	Create       string `gorm:"type:varchar(30)" json:"create"`
	Introduction string `gorm:"type:varchar(200)" json:"introduction"`
}
