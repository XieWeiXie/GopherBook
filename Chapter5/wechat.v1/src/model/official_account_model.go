package model

import "github.com/jinzhu/gorm"

type OfficialAccount struct {
	gorm.Model
	WeChatAccount string
	Company       string
	PersonID      uint
}
