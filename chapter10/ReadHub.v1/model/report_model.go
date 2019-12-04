package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Report struct {
	gorm.Model
	Title          string `gorm:"type:varchar" json:"title"`
	Reporter       string `gorm:"type:varchar" json:"reporter"`
	Content        string `gorm:"type:text" json:"content"`
	SupportID      uint
	Support        Support
	RelatedReports []RelatedReport `gorm:"many2many:report2related_report" json:"related_reports"`
	Events         []Event         `gorm:"many2many:report2event" json:"events"`
}

type Support struct {
	gorm.Model
	AffirmativeNumber int    `gorm:"type:integer" json:"affirmative_number"`
	NegativeNumber    int    `gorm:"type:integer" json:"negative_number"`
	Type              string `gorm:"type:varchar" json:"type"`
}

type RelatedReport struct {
	gorm.Model
	Title    string   `gorm:"type:varchar" json:"title"`
	Reporter string   `gorm:"type:varchar" json:"reporter"`
	Url      string   `gorm:"type:varchar" json:"url"`
	Reports  []Report `gorm:"many2many:report2related_report" json:"reports"`
}

type Event struct {
	gorm.Model
	ReportTime time.Time `gorm:"type:timestamp with time zone" json:"report_time"`
	Title      string    `gorm:"type:varchar" json:"title"`
	Reports    []Report  `gorm:"many2many:report2event" json:"reports"`
}
