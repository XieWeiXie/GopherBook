package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Activity struct {
	gorm.Model
	Title    string    `gorm:"type:varchar" json:"title"`
	FromDate time.Time `gorm:"type:timestamp with time zone" json:"from_date"`
	ToDate   time.Time `gorm:"type:timestamp with time zone" json:"to_date"`
	Products []Product `gorm:"type:many2many: activity2products" json:"products"`
}
