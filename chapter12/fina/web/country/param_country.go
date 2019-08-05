package country

import "gopkg.in/go-playground/validator.v9"

type GetCountryParam struct {
	Name  string `json:"name" validate:"required"`
	Short string `json:"short"`
}

func (G GetCountryParam) Valid() error {
	return validator.New().Struct(G)
}
