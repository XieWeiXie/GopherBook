package BizCharts

import "net/http"

type Bar struct {
	BaseData
	BaseCols
	value interface{}
	theme string
}

func (B Bar) Plot(w http.ResponseWriter, r *http.Request) {
	toHandler(w, r, B.value)
}

func (B *Bar) valueFormat() {}

func (B *Bar) Save(name string) bool {
	if name == "" {
		name = "bar"
	}
	return toSave(name, B.value)
}
