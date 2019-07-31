package models

import "time"

type FiFa struct {
	Base                `xorm:"extends"`
	Description         string `json:"description"`
	Established         string `xorm:"varchar(32) notnull 'established'" json:"established"`
	Headquarters        string `xorm:"varchar(32) notnull 'headquarters'" json:"headquarters"`
	NationalMember      int    `xorm:"integer(11) notnull 'national_member'" json:"national_member"`
	NumberOfDisciplines int    `xorm:"integer(11) notnull 'number_of_disciplines'" json:"number_of_disciplines"`
	Detail              string `json:"detail"`
}

func (F FiFa) TableName() string {
	return "fifa"
}

type FiFaHistory struct {
	Base   `xorm:"extends"`
	Year   int    `json:"year"`
	Detail string `json:"detail"`
}

func (F FiFaHistory) TableName() string {
	return "fifa_history"
}

type PastEvents struct {
	Base           `xorm:"extends"`
	Number         int       `json:"number"`
	Year           time.Time `json:"year"`
	HostCountry    string    `xorm:"varchar(32) 'host_country'"json:"host_country"`
	City           string    `xorm:"varchar(12) 'city'" json:"city"`
	AthletesNumber int       `xorm:"integer(11) 'athletes_number'" json:"athletes_number"`
	CountryNumber  int       `xorm:"integer(11) 'country_number'" json:"country_number"`
}

func (P PastEvents) TableName() string {
	return "past_events"
}
