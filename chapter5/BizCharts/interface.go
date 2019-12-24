package BizCharts

import "net/http"

type BizChartInterface interface {
	Plot(w http.ResponseWriter, r *http.Request)
	Save(string) bool
	Name() string
	Type() string
}
