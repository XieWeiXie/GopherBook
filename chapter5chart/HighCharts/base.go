package HighCharts

import "net/http"

type ChartBase interface {
	Plot(w http.ResponseWriter, r *http.Request)
	Save(string) bool
	Name() string
	Type() string
}

type Base struct {
	Chart  Chart  `json:"chart"`
	Title  Title  `json:"title"`
	Series Series `json:"series"`
}

type Chart struct {
	TypeName string `json:"type"`
}

type Title struct {
	Text string `json:"text"`
}

type Series struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}
