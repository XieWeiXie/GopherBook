package HighCharts

import "net/http"

type ChartBase interface {
	Plot(w http.ResponseWriter, r *http.Request)
	Save(string) bool
	Name() string
	Type() string
}
