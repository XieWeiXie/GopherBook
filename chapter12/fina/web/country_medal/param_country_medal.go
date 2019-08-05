package country_medal

import "gopkg.in/go-playground/validator.v9"

type GetCountryMedalParam struct {
	Name string `json:"name" validator:"name"`
	Year int    `json:"year" validator:"min=1975"`
}

func (G GetCountryMedalParam) Valid() error {
	return validator.New().Struct(G)
}

type RankCountryMedalParam struct {
	Year   int    `json:"year" validator:"min=1973"`
	SortBy string `json:"sort_by" validator:"eq=gold|eq=silver|eq=bronze|eq=total"`
}

func (R RankCountryMedalParam) Valid() error {
	return validator.New().Struct(R)
}
