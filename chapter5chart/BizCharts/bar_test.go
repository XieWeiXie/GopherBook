package BizCharts

import (
	"fmt"
	"testing"
)

func TestNewBar(t *testing.T) {
	bar := NewBar("top")
	type one struct {
		Genre  string
		Sold   int
		Income int
	}
	data := BaseData{
		Data: []interface{}{
			one{
				"Sports", 275, 2300,
			}, one{
				"Strategy", 115, 667,
			}, one{
				"Action", 120, 982,
			}, one{
				"Shooter", 350, 5271,
			}, one{
				"Other", 150, 3710,
			},
		},
	}
	bar.Data = data
	cols := BaseCols{
		X: OneCol{
			Alias: "销售额",
		},
		Y: OneCol{
			Alias: "游戏种类",
		},
	}
	bar.BaseCols = cols
	bar.valueFormat()
	fmt.Println(bar.value)
}
