package model

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Avatar        string `gorm:"type:varchar(255)" json:"avatar"`
	NickName      string `gorm:"type:varchar(10)" json:"nick_name"`
	AccountString string `gorm:"type:varchar(15)" json:"account_string"`
	AccountQR     string `gorm:"type:varchar(255)" json:"account_qr"`
	Gender        int    `gorm:"type:integer" json:"gender"`
	Location      string `gorm:"type:varchar(255)" json:"location"`
	Signal        string `gorm:"type:varchar(64)" json:"signal"`
	Addresses     []Address
	Receipts      []Receipt
}

func (Person) TableName() string {
	return "wechat_persons"
}

type Address struct {
	gorm.Model
	PersonID     uint
	Name         string `gorm:"type:varchar(10)" json:"name"`
	Phone        string `gorm:"type:varchar(11)" json:"phone"`
	LocationInfo string `gorm:"type:varchar(64);not null" json:"location_info"`
	Detail       string `gorm:"type:varchar(128);not null" json:"detail"`
	Code         string `gorm:"type:varchar(6);not null" json:"code"`
}

func (Address) TableName() string {
	return "wechat_address"
}

const (
	TYPEPERSONALE = iota
	TYPECOMPANY
)

type Receipt struct {
	gorm.Model
	PersonID       uint
	Type           int    `gorm:"type:integer" json:"type"`
	Name           string `gorm:"type:varchar(32)" json:"name"`
	TaxNumber      string `gorm:"type:varchar(32)"`
	CompanyAddress string `gorm:"type:varchar(64)" json:"company_address"`
	Phone          string `gorm:"type:varchar(11)" json:"phone"`
	Bank           string `gorm:"type:varchar(32)" json:"bank"`
	BankCount      string `gorm:"type:varchar(18)" json:"bank_count"`
}

func (Receipt) TableName() string {
	return "wechat_receipt"
}
