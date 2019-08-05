package kind

import "gopkg.in/go-playground/validator.v9"

type GetKindParam struct {
	Class int `json:"class" validator:"eq=0|eq=1"`
}

func (G GetKindParam) Valid() error {
	return validator.New().Struct(G)
}
