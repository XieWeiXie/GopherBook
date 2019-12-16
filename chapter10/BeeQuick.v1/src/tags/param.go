package tags

import (
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/make_param"
	"gopkg.in/go-playground/validator.v9"
)

type CreateTagParam struct {
	Name string `json:"name" validate:"required"`
}

func (c CreateTagParam) Valid() error {
	return validator.New().Struct(c)
}

type CreateTagsParam struct {
	Data []CreateTagParam `json:"data" validate:"required,dive,required"`
}

func (c CreateTagsParam) Valid() error {
	return validator.New().Struct(c)
}

type GetTagsParam struct {
	make_param.ReturnAll
}
