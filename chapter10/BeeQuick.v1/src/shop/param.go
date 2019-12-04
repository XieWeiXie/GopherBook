package shop

import "gopkg.in/go-playground/validator.v9"

type PostShopParam struct {
	Name       string `json:"name" validate:"required"`
	Location   string `json:"location" validate:"required"`
	ProvinceId int    `json:"province_id" validate:"required"`
}

func (p PostShopParam) Valid() error {
	return validator.New().Struct(p)
}
