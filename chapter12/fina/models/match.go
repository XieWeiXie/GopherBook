package models

import (
	"GopherBook/chapter12/fina/pkg/database"
	"time"
)

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

type Symbol struct {
	Base                    `xorm:"extends"`
	SymbolText              string  `xorm:"text 'symbol_text'" json:"symbol_text"`
	SymbolDescription       string  `xorm:"text 'symbol_description'" json:"symbol_description"`
	SymbolAnimalImage       string  `xorm:"text 'symbol_animal_image'" json:"symbol_animal_image"`
	SymbolAnimalDescription string  `xorm:"text 'symbol_animal_description'" json:"symbol_animal_description"`
	BlueVersions            []int64 `xorm:"'blue_versions'" json:"blue_versions"`
}

func (S Symbol) TableName() string { return "symbol" }

type SymbolSerializer struct {
	Id                      int64     `json:"id"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	SymbolText              string    `json:"symbol_text"`
	SymbolDescription       string    `json:"symbol_description"`
	SymbolAnimalImage       string    `json:"symbol_animal_image"`
	SymbolAnimalDescription string    `json:"symbol_animal_description"`
	BlueVersions            []BlueSerializer
}

func (S Symbol) Serializer() SymbolSerializer {

	blueVersions := func(ids []int64) []BlueSerializer {
		var blues []Blue
		database.MySQL.In("id", S.BlueVersions).Find(&blues)
		var results []BlueSerializer
		for _, i := range blues {
			results = append(results, i.Serializer())
		}
		return results
	}

	return SymbolSerializer{
		Id:                      S.Id,
		CreatedAt:               S.CreatedAt,
		UpdatedAt:               S.UpdatedAt,
		SymbolText:              S.SymbolText,
		SymbolAnimalImage:       S.SymbolAnimalImage,
		SymbolAnimalDescription: S.SymbolAnimalDescription,
		BlueVersions:            blueVersions(S.BlueVersions),
	}
}

/**/

type Blue struct {
	Base        `xorm:"extends"`
	Short       string `json:"short"`
	EnName      string `xorm:"'en_name'"json:"en_name"`
	ChName      string `xorm:"'ch_name'"json:"ch_name"`
	Description string `xorm:"'description'" json:"description"`
}

func (B Blue) TableName() string { return "blue" }

type BlueSerializer struct {
	Id          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Short       string    `json:"short"`
	EnName      string    `json:"en_name"`
	ChName      string    `json:"ch_name"`
	Description string    `json:"description"`
}

func (B Blue) Serializer() BlueSerializer {
	return BlueSerializer{
		Id:          B.Id,
		CreatedAt:   B.CreatedAt.Truncate(time.Second),
		UpdatedAt:   B.UpdatedAt.Truncate(time.Second),
		Short:       B.Short,
		EnName:      B.EnName,
		ChName:      B.ChName,
		Description: B.Description,
	}
}
