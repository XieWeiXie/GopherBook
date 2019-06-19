package model_v1

import "time"

const (
	// 准备状态、未付款状态、已付款状态
	READINESS = iota
	BALANCE
	PAID
)

var (
	STATUS_MAP = make(map[int]string)
)

func init() {
	STATUS_MAP[READINESS] = "准备状态"
	STATUS_MAP[BALANCE] = "未付款状态"
	STATUS_MAP[PAID] = "已付款状态"
}

type Order struct {
	base       `xorm:"extends"`
	ProductIds []int `xorm:"blob"`
	Status     int
	AccountId  int64
	Account    Account `xorm:"-"`
}
type OrderSerializer struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
	Phone     string    `json:"phone"`
	AccountId uint      `json:"account_id"`
}

func (o Order) Serializer() OrderSerializer {
	return OrderSerializer{
		Id:        o.ID,
		CreatedAt: o.CreatedAt.Truncate(time.Second),
		UpdatedAt: o.UpdatedAt.Truncate(time.Second),
		Status:    STATUS_MAP[o.Status],
		AccountId: o.Account.ID,
		Phone:     o.Account.Phone,
	}
}
