package models

import (
	"GopherBook/chapter12/fina/pkg/database"
	"time"
)

type RecordMax struct {
	Base             `xorm:"extends"`
	EventName        string    `xorm:"varchar(32)" json:"event_name"`
	Record           string    `xorm:"varchar(32)" json:"record"`
	CountryId        int64     `xorm:"integer(3)" json:"country_id"`
	Date             time.Time `json:"date"`
	Location         string    `json:"location"`
	CompetitionClass int       `json:"competition_class"`
	SportClass       int       `json:"sport_class"`
	Name             string    `json:"name"`
}

func (R RecordMax) TableName() string { return "records" }

type RecordsMaxSerializer struct {
	Id                     int64     `json:"id"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
	EventName              string    `json:"event_name"`
	Record                 string    `json:"record"`
	CountryId              int64     `json:"country_id"`
	CountryName            string    `json:"country_name"`
	Date                   time.Time `json:"date"`
	Location               string    `json:"location"`
	CompetitionClass       int       `json:"competition_class"`
	CompetitionClassString string    `json:"competition_class_string"`
	SportClass             int       `json:"sport_class"`
	SportClassString       string    `json:"sport_class_string"`
	Name                   string    `json:"name"`
}

func (R RecordMax) Serializer() RecordsMaxSerializer {
	var country Country
	database.MySQL.ID(R.CountryId).Get(&country)
	return RecordsMaxSerializer{
		Id:                     R.Id,
		CreatedAt:              R.CreatedAt.Truncate(time.Second),
		UpdatedAt:              R.UpdatedAt.Truncate(time.Second),
		EventName:              R.EventName,
		Record:                 R.Record,
		CountryId:              R.CountryId,
		CountryName:            country.Name,
		Date:                   R.Date,
		Location:               R.Location,
		CompetitionClass:       R.CompetitionClass,
		CompetitionClassString: CompetitionClass[R.CompetitionClass],
		SportClass:             R.SportClass,
		SportClassString:       SportClass[R.SportClass],
		Name:                   R.Name,
	}
}
