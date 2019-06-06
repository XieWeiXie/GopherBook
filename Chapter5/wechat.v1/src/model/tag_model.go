package model

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name    string `gorm:"type:varchar"json:"name"`
	Count   int    `gorm:"type:integer" json:"count"`
	Persons []Person
}
