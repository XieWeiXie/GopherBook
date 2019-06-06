package model

import "github.com/jinzhu/gorm"

type Level struct {
	gorm.Model
	Name      string `gorm:"type:varchar" json:"name"`
	ZeroValue int    `gorm:"type:integer" json:"zero_value"`
	EndValue  int    `gorm:"type:integer" json:"end_value"`
	Privilege string `gorm:"type:varchar" json:"privilege"`
	Validity  string `gorm:"type:varchar" json:"validity"`
}
