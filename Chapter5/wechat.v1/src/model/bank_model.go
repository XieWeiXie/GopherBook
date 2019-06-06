package model

import "github.com/jinzhu/gorm"

type Bank struct {
	gorm.Model
	Name    string
	Account string
}
