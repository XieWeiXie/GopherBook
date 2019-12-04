package BizCharts

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

type Data struct {
	Year       string  `json:"year"`
	Population float32 `json:"population"`
}

func TestPolar(tests *testing.T) {
	polar := NewPolar(DEFAULT_THEME)
	//{year: '2001', population: 41.8 },
	//{year: '2002', population: 38 },
	//{year: '2003', population: 33.7 },
	//{year: '2004', population: 30.7 },
	//{year: '2005', population: 25.8 },
	//{year: '2006', population: 31.7 },
	//{year: '2007', population: 33 },
	//{year: '2008', population: 46 },
	//{year: '2009', population: 38.3 },
	//{year: '2010', population: 28 },
	//{year: '2011', population: 42.5 },
	//{year: '2012', population: 30.3 }
	data := []Data{
		{Year: "2001", Population: 41.8},
		{Year: "2002", Population: 38},
		{Year: "2003", Population: 33.7},
		{Year: "2004", Population: 30.7},
		{Year: "2005", Population: 25.8},
		{Year: "2006", Population: 31.7},
		{Year: "2007", Population: 33},
		{Year: "2008", Population: 46},
		{Year: "2009", Population: 38.3},
		{Year: "2010", Population: 28},
		{Year: "2011", Population: 42.5},
		{Year: "2012", Population: 30.3},
	}
	var temp []interface{}
	for _, i := range data {
		r, _ := ToMap(i)
		temp = append(temp, r)
	}
	tt, _ := json.Marshal(temp)
	polar.Data = string(tt)
	polar.position = "year*population"
	polar.color = "year"

	http.HandleFunc("/", polar.Plot)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
