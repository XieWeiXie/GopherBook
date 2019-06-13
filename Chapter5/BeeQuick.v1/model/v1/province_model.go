package model_v1

type Province struct {
	base `xorm:"extends"`
	Name string `xorm:"varchar(10)" json:"name"`
	Code string
}

func (p Province) TableName() string {
	return "beeQuick_province"
}

type District struct {
	base       `xorm:"extends"`
	Name       string `xorm:"varchar(10)" json:"name"`
	ProvinceId int64  `xorm:"index"`
}
