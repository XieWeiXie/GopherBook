package BizCharts

import (
	"github.com/gobuffalo/packr"
	"github.com/qiniu/x/log.v7"
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
