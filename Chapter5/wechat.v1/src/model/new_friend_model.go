package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	STATUSREJECT = iota
	STATUSACCEPT
	STATUSEXPIRE
	STATUSWAIT
)

type NewFriend struct {
	gorm.Model
	Time     time.Time `gorm:"type:timestamp with time zone" json:"time"`
	PersonID uint
	Comment  string `gorm:"type:varchar" json:"comment"`
	Status   int    `gorm:"type:integer" json:"status"`
}
