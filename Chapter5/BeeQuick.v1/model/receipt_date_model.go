package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ReceiptDate struct {
	gorm.Model
	ReceiveDateID uint
	ReceiveDate   ReceiveDate
	FormTime      time.Time `gorm:"type:timestamp with time zone" json:"form_time"`
	ToTime        time.Time `gorm:"type:timestamp with time zone" json:"to_time"`
}

type ReceiveDate struct {
	gorm.Model
	Item string `gorm:"type:varchar" json:"item"`
}
