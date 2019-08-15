package danjuan

import "testing"

func TestParseConf(t *testing.T) {
	ParseConf(CONF)
}
func TestParseFund(t *testing.T) {
	ParseFund("https://danjuanapp.com/djapi/v3/filter/fund?type=1&order_by=1m&size=20&page=1")
}

func TestParseEveryFund(t *testing.T) {
	ParseEveryFund("https://danjuanapp.com/djapi/index_eva/dj")
}
