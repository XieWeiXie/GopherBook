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
