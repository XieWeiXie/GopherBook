package model

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"type:varchar" json:"name"`
}
