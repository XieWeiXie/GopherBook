package model

import "github.com/jinzhu/gorm"

type ReceiveAddress struct {
	gorm.Model
	Name     string
	Phone    string
	RegionID uint
}
