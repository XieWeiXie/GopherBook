package Echarts

import (
	"net/http"
)

type Scatter struct {
	BaseType
	BaseTitle
	BaseOptions
	BaseData
	BackgroundOpts
	json  interface{}
	theme string
}

func (S *Scatter) toJSON() {
	var V map[string]interface{}
	V = make(map[string]interface{})
	V["title"] = S.Title
	V["series"] = S.Series.Data
	V["xAxis"] = S.XAxis
	V["yAxis"] = S.YAxis
	V["legend"] = S.Legend
	V["tooltip"] = S.ToolTip
	S.json = V
}

func (S *Scatter) SetTheme(name string) {
	S.theme = name
}

func (S Scatter) Plot(w http.ResponseWriter, r *http.Request) {
	S.toJSON()
	var theme map[string]interface{}
	theme = make(map[string]interface{})
	theme["Theme"] = S.theme
	theme["Options"] = S.json
	toHandler(w, r, theme)
}

func (S Scatter) Save(name string) bool {
	S.toJSON()
	if name == "" {
		name = SCATTER
	}
	return toSave(S.json, name)
}

func (S Scatter) Name() string {
	return SCATTER
}
func (S Scatter) Type() string {
	return S.Name()
}

func NewScatter(title string) *Scatter {
	t := BaseTitle{
		Title: TitleOpts{
			Text: title,
		},
	}
	return &Scatter{
		BaseTitle: t,
		BaseType: BaseType{
			Type: SCATTER,
		},
	}
}
