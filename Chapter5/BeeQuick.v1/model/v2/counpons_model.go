package v2

import (
	"github.com/jinzhu/gorm"
)

type Coupon struct {
	gorm.Model
	Exchange
	Token string `gorm:"type:varchar" json:"token"`
}
