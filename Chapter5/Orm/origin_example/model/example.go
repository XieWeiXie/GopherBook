package model

import "time"

type Base struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
type Person struct {
	Base
	Avatar        string
	NickName      string
	AccountString string
	AccountQR     string
	Gender        int
	Location      string
	Signal        string
	Addresses     []Address
	Receipts      []Receipt
}

func (Person) TableName() string {
	return "wechat_persons"
}

type PersonSerializer struct {
	ID            uint      `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Avatar        string    `json:"avatar"`
	NickName      string    `json:"nick_name"`
	AccountString string    `json:"account_string"`
	Gender        string    `json:"gender"`
	Location      string    `json:"location"`
	Signal        string    `json:"signal"`
}

func (p Person) JSONSerializer() PersonSerializer {
	genderString := func(gender int) string {
		if gender == 0 {
			return "男"
		}
		return "女"
	}
	return PersonSerializer{
		ID:            p.ID,
		CreatedAt:     p.CreatedAt.Truncate(time.Hour),
		UpdatedAt:     p.UpdatedAt.Truncate(time.Second),
		Avatar:        p.Avatar,
		NickName:      p.NickName,
		AccountString: p.AccountString,
		Gender:        genderString(p.Gender),
		Location:      p.Location,
		Signal:        p.Signal,
	}
}

type Address struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	PersonID     uint
	Name         string
	Phone        string
	LocationInfo string
	Detail       string
	Code         string
}

func (Address) TableName() string {
	return "wechat_address"
}

const (
	TYPEPERSONALE = iota
	TYPECOMPANY
)

type Receipt struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	PersonID       uint
	Type           int
	Name           string
	TaxNumber      string
	CompanyAddress string
	Phone          string
	Bank           string
	BankCount      string
}
