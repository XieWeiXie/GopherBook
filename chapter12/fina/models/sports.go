package models

const (
	SWIMMING = iota
	DIVING
	HIGHDIVING
	ARTISICSWIMMING
	OPENWATER
	WATERPOLO
)

var SportClass = map[int]string{}

func init() {
	SportClass = make(map[int]string)
	SportClass[SWIMMING] = "游泳"
	SportClass[DIVING] = "跳水"
	SportClass[HIGHDIVING] = "高空跳水"
	SportClass[ARTISICSWIMMING] = "花样游泳"
	SportClass[OPENWATER] = "公开水域游泳"
	SportClass[WATERPOLO] = "水球"

}

type Sports struct {
	Base           `xorm:"extends"`
	SportClass     int     `xorm:"'sport_class'"`
	SportName      string  `xorm:"'sport_name'" json:"sport_name"`
	Description    string  `json:"description"`
	CompetitionIds []int64 `xorm:"'competition_ids'" json:"competition_ids"`
	Rule           string  `xorm:"'rule'" json:"rule"`
	BriefHistory   string  `xorm:"'brief_history'" json:"brief_history"`
	Top3Ids        []int64 `xorm:"'top3_ids'" json:"top3"`
}

func (S Sports) TableName() string { return "sports" }

type Top3 struct {
	Base    `xorm:"extends"`
	Rank    int    `xorm:"integer(11)" json:"rank"`
	Number  int    `xorm:"integer(11)" json:"number"`
	Country string `xorm:"integer(11)" json:"country"`
}

func (T Top3) TableName() string { return "top3" }

const (
	MAN = iota
	WOMAN
	MIX
	TEAM
)

var CompetitionClass = map[int]string{}

func init() {
	CompetitionClass = make(map[int]string)
	CompetitionClass[MAN] = "MAN"
	CompetitionClass[WOMAN] = "WOMAN"
	CompetitionClass[MIX] = "MIX"
	CompetitionClass[TEAM] = "TEAM"

}

type Competitions struct {
	Base             `xorm:"extends"`
	CompetitionClass int    `xorm:"'competition_class'" json:"competition_class"`
	Detail           string `xorm:"'detail'" json:"detail"`
}

func (C Competitions) TableName() string { return "competitions" }
