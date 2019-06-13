package model_v1

type Company struct {
	base       `xorm:"extends"`
	Location   string `xorm:"varchar(255)" json:"location"`
	ProvinceId int64  `xorm:"index"`
	District   int64  `xorm:"index"`
	Name       string `xorm:"varchar(64)"`
}
