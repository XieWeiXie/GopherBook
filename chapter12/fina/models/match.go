package models

import "time"

/*
The 18th FINA World Championships Gwangju
- about: 介绍
- symbols: 象征
*/

type FiFaChampionships struct {
	Base           `xorm:"extends"`
	ShortSlogan    string    `xorm:"text 'short_slogan'"`
	StartDate      time.Time `xorm:"datetime notnull 'start_date'" json:"start_date"`
	EndDate        time.Time `xorm:"datetime notnull 'end_date'" json:"end_date"`
	DisciplinesIds []int64   `xorm:"blob 'disciplines_ids'" json:"disciplines_ids"`
	VenuesIds      []int64   `xorm:"blob 'venus_ids'" json:"venus_ids"`
}

func (F FiFaChampionships) TableName() string {
	return "fifa_championships"
}

type Disciplines struct {
	Base `xorm:"extends"`
	Name string `xorm:"varchar(32) 'name'" json:"name"`
}

func (D Disciplines) TableName() string {
	return "disciplines"
}

type Venues struct {
	Base `xorm:"extends"`
	Name string `xorm:"varchar(32) 'name'" json:"name"`
}

func (V Venues) TableName() string {
	return "venues"
}
