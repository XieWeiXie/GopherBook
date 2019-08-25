package Echarts

import (
	"log"
	"net/http"
	"testing"
)

func TestScatter(test *testing.T) {
	s := NewScatter("Hello World Scatter")
	s.Series.Add(
		OneSeries{
			SymbolSize: 20,
			Name:       "Brand",
			Type:       s.Type(),
			Data: []interface{}{
				[]float32{10.0, 8.04},
				[]float32{8.0, 6.95},
				[]float32{13.0, 7.58},
				[]float32{9.0, 8.81},
				[]float32{11.0, 8.33},
				[]float32{14.0, 9.96},
				[]float32{6.0, 7.24},
				[]float32{4.0, 4.26},
				[]float32{12.0, 10.84},
				[]float32{7.0, 4.82},
				[]float32{5.0, 5.68},
			},
		},
	)
	s.SetTheme(LightTheme)

	s.BaseTitle.Title.SetPositions(BOTTOM, CENTER)
	http.HandleFunc("/", s.Plot)
	log.Fatal(http.ListenAndServe(":9998", nil))
}
