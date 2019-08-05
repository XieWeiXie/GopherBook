package models

import "time"

type FiFa struct {
	Base                `xorm:"extends"`
	Description         string `json:"description"`
	Established         string `xorm:"varchar(32) notnull 'established'" json:"established"`
	Headquarters        string `xorm:"varchar(32) notnull 'headquarters'" json:"headquarters"`
	NationalMember      string `xorm:"varchar(24) notnull 'national_member'" json:"national_member"`
	NumberOfDisciplines string `xorm:"varchar(24) notnull 'number_of_disciplines'" json:"number_of_disciplines"`
}

func (F FiFa) TableName() string {
	return "fifa"
}

type FiFaSerializer struct {
	Id                  int64     `json:"id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	Description         string    `json:"description"`
	Established         string    `json:"established"`
	Headquarters        string    `json:"headquarters"`
	NationalMember      string    `json:"national_member"`
	NumberOfDisciplines string    `json:"number_of_disciplines"`
}

func (F FiFa) Serializer() FiFaSerializer {
	return FiFaSerializer{
		Id:                  F.Id,
		CreatedAt:           F.CreatedAt.Truncate(time.Second),
		UpdatedAt:           F.UpdatedAt.Truncate(time.Second),
		Description:         F.Description,
		Established:         F.Established,
		Headquarters:        F.Headquarters,
		NumberOfDisciplines: F.NumberOfDisciplines,
		NationalMember:      F.NationalMember,
	}
}

type FiFaHistory struct {
	Base   `xorm:"extends"`
	Year   int    `json:"year"`
	Detail string `json:"detail"`
}

func (F FiFaHistory) TableName() string {
	return "history"
}

type FiFaHistorySerializer struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Year      int       `json:"year"`
	Detail    string    `json:"detail"`
}

func (F FiFaHistory) Serializer() FiFaHistorySerializer {
	return FiFaHistorySerializer{
		Id:        F.Id,
		CreatedAt: F.CreatedAt.Truncate(time.Second),
		UpdatedAt: F.UpdatedAt.Truncate(time.Second),
		Year:      F.Year,
		Detail:    F.Detail,
	}
}
