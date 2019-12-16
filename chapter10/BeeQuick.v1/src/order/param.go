package order

import (
	"github.com/wuxiaoxiaoshen/GopherBook/chapter10/BeeQuick.v1/src/make_param"
	"gopkg.in/go-playground/validator.v9"
)

type GetOrderParam struct {
	make_param.ReturnAll
}

type PatchOrderParam struct {
	Status string `json:"status" validate:"eq=readiness|eq=balance|eq=paid"`
}

func (p PatchOrderParam) Valid() error {
	return validator.New().Struct(p)
}

type PostOrderParam struct {
	ProductIds []int `json:"product_ids" validate:"required"`
	AccountId  int   `json:"account_id" validate:"required"`
}

func (p PostOrderParam) Valid() error {
	return validator.New().Struct(p)
}
