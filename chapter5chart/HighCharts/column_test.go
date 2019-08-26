package HighCharts

import (
	"net/http"
	"testing"

	"github.com/micro/go-log"
)

func TestColumn(tests *testing.T) {
	column := NewColumn("3D Example")
	column.AddField("options3d", Options3d{
		Enabled:      true,
		Alpha:        15,
		Beta:         15,
		ViewDistance: 25,
		Depth:        40,
	})
	column.AddCategories([]string{"苹果", "橘子", "梨", "葡萄", "香蕉"})
	column.AddYAxisTitle("水果数量")
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
	column.Series = data
	http.HandleFunc("/", column.Plot)
	log.Fatal(http.ListenAndServe(":7777", nil))

}
