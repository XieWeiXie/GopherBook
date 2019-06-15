package activity

import (
	"fmt"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type CreateActivityParam struct {
	Name    string `json:"name" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Start   string `json:"start" validate:"required"`
	End     string `json:"end" validate:"required"`
	Avatar  string `json:"avatar"`
	ShopIds []int  `json:"shop_ids" validate:"required"`
}

func (c CreateActivityParam) Valid() error {
	return validator.New().Struct(c)
}

func (c CreateActivityParam) timeHandle() (time.Time, time.Time, error) {
	start := toTime(c.Start)
	end := toTime(c.End)
	if end.Before(start) {
		return time.Time{}, time.Time{}, fmt.Errorf("end before start")
	}
	return start, end, nil
}
