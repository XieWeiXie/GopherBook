package model

import "github.com/jinzhu/gorm"

type SingleChat struct {
	gorm.Model
	PersonID uint
	Contents []Content
}
