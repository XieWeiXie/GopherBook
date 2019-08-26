package Apexcharts

import (
	"encoding/json"
	"net/http"
)

type HeatMap struct {
	Chart
	DataLabels
	Colors
	Series
	XAxis
	Title
	Grid
	format map[string]interface{}
}

func (H *HeatMap) Format() {
	H.format["chart"] = H.Chart
	H.format["dataLabels"] = H.DataLabels
	H.format["colors"] = H.Colors.Data
	H.format["series"] = H.Series.Data
	H.format["xaxis"] = H.XAxis
	H.format["title"] = H.Title
	H.format["grid"] = H.Grid
}

func (H *HeatMap) Plot(w http.ResponseWriter, r *http.Request) {
	H.Format()
	t, _ := json.Marshal(H.format)
	toHandler(w, r, string(t))
}
func (H *HeatMap) Save(name string) bool {
	if name == "" {
		name = H.Chart.Type
	}
	H.Format()
	t, _ := json.Marshal(H.format)
	return toSave(name, string(t))
}

func (H *HeatMap) Name() string {
	return H.Chart.Type
}
func (H *HeatMap) Type() string {
	return H.Name()
}

func NewHeatMap(title string) *HeatMap {
	return &HeatMap{
		Chart: Chart{
			Height: 800,
			Type:   HEATMAPTYPE,
		},
		DataLabels: DataLabels{
			Enabled: true,
		},
		Title: Title{
			Text:  title,
			Align: "center",
		},
		format: make(map[string]interface{}),
	}
}
