package model

import "github.com/jinzhu/gorm"

const (
	PERSONALE = iota
	COMPANY
)

type ReceiptDetail struct {
	gorm.Model
	Type       int
	Name       string
	TaxNumber  string
	LocationID uint
	Phone      string
	BankID     uint
}

type Location struct {
	gorm.Model
	RegionID uint
	Detail   string
}
