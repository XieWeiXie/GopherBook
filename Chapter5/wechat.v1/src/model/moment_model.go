package model

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
)

const (
	MOMENTTYPETEXT = iota
	MOMENTTYPEVIDEO
	MOMENTPICTURE
	MOMENTSONG
	MOMENTPASSAGE
	MOMENTOTHER
)

type Moment struct {
	gorm.Model
	PersonID   uint
	SettingID  uint
	Time       time.Time `gorm:"type:timestamp with time zone"`
	MomentType int       `gorm:"type:integer" json:"moment_type"`
	Like       []Like
}

const (
	THREEDAYMOMENT = iota
	ONEMONTHMOMENT
	HALFYEARMOMENT
	ALLMOMENT
)

type Setting struct {
	gorm.Model
	Code int `gorm:"type:integer" json:"code"`
}

type Like struct {
	gorm.Model
	MomentID  uint
	PersonIDs postgres.Jsonb
}

type Comment struct {
	gorm.Model
	Detail   string
	PersonID uint
}

type CommentAll struct {
	gorm.Model
	MomentID uint
	Comments []Comment
}
