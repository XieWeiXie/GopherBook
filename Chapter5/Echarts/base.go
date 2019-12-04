package Echarts

import "net/http"

// 定义 图表接口
type EChartInterface interface {
	Plot(w http.ResponseWriter, r *http.Request)
	Save(string) bool
	Name() string
	Type() string
}

// 图表类型
type BaseType struct {
	Type string `json:"type"`
}

// 图表标题
type BaseTitle struct {
	Title TitleOpts `json:"title"`
}

// 基本数据
type BaseData struct {
	Series Series `json:"series"`
}

// 基本配置项
type BaseOptions struct {
	XAxis   AxisOpts    `json:"xAxis,omitempty"`
	YAxis   AxisOpts    `json:"yAxis,omitempty"`
	ToolTip ToolTipOpts `json:"tooltip,omitempty"`
	Legend  LegendOpts  `json:"legend,omitempty"`
}

// 背景选项
type BackgroundOpts struct {
	Data string `json:"backgroundColor"`
}
