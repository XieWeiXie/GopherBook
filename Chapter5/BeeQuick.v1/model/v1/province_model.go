package model_v1

import "time"

type Province struct {
	base     `xorm:"extends"`
	Name     string `xorm:"varchar(10)" json:"name"`
	AdCode   string `xorm:"varchar(10)" json:"ad_code"`
	CityCode string `xorm:"varchar(6)" json:"city_code"`
	Center   string `xorm:"varchar(32)" json:"center"`
	Level    string `xorm:"varchar(10)" json:"level"`
}

func (p Province) TableName() string {
	return "beeQuick_province"
}

type ProvinceSerializer struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	AdCode    string    `json:"ad_code"`
	CityCode  string    `json:"city_code"`
	Center    string    `json:"center"`
	Level     string    `json:"level"`
}

func (p Province) Serializer() ProvinceSerializer {
	return ProvinceSerializer{
		Id:        int(p.ID),
		CreatedAt: p.CreatedAt.Truncate(time.Second),
		UpdatedAt: p.UpdatedAt.Truncate(time.Second),
		Name:      p.Name,
		AdCode:    p.AdCode,
		Center:    p.Center,
		Level:     p.Level,
		CityCode:  p.CityCode,
	}
}
