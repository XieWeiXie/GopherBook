package rule

import (
	"gopkg.in/go-playground/validator.v9"
)

type PostRuleParam struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Type     int    `json:"type" validate:"eq=0|eq=1"`
}

func (p PostRuleParam) notNull() bool {
	if p.Question == "" || p.Answer == "" {
		return false
	}
	return true
}

func (p PostRuleParam) Valid() error {
	valid := validator.New()
	return valid.Struct(p)
}

type GetRuleParam struct {
	Return string `json:"return" validate:"eq=all_list|eq=all_count"`
}

func (g GetRuleParam) Valid() error {
	valid := validator.New()
	return valid.Struct(g)
}
