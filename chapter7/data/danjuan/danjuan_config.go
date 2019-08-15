package danjuan

var (
	CONF      = "https://danjuanapp.com/djapi/v3/filter/conf?type=yield"
	FUND      = "https://danjuanapp.com/djapi/v3/filter/fund?type=%d&order_by=%s&size=%d&page=%d"
	EveryFund = "https://danjuanapp.com/djapi/index_eva/dj"
)

type FundQuery struct {
	TypeID  int
	OrderBy string
	Size    int
	Page    int
}

var ConfType = []string{"yield", "zsph", "pdph"}
