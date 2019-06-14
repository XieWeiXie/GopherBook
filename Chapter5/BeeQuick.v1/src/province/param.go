package province

import "gopkg.in/go-playground/validator.v9"

type GetProvinceParam struct {
	Level  string `json:"level" validate:"eq=province|eq=city|eq=district"`
	Return string `json:"return" validate:"eq=all_list|eq=all_count"`
}

func (g GetProvinceParam) Valid() error {
	valid := validator.New()
	return valid.Struct(g)
}
