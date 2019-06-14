package model_v1

import "time"

type Shop struct {
	base       `xorm:"extends"`
	Location   string   `xorm:"varchar(255)" json:"location"`
	ProvinceId int64    `xorm:"index"`
	Province   Province `xorm:"-"`
	Name       string   `xorm:"varchar(64)"`
}

func (c Shop) TableName() string {
	return "beeQuick_shop"
}

type ShopSerializer struct {
	Id         int64              `json:"id"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	ProvinceId int64              `json:"province_id"`
	Province   ProvinceSerializer `json:"province"`
	Name       string             `json:"name"`
	Location   string             `json:"location"`
}

func (c Shop) Serializer() ShopSerializer {
	return ShopSerializer{
		Id:         int64(c.ID),
		CreatedAt:  c.CreatedAt.Truncate(time.Second),
		UpdatedAt:  c.UpdatedAt.Truncate(time.Second),
		Province:   c.Province.Serializer(),
		ProvinceId: c.ProvinceId,
		Name:       c.Name,
		Location:   c.Location,
	}
}
