package HighCharts

type PlotOptions struct {
	Column Column `json:"column"`
}

type Column struct {
	Stacking string `json:"stacking"`
	Depth    int    `json:"depth"`
}

type Options3d struct {
	Enabled      bool `json:"enabled"`
	Alpha        int  `json:"alpha"`
	Beta         int  `json:"beta"`
	Depth        int  `json:"depth"`
	ViewDistance int  `json:"viewDistance"`
}
