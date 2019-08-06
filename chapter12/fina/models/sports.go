package models

import (
	"GopherBook/chapter12/fina/pkg/database"
	"time"
)

const (
	SWIMMING = iota
	DIVING
	HIGHDIVING
	ARTISTICSWIMMING
	OPENWATER
	WATERPOLO
)

var SportClass = map[int]string{}
var SportClassEn = map[string]int{}

func init() {
	SportClass = make(map[int]string)
	SportClassEn = make(map[string]int)
	SportClass[SWIMMING] = "游泳"
	SportClass[DIVING] = "跳水"
	SportClass[HIGHDIVING] = "高空跳水"
	SportClass[ARTISTICSWIMMING] = "花样游泳"
	SportClass[OPENWATER] = "公开水域游泳"
	SportClass[WATERPOLO] = "水球"

	SportClassEn["SWIMMING"] = SWIMMING
	SportClassEn["DIVING"] = DIVING
	SportClassEn["HIGHDIVING"] = HIGHDIVING
	SportClassEn["ARTISTICSWIMMING"] = ARTISTICSWIMMING
	SportClassEn["WATERPOLO"] = WATERPOLO

}

type Sports struct {
	Base           `xorm:"extends"`
	Total          int     `xorm:"integer(3) 'total'" json:"total"`
	SportClass     int     `xorm:"'sport_class'"`
	SportName      string  `xorm:"'sport_name'" json:"sport_name"`
	Description    string  `json:"description"`
	CompetitionIds []int64 `xorm:"'competition_ids'" json:"competition_ids"`
	Rule           string  `xorm:"varchar(1024) 'rule'" json:"rule"`
}

func (S Sports) TableName() string { return "sports" }

type SportSerializer struct {
	Id               int64                   `json:"id"`
	CreatedAt        time.Time               `json:"created_at"`
	UpdatedAt        time.Time               `json:"updated_at"`
	Total            int                     `json:"total"`
	SportClass       int                     `json:"sport_class"`
	SportClassString string                  `json:"sport_class_string"`
	SportName        string                  `json:"sport_name"`
	Description      string                  `json:"description"`
	Competitions     []CompetitionSerializer `json:"competitions"`
	Rule             string                  `json:"rule"`
}

func (S Sports) Serializer() SportSerializer {

	competitions := func(ids []int64) []CompetitionSerializer {
		var coms []Competitions
		database.MySQL.In("id", ids).Find(&coms)
		var results []CompetitionSerializer
		for _, i := range coms {
			results = append(results, i.Serializer())
		}
		return results
	}
	return SportSerializer{
		Id:               S.Id,
		CreatedAt:        S.CreatedAt.Truncate(time.Second),
		UpdatedAt:        S.UpdatedAt.Truncate(time.Second),
		Total:            S.Total,
		SportClass:       S.SportClass,
		SportClassString: SportClass[S.SportClass],
		SportName:        S.SportName,
		Description:      S.Description,
		Rule:             S.Rule,
		Competitions:     competitions(S.CompetitionIds),
	}
}

const (
	MAN = iota
	WOMAN
	TEAM
)

var CompetitionClass = map[int]string{}
var CompetitionClassEn = map[string]int{}

func init() {
	CompetitionClass = make(map[int]string)
	CompetitionClassEn = make(map[string]int)
	CompetitionClass[MAN] = "MAN"
	CompetitionClass[WOMAN] = "WOMEN"
	CompetitionClass[TEAM] = "TEAM"
	CompetitionClassEn["MAN"] = MAN
	CompetitionClassEn["WOMEN"] = WOMAN
	CompetitionClassEn["TEAM"] = TEAM
	CompetitionClassEn["MIXED"] = TEAM

}

type Competitions struct {
	Base             `xorm:"extends"`
	CompetitionClass int    `xorm:"'competition_class'" json:"competition_class"`
	Detail           string `xorm:"'detail'" json:"detail"`
}

func (C Competitions) TableName() string { return "competitions" }

type CompetitionSerializer struct {
	Id                     int64     `json:"id"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
	CompetitionClass       int       `json:"competition_class"`
	CompetitionClassString string    `json:"competition_class_string"`
	Detail                 string    `json:"detail"`
}

func (C Competitions) Serializer() CompetitionSerializer {
	return CompetitionSerializer{
		Id:                     C.Id,
		CreatedAt:              C.CreatedAt.Truncate(time.Second),
		UpdatedAt:              C.UpdatedAt.Truncate(time.Second),
		CompetitionClass:       C.CompetitionClass,
		CompetitionClassString: CompetitionClass[C.CompetitionClass],
		Detail:                 C.Detail,
	}
}
