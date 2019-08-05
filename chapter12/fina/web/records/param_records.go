package records

import "gopkg.in/go-playground/validator.v9"

type GetRecordParam struct {
	Name             string `json:"name" validator:"required"`
	All              bool   `json:"all"`
	SportClass       int    `json:"sport_class" validator:"min=0,max=5"`
	CompetitionClass int    `json:"competition_class" validator:"min=0,max=2"`
}

func (G GetRecordParam) Valid() error {
	return validator.New().Struct(G)
}
