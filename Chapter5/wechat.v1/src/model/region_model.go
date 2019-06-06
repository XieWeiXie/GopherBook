package model

import "github.com/jinzhu/gorm"

type Region struct {
	gorm.Model
	Code   string `gorm:"type:varchar" json:"code"`
	CityID uint
}

type Province struct {
	gorm.Model
	Code   string `gorm:"type:varchar" json:"code"`
	Cities []City
}

type City struct {
	gorm.Model
	Name       string `json:"name"`
	ProvinceID uint
}

type From struct {
	gorm.Model
	Location string `gorm:"type:varchar" json:"location"`
}
