package model

import "github.com/jinzhu/gorm"

type Brand struct {
	gorm.Model
	EnName string `gorm:"type:varchar" json:"en_name"`
	ChName string `gorm:"type:varchar" json:"ch_name"`
}
