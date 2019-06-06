package model

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	Name         string `gorm:"type:varchar" json:"name"`
	Count        int    `gorm:"type:integer" json:"count"`
	Announcement string `gorm:"type:varchar(100)" json:"announcement"`
	GroupChatID  uint
}
