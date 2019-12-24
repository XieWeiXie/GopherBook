package HighCharts

import (
	"log"

	"github.com/gobuffalo/packr"
)

func PlotText() string {
	box := packr.NewBox("./template")
	plot, err := box.FindString("plot.html")
	if err != nil {
		log.Println(err)
		return "-1"
	}
	return plot
}
