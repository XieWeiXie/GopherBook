package model

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	LevelID  uint
	Phone    string    `gorm:"type:varchar" json:"phone"`
	Avatar   string    `gorm:"type:varchar" json:"avatar"`
	Name     string    `gorm:"type:varchar" json:"name"`
	Gender   int       `gorm:"type:integer" json:"gender"` // 0 男 1 女
	Birthday time.Time `gorm:"type:timestamp with time zone" json:"birthday"`
	Points   sql.NullFloat64
}
