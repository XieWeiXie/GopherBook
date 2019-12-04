package HighCharts

import (
	"log"
	"net/http"
	"testing"
)

func TestColumn(tests *testing.T) {
	c := NewColumn("3D Example")
	c.Options3d = Options3d{
		Enabled:      true,
		Alpha:        15,
		Beta:         15,
		ViewDistance: 25,
		Depth:        40,
	}
	c.XAxis.AddCategories([]string{"苹果", "橘子", "梨", "葡萄", "香蕉"})
	c.YAxis.AddYAxisTitle("水果数量")
	c.YAxis.AllowDecimal()
	c.AddProperty("plotOptions", PlotOptions{
		Column: Column{
			Stacking: "normal",
			Depth:    40,
		},
	})
	data := Series{
		Data: []OneSeries{
			{
				Name:  "小张",
				Data:  []int{5, 3, 4, 7, 2},
				Stack: "male",
			}, {
				Name:  "小王",
				Data:  []int{3, 4, 4, 2, 5},
				Stack: "male",
			}, {
				Name:  "小彭",
				Data:  []int{2, 5, 6, 2, 1},
				Stack: "female",
			}, {
				Name:  "小潘",
				Data:  []int{3, 0, 4, 4, 3},
				Stack: "female",
			},
		},
	}
	c.Series = data
	http.HandleFunc("/", c.Plot)
	log.Fatal(http.ListenAndServe(":7777", nil))

}
