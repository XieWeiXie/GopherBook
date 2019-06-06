package model

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type Exchange struct {
	gorm.Model
	Name     string    `gorm:"type:varchar" json:"name"`
	ZeroTime time.Time `gorm:"type:timestamp with time zone" json:"zero_time"`
	EndTime  time.Time `gorm:"type:timestamp with time zone" json:"end_time"`
	Price    sql.NullFloat64
}
