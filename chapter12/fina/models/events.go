package models

import (
	"github.com/wuxiaoxiaoshen/GopherBook/chapter12/fina/pkg/database"
	"time"
)

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

type CountryMedal struct {
	Base      `xorm:"extends"`
	Year      int   `json:"year"`
	CountryId int64 `xorm:"index 'country_id'" json:"country_id"`
	Gold      int   `xorm:"'gold'" json:"gold"`
	Silver    int   `xorm:"'silver'" json:"silver"`
	Bronze    int   `xorm:"'bronze'" json:"bronze"`
}

func (CC CountryMedal) TableName() string { return "medal" }

type CountryMedalSerializer struct {
	Id          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Year        int       `json:"year"`
	CountryId   int64     `json:"country_id"`
	CountryName string    `json:"country_name"`
	Gold        int       `json:"gold"`
	Silver      int       `json:"silver"`
	Bronze      int       `json:"bronze"`
}

func (CC CountryMedal) Serializer() CountryMedalSerializer {
	var Country Country
	database.MySQL.ID(CC.CountryId).Get(&Country)
	return CountryMedalSerializer{
		Id:          CC.Id,
		CreatedAt:   CC.CreatedAt.Truncate(time.Second),
		UpdatedAt:   CC.UpdatedAt.Truncate(time.Second),
		Year:        CC.Year,
		CountryId:   CC.CountryId,
		CountryName: Country.Name,
		Gold:        CC.Gold,
		Silver:      CC.Silver,
		Bronze:      CC.Bronze,
	}
}

type Country struct {
	Base  `xorm:"extends"`
	Name  string `xorm:"unique" json:"name"`
	Short string `xorm:"unique" json:"short"`
}

func (Cry Country) TableName() string {
	return "country"
}

type CountrySerializer struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Short     string    `json:"short"`
}

func (Cry Country) Serializer() CountrySerializer {
	return CountrySerializer{
		Id:        Cry.Id,
		CreatedAt: Cry.CreatedAt.Truncate(time.Second),
		UpdatedAt: Cry.UpdatedAt.Truncate(time.Second),
		Name:      Cry.Name,
		Short:     Cry.Short,
	}
}
