package competition

import "gopkg.in/go-playground/validator.v9"

type GetCompetitionParam struct {
	Class int `json:"class" validator:"eq=0|eq=1|eq=2"`
}

func (G GetCompetitionParam) Valid() error {
	return validator.New().Struct(G)

}
