package model

import "time"

type Base struct {
	ID        uint       `xorm:"pk 'id'"`
	CreatedAt time.Time  `xorm:"created"`
	UpdatedAt time.Time  `xorm:"updated"`
	DeletedAt *time.Time `xorm:"deleted index"`
}
type Person struct {
	Base          `xorm:"extends"`
	Avatar        string `xorm:"varchar(255)" json:"avatar"`
	NickName      string `xorm:"varchar(10)" json:"nick_name"`
	AccountString string `xorm:" varchar(15)" json:"account_string"`
	AccountQR     string `xorm:" varchar(255)" json:"account_qr"`
	Gender        int    `xorm:" integer" json:"gender"`
	Location      string `xorm:" varchar(255)" json:"location"`
	Signal        string `xorm:" varchar(64)" json:"signal"`
	Addresses     []Address
	Receipts      []Receipt
}

func (Person) TableName() string {
	return "wechat_persons"
}

type Address struct {
	ID           uint       `xorm:"pk 'id'"`
	CreatedAt    time.Time  `xorm:"created"`
	UpdatedAt    time.Time  `xorm:"updated"`
	DeletedAt    *time.Time `xorm:"deleted index"`
	PersonID     uint
	Name         string `xorm:" varchar(10)" json:"name"`
	Phone        string `xorm:" varchar(11)" json:"phone"`
	LocationInfo string `xorm:" varchar(64)" json:"location_info"`
	Detail       string `xorm:" varchar(128)" json:"detail"`
	Code         string `xorm:" varchar(6)  notnull" json:"code"`
}

func (Address) TableName() string {
	return "wechat_address"
}

const (
	TYPEPERSONALE = iota
	TYPECOMPANY
)

type Receipt struct {
	ID             uint       `xorm:"pk 'id'"`
	CreatedAt      time.Time  `xorm:"created"`
	UpdatedAt      time.Time  `xorm:"updated"`
	DeletedAt      *time.Time `xorm:"deleted index"`
	PersonID       uint
	Type           int    `xorm:" integer" json:"type"`
	Name           string `xorm:" varchar(32)" json:"name"`
	TaxNumber      string `xorm:" varchar(32)" json:"tax_number"`
	CompanyAddress string `xorm:" varchar(64)" json:"company_address"`
	Phone          string `xorm:" varchar(11)" json:"phone"`
	Bank           string `xorm:" varchar(32)" json:"bank"`
	BankCount      string `xorm:" varchar(18)" json:"bank_count"`
}
