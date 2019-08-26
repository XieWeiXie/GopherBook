package HighCharts

import "net/http"

type ColumnHighCharts struct {
	Chart
	Title
	XAixs
	YAixs
	PlotOptions
	Series
	typeName string `json:"type"`
}

func (C ColumnHighCharts) Plot(w http.ResponseWriter, r *http.Request) {
	toHandler(w, r, C)
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
func (C ColumnHighCharts) String() string {
	return C.Name()
}

func NewColumn(title string) *ColumnHighCharts {
	return &ColumnHighCharts{
		Chart: Chart{
			TypeName:    COLUMNTYPE,
			MarginRight: 40,
			MarginTop:   80,
			value:       make(map[string]interface{}),
		},
		Title: Title{
			Text: title,
		},
	}
}
