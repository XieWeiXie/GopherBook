package product

type CreateProductParam struct {
	Name          string  `json:"name" validate:"required"`
	ShopId        int64   `json:"shop_id" validate:"required"`
	Avatar        string  `json:"avatar"`
	Price         float64 `json:"price" validate:"required"`
	Discount      float64 `json:"discount"`
	Specification string  `json:"specification"`
}
