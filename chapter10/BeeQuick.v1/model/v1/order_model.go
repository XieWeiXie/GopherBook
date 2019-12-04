package model_v1

import "time"

const (
	// 准备状态、未付款状态、已付款状态
	READINESS = iota
	BALANCE
	PAID
)

var (
	STATUS_MAP    = make(map[int]string)
	STATUS_MAP_EN = make(map[int]string)
)

func init() {
	STATUS_MAP[READINESS] = "准备状态"
	STATUS_MAP[BALANCE] = "未付款状态"
	STATUS_MAP[PAID] = "已付款状态"
	STATUS_MAP_EN[READINESS] = "readiness"
	STATUS_MAP_EN[BALANCE] = "balance"
	STATUS_MAP_EN[PAID] = "paid"

}

type Order struct {
	base       `xorm:"extends"`
	ProductIds []int `xorm:"blob"`
	Status     int
	AccountId  int64
	Account    Account `xorm:"-"`
	Total      float64
}

func (o Order) TableName() string {
	return "beeQuick_order"
}

type OrderSerializer struct {
	Id         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Status     string    `json:"status"`
	Phone      string    `json:"phone"`
	AccountId  uint      `json:"account_id"`
	Total      float64   `json:"total"`
	ProductIds []int     `json:"product_ids"`
}

func (o Order) Serializer() OrderSerializer {
	return OrderSerializer{
		Id:         o.ID,
		CreatedAt:  o.CreatedAt.Truncate(time.Second),
		UpdatedAt:  o.UpdatedAt.Truncate(time.Second),
		Status:     STATUS_MAP[o.Status],
		AccountId:  o.Account.ID,
		Phone:      o.Account.Phone,
		Total:      o.Total,
		ProductIds: o.ProductIds,
	}
}
