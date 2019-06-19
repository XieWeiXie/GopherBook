package product

import (
	"GopherBook/Chapter5/BeeQuick.v1/src/make_param"
	"gopkg.in/go-playground/validator.v9"
)

type CreateProductParam struct {
	Name          string  `json:"name" validate:"required"`
	ShopId        int64   `json:"shop_id" validate:"required"`
	Avatar        string  `json:"avatar"`
	Price         float64 `json:"price" validate:"required"`
	Discount      float64 `json:"discount"`
	Specification string  `json:"specification" validate:"required"`
	Period        string  `json:"period" validate:"required"`
	BrandId       int64   `json:"brand_id" validate:"gt=0"`
	TagId         int64   `json:"tag_id" validate:"gt=0"`
	UnitId        int64   `json:"unit_id" validate:"gt=0"`
}

func (c CreateProductParam) Valid() error {
	return validator.New().Struct(c)
}

type GetAllProductParam struct {
	make_param.ReturnAll
}

type PatchOneParam struct {
	Name     string  `json:"name" validate:"required_with_all"`
	Price    float64 `json:"price" validate:"required_with_all"`
	Discount float64 `json:"discount" validate:"required_with_all"`
}

func (p PatchOneParam) Valid() error {
	return validator.New().Struct(p)
}

type CreateMultiplyParam struct {
	Data []CreateProductParam `json:"data" validate:"required, min=1, dive, required"`
}

func (c CreateMultiplyParam) Valid() error {
	return validator.New().Struct(c)
}
