package model_v1

import "time"

type Product struct {
	base          `xorm:"extends"`
	ShopId        int64   `xorm:"index"`
	Name          string  `xorm:"varchar(128) 'name'" json:"name"`
	Avatar        string  `xorm:"varchar(255) 'avatar'" json:"avatar"`
	Price         float64 `xorm:"double 'price'" json:"price"`
	Discount      float64 `xorm:"double default(1) 'discount'" json:"discount"` // 默认为 1
	Specification string  `xorm:"varchar(128) 'specification'" json:"specification"`
	BrandId       int64   `xorm:"index"`
	TagsId        int64   `xorm:"index"`
	Period        string  `xorm:"varchar(64)" json:"period"`
	UintsId       int64   `xorm:"index"`
}

func (p Product) TableName() string {
	return "beeQuick_products"
}

type Uints struct {
	base      `xorm:"extends"`
	Name      string `json:"name"`
	EnName    string `json:"en_name"`
	ShortCode string `json:"short_code"`
}

func (u Uints) TableName() string {
	return "beeQuick_uints"
}

type UintsSerializer struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	EnName    string    `json:"en_name"`
	ShortCode string    `json:"short_code"`
}

func (u Uints) Serializer() UintsSerializer {
	return UintsSerializer{
		Id:        int64(u.ID),
		CreatedAt: u.CreatedAt.Truncate(time.Second),
		UpdatedAt: u.UpdatedAt.Truncate(time.Second),
		Name:      u.Name,
		EnName:    u.Name,
		ShortCode: u.ShortCode,
	}
}

type Brands struct {
	base   `xorm:"extends"`
	ChName string `json:"ch_name"`
	EnName string `json:"en_name"`
}

func (b Brands) TableName() string {
	return "beeQuick_brands"
}

type Tags struct {
	base `xorm:"extends"`
	Name string `json:"name"`
}

func (t Tags) TableName() string {
	return "beeQuick_tags"
}

type Shop2Tags struct {
	TagsId int64 `xorm:"index"`
	ShopId int64 `xorm:"index"`
}

func (s2t Shop2Tags) TableName() string {
	return "beeQuick_shop2Tags"
}

type Product2Tags struct {
	TagsId    int64 `xorm:"index"`
	ProductId int64 `xorm:"index"`
}

func (p2t Product2Tags) TableName() string {
	return "beeQuick_product2Tags"
}