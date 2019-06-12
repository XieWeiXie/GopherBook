package model_v1

import "time"

const (
	// 兑换券，优惠券
	EXCHANGE = iota
	COUPON
)

type ExchangeCoupon struct {
	base  `xorm:"extends"`
	Name  string    `xorm:"varchar(32)" json:"name"`
	Price float64   `json:"price"`
	Total float64   `json:"total"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Token string    `json:"token"`
	Type  int       `json:"type"` // 0,1 : 兑换券 抵消价格，优惠券 类似几折
}

func (exchange ExchangeCoupon) TableName() string {
	return "beeQuick_exchange_coupons"
}

type Account2ExchangeCoupon struct {
	AccountId        int64 `xorm:"index"`
	ExchangeCouponId int64 `xorm:"index"`
	Status           int   `json:"status"` // 0,1,2:未使用，已使用，已过期
}

func (a2e Account2ExchangeCoupon) TableName() string {
	return "beeQuick_account2exchange_coupon"
}

const (
	// 未使用、已使用、已过期
	NEW = iota
	USED
	EXPIRE
)

var StatusMap = make(map[int]string)

func init() {
	StatusMap[NEW] = "未使用"
	StatusMap[USED] = "已使用"
	StatusMap[EXPIRE] = "已过期"
}

type ExchangeCouponSerializer struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Start     string    `json:"start"` //  格式：2006/01/02
	End       string    `json:"end"`   // 格式：2006/01/02
	Status    string    `json:"status"`
}

func (exchange ExchangeCoupon) Serializer(status string) ExchangeCouponSerializer {

	//status := func(now time.Time) string {
	//	start, _ := time.Parse("2006/01/02", exchange.Start)
	//	end, _ := time.Parse("2006/01/02", exchange.End)
	//	if now.After(start) && now.Before(end) {
	//		return StatusMap[NEW]
	//	}
	//	if now.After(end) {
	//		return StatusMap[EXPIRE]
	//	}
	//	return StatusMap[USED]
	//}

	return ExchangeCouponSerializer{
		ID:        exchange.ID,
		CreatedAt: exchange.CreatedAt.Truncate(time.Second),
		UpdatedAt: exchange.UpdatedAt.Truncate(time.Second),
		Name:      exchange.Name,
		Price:     exchange.Price,
		Start:     exchange.Start.Format("2006-01-02 15:04:05"),
		End:       exchange.End.Format("2006-01-02 15:04:05"),
		Status:    status,
	}
}
