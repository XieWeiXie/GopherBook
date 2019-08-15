package danjuan

import "time"

type OrderModel struct {
	Name string
	Key  string
}

type OrderModels []OrderModel

type TypeModel struct {
	Name string
	Key  string
}

type TypeModels []TypeModel

type Conf struct {
	OrderModels
	TypeModels
}

type Fund struct {
	T       string `json:"type"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	SFType  string `json:"sf_type"`
	UnitNav string `json:"unit_nav"`
	Yield   string `json:"yield"`
}

type FundResult struct {
	Funds      []Fund `json:"funds"`
	TotalItems int64  `json:"total_items"`
}

type FundEvery struct {
	BeginAt      time.Time `json:"begin_at"`
	EvaType      string    `json:"eva_type"`
	IndexCode    string    `json:"index_code"`
	Name         string    `json:"name"`
	PB           float64   `json:"pb"`
	PBPercentile float64   `json:"pb_percentile"`
	PE           float64   `json:"pe"`
	PEPercentile float64   `json:"pe_percentile"`
	ROE          float64   `json:"roe"`
	CurrentDay   time.Time `json:"current_day"`
	Yield        float64   `json:"yield"`
}

type FundEveryResult struct {
	Funds      []FundEvery
	TotalItems int64 `json:"total_items"`
}
