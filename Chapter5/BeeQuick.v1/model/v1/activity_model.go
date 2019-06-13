package model_v1

import "time"

type Activity struct {
	base      `xorm:"extends"`
	Name      string `xorm:"varchar(32)" json:"name"`
	Title     string `xorm:"varchar(32)" json:"title"`
	Start     time.Time
	End       time.Time
	Avatar    string  `xorm:"varchar(255)" json:"avatar"`
	CompanyId int64   `xorm:"index"`
	Company   Company `xorm:"-"`
}
