package v2

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	AccountID       uint
	AccountBalance  sql.NullFloat64
	ExchangesNumber int `gorm:"type:integer" json:"exchanges_number"`
	CouponsNumber   int `gorm:"type:integer" json:"coupons_number"`
	Exchanges       []Exchange
	Coupons         []Coupon
}
