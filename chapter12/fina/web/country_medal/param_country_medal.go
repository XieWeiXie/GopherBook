package country_medal

import "gopkg.in/go-playground/validator.v9"

type GetCountryMedalParam struct {
	Name string `json:"name" validator:"name"`
	Year int    `json:"year" validator:"min=1975"`
}

func (G GetCountryMedalParam) Valid() error {
	return validator.New().Struct(G)
}
