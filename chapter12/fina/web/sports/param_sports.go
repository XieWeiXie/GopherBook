package sports

import "gopkg.in/go-playground/validator.v9"

type GetSportParam struct {
	Class int `json:"class" validator:"eq=0|eq=1|eq=2|eq=3|eq=4|eq=5"`
}

func (G GetSportParam) Valid() error {
	return validator.New().Struct(G)
}
