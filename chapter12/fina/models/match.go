package models

import (
	"GopherBook/chapter12/fina/pkg/database"
	"time"
)

type FiFaChampionships struct {
	Base           `xorm:"extends"`
	NumberOlympic  int       `xorm:"integer(4) 'number_olympic'" json:"number_olympic"`
	ShortSlogan    string    `xorm:"text 'short_slogan'"`
	StartDate      time.Time `xorm:"datetime notnull 'start_date'" json:"start_date"`
	EndDate        time.Time `xorm:"datetime notnull 'end_date'" json:"end_date"`
	DisciplinesIds []int64   `xorm:"blob 'disciplines_ids'" json:"disciplines_ids"`
	VenuesIds      []int64   `xorm:"blob 'venus_ids'" json:"venus_ids"`
}

func (F FiFaChampionships) TableName() string {
	return "championships"
}

type FiFaChampionshipSerializer struct {
	Id            int64            `json:"id"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
	NumberOlympic int              `json:"number_olympic"`
	ShortSlogan   string           `json:"short_slogan"`
	Date          string           `json:"date"`
	Disciplines   []KindSerializer `json:"disciplines"`
	Venues        []KindSerializer `json:"venues"`
}

const (
	DISCIPLINE = iota
	VENUES
)

var KindClass map[int]string

func init() {
	KindClass = make(map[int]string)
	KindClass[DISCIPLINE] = "discipline"
	KindClass[VENUES] = "venues"
}

type Kinds struct {
	Base  `xorm:"extends"`
	Name  string `xorm:"varchar(32) 'name'" json:"name"`
	Class int    `xorm:"integer(1)" json:"class"`
}

func (K Kinds) TableName() string {
	return "kinds"
}

type KindSerializer struct {
	Id          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Class       int       `json:"class"`
	ClassString string    `json:"class_string"`
}

func (K Kinds) Serializer() KindSerializer {
	return KindSerializer{
		Id:          K.Id,
		CreatedAt:   K.CreatedAt.Truncate(time.Second),
		UpdatedAt:   K.UpdatedAt.Truncate(time.Second),
		Name:        K.Name,
		Class:       K.Class,
		ClassString: KindClass[K.Class],
	}
}

type Symbol struct {
	Base                    `xorm:"extends"`
	SymbolText              string  `xorm:"varchar(64) 'symbol_text'" json:"symbol_text"` // 标志
	SymbolTextImage         string  `xorm:"varchar(128) 'symbol_text_image'" json:"symbol_text_image"`
	SymbolTextShort         string  `xorm:"varchar(12) 'symbol_text_short'" json:"symbol_text_short"`
	SymbolDescription       string  `xorm:"varchar(64) 'symbol_description'" json:"symbol_description"` // 标语
	SymbolDescriptionImage  string  `xorm:"varchar(128) 'symbol_description_image'" json:"symbol_description_image"`
	SymbolDescriptionShort  string  `xorm:"varchar(12) 'symbol_description_short'" json:"symbol_description_short"`
	SymbolAnimalImage       string  `xorm:"varchar(128) 'symbol_animal_image'" json:"symbol_animal_image"` // 吉祥物
	SymbolAnimalDescription string  `xorm:"varchar(64) 'symbol_animal_description'" json:"symbol_animal_description"`
	SymbolAnimalShort       string  `xorm:"varchar(12) 'symbol_animal_short'" json:"symbol_animal_short"`
	BlueVersions            []int64 `xorm:"'blue_versions'" json:"blue_versions"`
}

func (S Symbol) TableName() string { return "symbol" }

type SymbolSerializer struct {
	Id                      int64     `json:"id"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	SymbolText              string    `json:"symbol_text"`
	SymbolTextImage         string    `json:"symbol_text_image"`
	SymbolTextShort         string    `json:"symbol_text_short"`
	SymbolDescription       string    `json:"symbol_description"`
	SymbolDescriptionImage  string    `json:"symbol_description_image"`
	SymbolDescriptionShort  string    `json:"symbol_description_short"`
	SymbolAnimalImage       string    `json:"symbol_animal_image"`
	SymbolAnimalDescription string    `json:"symbol_animal_description"`
	SymbolAnimalShort       string    `json:"symbol_animal_short"`
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
	Short       string `xorm:"varchar(1)"json:"short"`
	EnName      string `xorm:"varchar(12) 'en_name'"json:"en_name"`
	ChName      string `xorm:"varchar(32) 'ch_name'"json:"ch_name"`
	Description string `xorm:"varchar(64) 'description'" json:"description"`
	Image       string `xorm:"varchar(128) 'image'" json:"image"`
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
	Image       string    `json:"image"`
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
		Image:       B.Image,
	}
}
