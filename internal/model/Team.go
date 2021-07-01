package model

import "gorm.io/gorm"

//团队模型
type Team struct {
	gorm.Model
	TeamName     string `gorm:"type:varchar(20);not null;default:''" json:"teamname"`
	Avatar       string `gorm:"type:varchar(100);not null;default:''" json:"avatar"`
	Email        string `gorm:"type:varchar(30);not null;default:''" json:"email"`
	Introduction string `gorm:"type:varchar(200);not null;default:''" json:"introduction"`
}
