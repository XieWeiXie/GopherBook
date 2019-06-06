package model

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	OrderStatus    []OrderStatus
	Status         string `gorm:"type:varchar" json:"status"`
	ShoppingCartID uint
}

type OrderStatus struct {
	gorm.Model
	Product Product
	Amount  int `gorm:"type:integer" json:"amount"`
}
