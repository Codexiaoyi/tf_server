package model

import (
	"gorm.io/gorm"
)

//用户相册
type UserAlbum struct {
	Email string `gorm:"type:varchar(30);index;not null" json:"email"`
	Album
}

//团队相册
type TeamAlbum struct {
	TeamId int `gorm:"type:int;index;not null" json:"teamid"`
	Album
}

//相册基本属性
type Album struct {
	gorm.Model
	Public       bool   `gorm:"type:int;not null" json:"public"` //是否公开
	Name         string `gorm:"type:type:varchar(20);not null;default:''" json:"name"`
	Introduction string `gorm:"type:type:varchar(200);not null;default:''" json:"introduction"`
	Cover        string `gorm:"type:type:varchar(100);not null;default:''" json:"cover"` //封面地址
	Like         int    `gorm:"type:int;not null" json:"like"`                           //点赞数量
	Collect      int    `gorm:"type:int;not null" json:"collect"`                        //收藏数量
}