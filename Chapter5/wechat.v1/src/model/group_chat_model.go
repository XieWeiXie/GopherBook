package model

import "github.com/jinzhu/gorm"

type GroupChat struct {
	gorm.Model
	Contents []Content
	Persons  []Person
}
