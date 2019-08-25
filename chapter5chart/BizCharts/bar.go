package BizCharts

import "net/http"

type Bar struct {
	BaseTheme
	BaseData
	BaseCols
	BaseLegend
	value    interface{}
	typename string
	color    string
}

func (B Bar) Plot(w http.ResponseWriter, r *http.Request) {
	B.valueFormat()
	toHandler(w, r, B.value)
}

func (B *Bar) valueFormat() {
	var v = make(map[string]interface{})
	v["Theme"] = B.Theme
	v["Data"] = B.Data
	v["XYAxis"] = B.BaseCols
	v["Axis"] = B.Keys()[0]
	v["Yxis"] = B.Keys()[1]
	v["Position"] = B.Position()
	v["Color"] = B.color
	v["Size"] = 10
	v["Type"] = B.typename
	v["Location"] = B.Location
	B.value = v
}

func (B *Bar) Save(name string) bool {
	if name == "" {
		name = "bar"
	}
	return toSave(name, B.value)
}

func (B *Bar) Name() string {
	return B.typename
}
func (B *Bar) String() string {
	return B.Name()
}

func NewBar(local string) *Bar {
	return &Bar{
		BaseTheme: BaseTheme{
			Theme: DEFAULT_THEME,
		},
		BaseLegend: BaseLegend{
			Location: local,
		},
		typename: BARTYPE,
	}
}
