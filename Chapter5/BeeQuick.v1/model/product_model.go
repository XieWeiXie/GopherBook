package model

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name          string          `gorm:"type:varchar" json:"name"`
	Avatar        string          `gorm:"type:varchar" json:"avatar"`
	Price         sql.NullFloat64 `json:"price"`
	Amount        int             `gorm:"type:integer" json:"amount"`
	Specification string          `gorm:"type:varchar" json:"specification"`
	Period        int             `gorm:"type:integer" json:"period"`
	BrandID       uint
	UintID        uint
	TagID         uint
}
