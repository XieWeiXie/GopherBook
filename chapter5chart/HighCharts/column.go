package HighCharts

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ColumnHighCharts struct {
	Chart
	Title
	XAxis       XAxis       `json:"xAxis"`
	YAxis       YAxis       `json:"yAxis"`
	PlotOptions PlotOptions `json:"plotOptions"`
	Series
	typeName string `json:"type"`
	format   map[string]interface{}
}

func (C *ColumnHighCharts) AddProperty(key string, v interface{}) {
	C.format[key] = v
}

func (C *ColumnHighCharts) Format() {
	C.format["chart"] = C.Chart
	C.format["title"] = C.Title
	C.format["xAxis"] = C.XAxis
	C.format["yAxis"] = C.YAxis
	C.format["series"] = C.Series.Data

}
func (C ColumnHighCharts) Plot(w http.ResponseWriter, r *http.Request) {
	C.Format()
	t, _ := json.Marshal(C.format)
	fmt.Println(string(t))
	toHandler(w, r, string(t))
}

func (C ColumnHighCharts) Save(name string) bool {
	if name == "" {
		name = C.typeName
	}
	return toSave(name, C)
}

func (C ColumnHighCharts) Name() string {
	return C.typeName
}

func NewColumn(title string) *ColumnHighCharts {
	return &ColumnHighCharts{
		Chart: Chart{
			TypeName:    COLUMNTYPE,
			MarginRight: 40,
			MarginTop:   80,
		},
		Title: Title{
			Text: title,
		},
		format:   make(map[string]interface{}),
		typeName: COLUMNTYPE,
	}
}
