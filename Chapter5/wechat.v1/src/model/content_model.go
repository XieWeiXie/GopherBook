package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	TYPEFORVIDEOCALL = iota // video call
	TYPEFORVOICECALL
	TYPEFOREMOJI
	TYPEFORTEXT
	TYPEFORVOICE
)

type Content struct {
	gorm.Model
	Date         time.Time `gorm:"type:timestamp with time zone" json:"date"`
	TypeID       uint
	PersonID     uint
	Detail       string `gorm:"type:varchar" json:"detail"`
	Status       Status
	SingleChatID uint
}

type Type struct {
	gorm.Model
	Code int `gorm:"type:integer" json:"code"`
}

type Status struct {
	gorm.Model
	Duration  string `gorm:"type:varchar" json:"duration"`
	ContentID uint
}
