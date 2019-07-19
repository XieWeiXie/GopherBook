package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type VoteOriginal struct {
	Title       string
	AdminId     uint
	Description string
	Choice      []Choice
	DeadLine    time.Time
	IsAnonymous bool
	IsSingle    bool
}

type Vote struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(32)"`
	AdminId     uint   `json:"admin_id"`
	Description string `json:"description" gorm:"type:varchar(64)"`
	Choice      []Choice
	DeadLine    time.Time
	IsAnonymous bool
	IsSingle    bool
}

type Choice struct {
	gorm.Model
	VoteId uint
	Title  string `gorm:"type:varchar(32)" json:"title"`
	Number int    `gorm:"type:integer" json:"number"`
}
