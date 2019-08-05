package history

import "gopkg.in/go-playground/validator.v9"

type GetHistoryParam struct {
	Year int `json:"year" validator:"min=1973"`
}

func (G GetHistoryParam) Valid() error {
	return validator.New().Struct(G)

}

type GetAllHistoryParam struct {
	OrderBy string `json:"order_by" validator:"eq=id|eq=year"`
}

func (G GetAllHistoryParam) Valid() error {
	return validator.New().Struct(G)
}
