package unit

type CreateUintParam struct {
	Data []OneUint `json:"data" validate:"required,dive"`
}

type OneUint struct {
	Name   string `json:"name" validate:"required"`
	EnName string `json:"en_name" validate:"required"`
}
