package models

var MedalString = "https://www.fina-gwangju2019.com/contentsView.do?pageId=eng107"

const (
	GOLD = iota
	SILVER
	BRONZE
)

var MedalClass = map[int]string{}

func init() {
	MedalClass = make(map[int]string)
	MedalClass[GOLD] = "GOLD"
	MedalClass[SILVER] = "SILVER"
	MedalClass[BRONZE] = "BRONZE"

}

type Medal struct {
	Base    `xorm:"extends"`
	Country string `xorm:"'country'" json:"country"`
	Gold    int    `xorm:"'gold'" json:"gold"`
	Silver  int    `xorm:"'silver''" json:"silver"`
	Bronze  int    `xorm:"'bronze'" json:"bronze"`
}

func (M Medal) TableName() string { return "medal" }

type SportMedal struct {
	Base       `xorm:"extends"`
	SportClass int    `xorm:"'sport_class'" json:"sport_class"`
	EventClass int    `xorm:"'event_class'" json:"event_class"`
	Event      string `xorm:"'event'" json:"event"`
	Gold       string `json:"gold"`
	Silver     string `json:"silver"`
	Bronze     string `json:"bronze"`
}

func (S SportMedal) TableName() string { return "sport_medal" }
