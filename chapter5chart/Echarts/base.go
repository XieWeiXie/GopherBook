package Echarts

import "net/http"

type EChartInterface interface {
	Plot(w http.ResponseWriter, r *http.Request)
	Save(string) bool
	Name() string
	Type() string
}

type BaseType struct {
	Type string `json:"type"`
}

type BaseTitle struct {
	Title TitleOpts `json:"title"`
}

type BaseData struct {
	Series Series `json:"series"`
}

type BaseOptions struct {
	XAxis   AxisOpts    `json:"xAxis,omitempty"`
	YAxis   AxisOpts    `json:"yAxis,omitempty"`
	ToolTip ToolTipOpts `json:"tooltip,omitempty"`
	Legend  LegendOpts  `json:"legend,omitempty"`
}

type BackgroundOpts struct {
	Data string `json:"backgroundColor"`
}
